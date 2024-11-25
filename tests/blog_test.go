package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/coredb/txengine"
	"github.com/olachat/gola/v2/golalib/testdata/blogs"
)

func TestBlogMethods(t *testing.T) {
	blog := blogs.New()
	blog.SetTitle("foo")
	blog.SetCount(55)
	e := blog.Insert()
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 1 {
		t.Error("Insert blog 1 failed")
	}
	if blog.GetCount() != 55 {
		t.Error("Set count failed")
	}

	e = blogs.DeleteByPK(blog.GetId())
	if e != nil {
		t.Error(e)
	}
	blog = blogs.FetchByPK(1)
	if blog != nil {
		t.Error("blog 1 delete failed")
	}

	blog = blogs.New()
	blog.SetTitle("foo")
	blog.SetCount(99)
	blog.Insert()

	blog = blogs.New()
	blog.SetTitle("bar")
	blog.SetCount(88)
	blog.SetSlug("slug")
	e = blog.Insert()
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 3 {
		t.Error("Insert blog 2 failed")
	}

	count, err := blogs.Count("")
	if err != nil || count != 2 {
		t.Error("Count all blogs failed")
	}

	count, err = blogs.Count("where title=?", "bar")
	if err != nil || count != 1 {
		t.Error("Count bar blog failed")
	}

	count, err = blogs.Count("where title=?", "barbar")
	if err != nil || count != 0 {
		t.Error("Count not-exist blog failed")
	}
}

func TestBlogFind(t *testing.T) {
	obj := blogs.FindOneFromMaster("where title = ?", "bar")
	if obj.GetId() != 3 {
		t.Error("Find blog with title bar failed")
	}

	objs, err := blogs.Find("where title = ?", "bar")
	if err != nil || len(objs) != 1 {
		t.Error("Find blogs with title bar failed: ")
	}
	if objs[0].GetId() != 3 {
		t.Error("Find blogs with title bar failed")
	}

	objs, err = blogs.FindFromMaster("where title = ?", "barbar")
	if err != nil || len(objs) != 0 {
		t.Error("Find blogs with non-exist title bar failed: ")
	}
}

func TestBlogTx(t *testing.T) {
	ctx := context.Background()
	tx, err := coredb.BeginTx(ctx, blogs.DBName, &coredb.DefaultTxOpts)
	if err != nil {
		t.Fatalf("fail to start tx: %v", err)
	}
	err = txengine.StartTx(ctx, tx, func(ctx context.Context, sqlTx *sql.Tx) error {
		blogRec, err := txengine.WithTypedTx[blogs.Blog](sqlTx).FindOne(ctx, blogs.TableName, coredb.NewWhere("where title = ?", "bar"))
		if err != nil {
			return err
		}
		if blogRec.GetId() != 3 {
			t.Error("Find blog with title bar failed")
		}
		blogRec.SetCountry("GuaGua")
		updateOk, err := blogRec.UpdateTx(ctx, sqlTx)
		if err != nil {
			return err
		}
		if !updateOk {
			t.Error("fail to update blog in tx")
		}
		c, err := txengine.WithTx(sqlTx).QueryInt(ctx, "select count(*) from blogs where country = ?", "GuaGua")
		if err != nil {
			return fmt.Errorf("QueryInt failed: %w", err)
		}
		if c != 1 {
			t.Error("expect count to be 1")
		}
		newBlog := blogs.New()
		newBlog.SetCountry("GuaGua")
		newBlog.SetSlug("oldSlug")
		err = newBlog.InsertTx(ctx, sqlTx)
		if err != nil {
			return fmt.Errorf("fail to insert with Tx")
		}
		newBlog.SetSlug("slugadded123")
		_, err = newBlog.UpdateTx(ctx, sqlTx)
		if err != nil {
			return fmt.Errorf("update failed: %w", err)
		}
		return nil
	})
	if err != nil {
		t.Fatal("error encountered", err)
	}
	obj := blogs.FindOneFromMaster("where title = ?", "bar")
	if obj.GetId() != 3 {
		t.Error("Find blog with title bar failed")
	}
	if obj.GetCountry() != "GuaGua" {
		t.Error("blog not updated")
	}
	recs, err := blogs.FindCtx(ctx, "where country=? order by id desc", "GuaGua")
	if err != nil {
		t.Error(err)
	}
	if len(recs) != 2 {
		t.Error("insertTx must have failed")
	}
	if recs[0].GetSlug() != "slugadded123" {
		t.Error("wrong slug gotten")
	}

	// start a new transaction
	tx, err = coredb.BeginTx(ctx, blogs.DBName, &coredb.DefaultTxOpts)
	if err != nil {
		t.Fatalf("fail to start tx: %v", err)
	}
	err = txengine.StartTx(ctx, tx, func(ctx context.Context, sqlTx *sql.Tx) error {
		blogRec, err := txengine.WithTypedTx[blogs.Blog](sqlTx).FindOne(ctx, blogs.TableName, coredb.NewWhere("where title = ?", "bar"))
		if err != nil {
			return err
		}
		if blogRec.GetId() != 3 {
			t.Error("Find blog with title bar failed")
		}
		blogRec.SetCountry("PuaPua")
		updateOk, err := blogRec.UpdateTx(ctx, sqlTx)
		if err != nil {
			return err
		}
		if !updateOk {
			t.Error("fail to update blog in tx")
		}
		return errors.New("rollback")
	})
	if err == nil {
		t.Fatal("should get rollback error but err is nil")
	}
	// currently there is no way to test rollback using dolt in memory DB as transaction is not supported

	// start a new transaction
	tx, err = coredb.BeginTx(ctx, blogs.DBName, &coredb.DefaultTxOpts)
	if err != nil {
		t.Fatalf("fail to start tx: %v", err)
	}
	err = txengine.StartTx(ctx, tx, func(ctx context.Context, sqlTx *sql.Tx) error {
		blogRecs, err := txengine.WithTypedTx[struct {
			blogs.Id
		}](sqlTx).Find(ctx, blogs.TableName, coredb.NewWhere("where slug = ?", "slugadded123"))
		if err != nil {
			return err
		}
		if len(blogRecs) != 1 {
			t.Errorf("expected to have 1 blog")
			return errors.New("expected to have 1 blog")
		}
		id := blogRecs[0].GetId()
		_, err = txengine.WithTx(sqlTx).Exec(ctx, "delete from blogs where id=?", id)
		return err
	})
	if err != nil {
		t.Fatalf("tx delete failed")
	}
	r := blogs.FindOne("where slug = ?", "slugadded123")
	if r != nil {
		t.Error("r should already been deleted")
	}

}

