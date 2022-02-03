package tests

import (
	"testing"

	"github.com/olachat/gola/testdata/blogs"
)

func TestBlogMethods(t *testing.T) {
	blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).All[struct{
		blogs.Id,
		blogs.Title,
	}]()
	blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).FindPrimayKeys()
	blogs.Query().CountryEqual("SG").Count()
	blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).Count()
	blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).OrderBy(
		blogs.IdAsc,
	).Limit(limit, offset).Select[struct{
		blogs.Id,
		blogs.Title,
	}]()
}
