package tests

import (
	"testing"

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
