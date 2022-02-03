package tests

import (
	"fmt"
	"testing"

	"github.com/olachat/gola/testdata/blogs"
)

func TestBlogMethods(t *testing.T) {
	data := blogs.Select[struct {
		blogs.Id
		blogs.Title
	}]().WhereCountryEQ("SG").AndCategoryIdIN(1, 2).All()

	fmt.Printf("data: %v", data)

	data2 := blogs.Select[blogs.Blog]().WhereCountryEQ("SG").OrderBy(
		blogs.CategoryIdAsc,
		blogs.IdDesc).Limit(10, 0)

	fmt.Printf("data: %v", data2)
}

// blogs.Query().CountryIN("SG", "CN").

// blogs.Query().CountryEQ("SG").CategoryIdIn(1, 2).All[struct{
// 	blogs.Id,
// 	blogs.Title,
// }]()
// blogs.Select[blogs.Blog]().WhereCountryEQ("SG").AndCategoryIdIN(1, 2)

// blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).OrderBy(
// 	blogs.IdAsc,
// ).Limit(limit, offset).Select[struct{
// 	blogs.Id,
// 	blogs.Title,
// }]()