func TestBlogFindT(t *testing.T) {
	obj := blogs.FindOneFieldsFromMaster[struct {
		blogs.Id
	}]("where title = ?", "bar")
	if obj.GetId() != 3 {
		t.Error("Find blog with title bar failed")
	}

	objs, err := blogs.FindFieldsFromMaster[struct {
		blogs.Id
	}]("where title = ?", "bar")
	if err != nil || len(objs) != 1 {
		t.Error("Find blogs with title bar failed: ")
	}
	if objs[0].GetId() != 3 {
		t.Error("Find blogs with title bar failed")
	}

	data, err := blogs.FindFieldsFromMaster[struct {
		blogs.Title
	}]("where title = ?", "barbar")
	if err != nil || len(data) != 0 {
		t.Error("Find blogs with non-exist title bar failed: ")
	}
}

func TestBlogSelect(t *testing.T) {
	objs := blogs.SelectFields[struct {
		blogs.Id
		blogs.Title
		blogs.Count_
	}]().OrderBy(blogs.IdAsc).AllFromMaster()

	if len(objs) != 2 {
		t.Error("Read all blog failed")
	}

	if objs[0].GetTitle() != "foo" {
		t.Error("Read blog 1 failed")
	}
	if objs[0].GetCount() != 99 {
		t.Error("Read blog 1 failed")
	}

	if objs[1].GetTitle() != "bar" {
		t.Error("Read blog 2 failed")
	}
	if objs[1].GetCount() != 88 {
		t.Error("Read blog 2 failed")
	}

	objs = blogs.SelectFields[struct {
		blogs.Id
		blogs.Title
		blogs.Count_
	}]().OrderBy(blogs.IdDesc).All()

	if len(objs) != 2 {
		t.Error("Read all blog failed")
	}

	if objs[0].GetTitle() != "bar" {
		t.Error("Read blog 1 failed")
	}

	if objs[1].GetTitle() != "foo" {
		t.Error("Read blog 2 failed")
	}

	data := blogs.FetchByPKsFromMaster(2, 3)
	if len(data) != 2 {
		t.Error("Read all blog failed")
	}

	for _, obj := range data {
		switch obj.GetId() {
		case 2:
			if obj.GetTitle() != "foo" {
				t.Error("Read blog 2 failed")
			}
		case 3:
			if obj.GetTitle() != "bar" {
				t.Error("Read blog 3 failed")
			}
		default:
			t.Error("blogs.FetchBlogByPKs load wrong id")
		}
	}
}
