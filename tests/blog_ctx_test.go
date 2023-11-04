package tests

import (
	"context"
	"testing"

	"github.com/olachat/gola/v2/golalib/testdata/blogs"
)

func TestBlogCtxMethods(t *testing.T) {
	ctx := context.Background()
	blog := blogs.New()
	blog.SetTitle("foo")
	e := blog.InsertCtx(ctx)
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 1 {
		t.Error("Insert blog 1 failed")
	}

	e = blogs.DeleteByPKCtx(ctx, blog.GetId())
	if e != nil {
		t.Error(e)
	}
	blog = blogs.FetchByPKCtx(ctx, 1)
	if blog != nil {
		t.Error("blog 1 delete failed")
	}

	blog = blogs.New()
	blog.SetTitle("foo")
	blog.Insert()

	blog = blogs.New()
	blog.SetTitle("bar")
	e = blog.Insert()
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 3 {
		t.Error("Insert blog 2 failed")
	}

	count, err := blogs.CountCtx(ctx, "")
	if err != nil || count != 2 {
		t.Error("Count all blogs failed")
	}

	count, err = blogs.CountCtx(ctx, "where title=?", "bar")
	if err != nil || count != 1 {
		t.Error("Count bar blog failed")
	}

	count, err = blogs.CountCtx(ctx, "where title=?", "barbar")
	if err != nil || count != 0 {
		t.Error("Count not-exist blog failed")
	}
}

func TestBlogFindCtx(t *testing.T) {
	ctx := context.Background()
	obj := blogs.FindOneFromMasterCtx(ctx, "where title = ?", "bar")
	if obj.GetId() != 3 {
		t.Error("Find blog with title bar failed")
	}

	objs, err := blogs.FindCtx(ctx, "where title = ?", "bar")
	if err != nil || len(objs) != 1 {
		t.Error("Find blogs with title bar failed: ")
	}
	if objs[0].GetId() != 3 {
		t.Error("Find blogs with title bar failed")
	}

	objs, err = blogs.FindFromMasterCtx(ctx, "where title = ?", "barbar")
	if err != nil || len(objs) != 0 {
		t.Error("Find blogs with non-exist title bar failed: ")
	}
}

func TestBlogFindTCtx(t *testing.T) {
	ctx := context.Background()
	obj := blogs.FindOneFieldsFromMasterCtx[struct {
		blogs.Id
	}](ctx, "where title = ?", "bar")
	if obj.GetId() != 3 {
		t.Error("Find blog with title bar failed")
	}

	objs, err := blogs.FindFieldsFromMasterCtx[struct {
		blogs.Id
	}](ctx, "where title = ?", "bar")
	if err != nil || len(objs) != 1 {
		t.Error("Find blogs with title bar failed: ")
	}
	if objs[0].GetId() != 3 {
		t.Error("Find blogs with title bar failed")
	}

	data, err := blogs.FindFieldsFromMasterCtx[struct {
		blogs.Title
	}](ctx, "where title = ?", "barbar")
	if err != nil || len(data) != 0 {
		t.Error("Find blogs with non-exist title bar failed: ")
	}
}

func TestBlogSelectCtx(t *testing.T) {
	ctx := context.Background()
	objs := blogs.SelectFields[struct {
		blogs.Id
		blogs.Title
	}]().OrderBy(blogs.IdAsc).AllFromMaster()

	if len(objs) != 2 {
		t.Error("Read all blog failed")
	}

	if objs[0].GetTitle() != "foo" {
		t.Error("Read blog 1 failed")
	}

	if objs[1].GetTitle() != "bar" {
		t.Error("Read blog 2 failed")
	}

	objs = blogs.SelectFields[struct {
		blogs.Id
		blogs.Title
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

	data := blogs.FetchByPKsFromMasterCtx(ctx, 2, 3)
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
