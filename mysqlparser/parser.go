package mysqlparser

import (
	"errors"
	"regexp"
	"strings"

	"log"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/structs"
)

type MySQLParser struct {
	tableSQLs             map[string]string
	tableCreateStatements map[string]*sqlparser.CreateTable
}

type MySQLParserConfig struct {
	DbName          string
	TableCreateSQLs []TableCreateSQL
}
type TableCreateSQL struct {
	Table     string
	CreateSQL string
}

func (p *MySQLParser) Assemble(c MySQLParserConfig) (dbinfo *structs.DBInfo, err error) {
	defer func() {
		if r := recover(); r != nil && err == nil {
			dbinfo = nil
			err = r.(error)
		}
	}()

	p.tableSQLs = map[string]string{}
	p.tableCreateStatements = map[string]*sqlparser.CreateTable{}
	dbinfo = &structs.DBInfo{}

	tables := make([]string, 0, len(c.TableCreateSQLs))
	for _, s := range c.TableCreateSQLs {
		tables = append(tables, s.Table)
		p.tableSQLs[s.Table] = s.CreateSQL
	}

	dbinfo.Schema = c.DbName
	dbinfo.Tables, err = structs.Tables(p, c.DbName, tables, []string{})
	if err != nil {
		return nil, err
	}

	return dbinfo, nil
}

func (p *MySQLParser) TableNames(schema string, whitelist, blacklist []string) ([]string, error) {
	out := make([]string, 0, len(whitelist))
	for _, table := range whitelist {
		sql := p.tableSQLs[table]
		createStatement, err := p.parse(sql)
		if err != nil {
			continue
		}
		p.tableCreateStatements[table] = createStatement
		out = append(out, table)
	}
	return out, nil
}

func (p *MySQLParser) parse(sql string) (*sqlparser.CreateTable, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		log.Printf("Error parsing SQL: %v", err)
		return nil, err
	}
	createStmt, ok := stmt.(*sqlparser.CreateTable)
	if !ok {
		log.Printf("Not a CREATE TABLE statement")
		return nil, err
	}

	return createStmt, nil
}

func (p *MySQLParser) Columns(schema string, table *structs.Table, tableName string, whitelist, blacklist []string) ([]structs.Column, error) {
	statement, found := p.tableCreateStatements[tableName]
	if !found {
		return nil, errors.New("statement not found")
	}
	cols := make([]structs.Column, 0, len(statement.Columns))
	for _, colDef := range statement.Columns {
		// opt := colDef.Type.Options
		col := structs.Column{
			Name:       colDef.Name,
			DBType:     strings.ToLower(colDef.Type),
			FullDBType: strings.ToLower(colDef.Type),
			Table:      table,
		}
		col.Unique = isColUnique(colDef.Options)
		col.Nullable = isColNullable(colDef.Options)
		col.Default = getColDefault(colDef.Options)

		re := regexp.MustCompile(`\([^)]*\)`)
		col.DBType = re.ReplaceAllString(col.DBType, "")
		col.DBType = strings.ReplaceAll(col.DBType, "  ", " ")

		col.Comment = getColComments(colDef.Options)
		col.Comment = strings.ReplaceAll(col.Comment, "\r\n", " ")
		col.Comment = strings.ReplaceAll(col.Comment, "\n", " ")
		col.Comment = strings.ReplaceAll(col.Comment, "\"", "'")
		col.Comment = strings.Trim(col.Comment, "'")
		if col.Comment == "''" {
			col.Comment = ""
		}

		cols = append(cols, col)
	}
	return cols, nil
}

func (p *MySQLParser) SetIndexAndKey(tables []*structs.Table) (err error) {
	for _, t := range tables {
		stmt := p.tableCreateStatements[t.Name]
		t.Indexes = make(map[string][]*structs.IndexDesc)
		for _, idx := range stmt.Constraints {
			hasSingleColPriKey := false
			if idx.Type == sqlparser.ConstraintPrimaryKey {
				t.PKey = &structs.PrimaryKey{}
				t.PKey.Name = idx.Name
				t.PKey.Columns = coredb.MapSlice(idx.Keys, func(col sqlparser.ColIdent) string {
					return col.Lowered()
				})
				if len(t.PKey.Columns) == 1 {
					hasSingleColPriKey = true
				}
			}
			if hasSingleColPriKey {
				continue
			}
			indicesDesc := make([]*structs.IndexDesc, 0, len(idx.Keys))
			for i, col := range idx.Keys {
				nonUniq := 1
				if idx.Type == sqlparser.ConstraintUniq ||
					idx.Type == sqlparser.ConstraintUniqIndex ||
					idx.Type == sqlparser.ConstraintPrimaryKey ||
					idx.Type == sqlparser.ConstraintUniqKey {
					nonUniq = 0
				}
				idxDesc := &structs.IndexDesc{
					Table:      t.Name,
					KeyName:    idx.Name,
					ColumnName: col.Lowered(),
					IndexType:  idx.Type.String(),
					NonUnique:  nonUniq,
					SeqInIndex: i + 1,
				}
				indicesDesc = append(indicesDesc, idxDesc)
			}
			t.Indexes[idx.Name] = indicesDesc
		}

		if t.PKey == nil {
			priKeyColumns := []string{}
			for _, col := range stmt.Columns {
				if isColPrimaryKey(col.Options) {
					priKeyColumns = append(priKeyColumns, col.Name)
				}
			}
			if len(priKeyColumns) > 0 {
				t.PKey = &structs.PrimaryKey{
					Name:    strings.Join(priKeyColumns, "_"),
					Columns: priKeyColumns,
				}
			}
			if len(priKeyColumns) > 1 {
				indicesDesc := make([]*structs.IndexDesc, 0, len(priKeyColumns))
				for i, col := range priKeyColumns {
					idxDesc := &structs.IndexDesc{
						Table:      t.Name,
						KeyName:    "PRIMARY",
						ColumnName: col,
						NonUnique:  1,
						SeqInIndex: i + 1,
					}
					indicesDesc = append(indicesDesc, idxDesc)
				}
				t.Indexes["PRIMARY"] = indicesDesc
			}
		}

		if t.PKey == nil {
			for _, col := range stmt.Columns {
				if isColUnique(col.Options) {
					t.PKey = &structs.PrimaryKey{}
					t.PKey.Name = col.Name
					t.PKey.Columns = []string{
						col.Name,
					}
					break
				}
			}
		}

		if t.PKey == nil {
			for _, idx := range stmt.Constraints {
				if idx.Type == sqlparser.ConstraintUniq ||
					idx.Type == sqlparser.ConstraintUniqIndex ||
					idx.Type == sqlparser.ConstraintUniqKey {
					t.PKey = &structs.PrimaryKey{
						Name: idx.Name,
						Columns: coredb.MapSlice(idx.Keys, func(col sqlparser.ColIdent) string {
							return col.Lowered()
						}),
					}
					break
				}
			}
		}
	}
	return nil
}

func (p *MySQLParser) ForeignKeyInfo(schema, tableName string) ([]structs.ForeignKey, error) {
	return []structs.ForeignKey{}, nil
}

func (p *MySQLParser) TranslateColumnType(c structs.Column) structs.Column {
	return c
}
