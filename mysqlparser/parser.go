package mysqlparser

import (
	"errors"
	"fmt"
	"strings"

	"log"

	"github.com/olachat/gola/structs"
	"vitess.io/vitess/go/vt/sqlparser"
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
	cols := make([]structs.Column, 0, len(statement.TableSpec.Columns))
	for _, colDef := range statement.TableSpec.Columns {
		opt := colDef.Type.Options
		col := structs.Column{
			Name:   colDef.Name.Lowered(),
			DBType: colDef.Type.Type,
			Table:  table,
		}
		col.FullDBType = col.DBType
		if colDef.Type.Length != nil {
			col.FullDBType = fmt.Sprintf("%s(%s)", col.FullDBType, colDef.Type.Length.Bytes())
		}
		if colDef.Type.Unsigned {
			col.FullDBType = fmt.Sprintf("%s unsigned", col.FullDBType)
		}
		if len(colDef.Type.EnumValues) > 0 {
			col.FullDBType = fmt.Sprintf("%s(%s)", col.FullDBType, strings.Join(colDef.Type.EnumValues, ","))
		}
		if opt != nil {
			if opt.Comment != nil {
				col.Comment = string(opt.Comment.Bytes())
			}
			if opt.Null == nil {
				col.Nullable = true
			} else if *opt.Null {
				col.Nullable = true
			}
			if opt.Default != nil {
				switch v := opt.Default.(type) {
				case *sqlparser.Literal:
					col.Default = v.Val
				case *sqlparser.CurTimeFuncExpr:
					col.Default = v.Name.String()
				}

			}
		}

		col.Comment = strings.ReplaceAll(col.Comment, "\r\n", " ")
		col.Comment = strings.ReplaceAll(col.Comment, "\n", " ")
		col.Comment = strings.ReplaceAll(col.Comment, "\"", "'")

		fmt.Printf("table:%s col:%s, type:%s \n", tableName, col.Name, col.FullDBType)

		cols = append(cols, col)
	}
	return cols, nil
}

func (p *MySQLParser) SetIndexAndKey(tables []*structs.Table) (err error) {
	return nil
}

func (p *MySQLParser) ForeignKeyInfo(schema, tableName string) ([]structs.ForeignKey, error) {
	return []structs.ForeignKey{}, nil
}

func (p *MySQLParser) TranslateColumnType(c structs.Column) structs.Column {
	return c
}
