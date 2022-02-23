package tests

import (
	"testing"

	"github.com/olachat/gola/testdata/blogs"
)

func TestBlogMethods(t *testing.T) {
	blog := blogs.NewBlog()
	blog.SetTitle("foo")
	e := blog.Insert()
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 1 {
		t.Error("Insert blog 1 failed")
	}

	blog = blogs.NewBlog()
	blog.SetTitle("bar")
	e = blog.Insert()
	if e != nil {
		t.Error(e)
	}

	if blog.GetId() != 2 {
		t.Error("Insert blog 2 failed")
	}

	objs := blogs.Select[struct {
		blogs.Id
		blogs.Title
	}]().OrderBy(blogs.IdAsc).All()

	if len(objs) != 2 {
		t.Error("Read all blog failed")
	}

	if objs[0].GetTitle() != "foo" {
		t.Error("Read blog 1 failed")
	}

	if objs[1].GetTitle() != "bar" {
		t.Error("Read blog 2 failed")
	}

	objs = blogs.Select[struct {
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
}
