```sql
CREATE TABLE `blogs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'User Id',
  `slug` varchar(255) NOT NULL DEFAULT '' COMMENT 'Slug',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'Title',
  `category_id` int(11) NOT NULL DEFAULT '' COMMENT 'Category Id',
  `country` varchar(255) NOT NULL DEFAULT '' COMMENT 'Country of the blog user',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Created Timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Updated Timestamp',
  PRIMARY KEY (`id`),
  KEY `user` (`user_id`),
  KEY `country_cate` (`country`, `category_id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

```go
func TestBlogMethods(t *testing.T) {
  blogs.IdxXXXX().CountryEqual("SG").Count()
  blogs.IdxXXXX().CountryEqual("SG").Find[]()
  blogs.IdxXXXX().CountryIn("SG", "CN").CategoryIdIn(1, 2)
  blogs.IdxXXXX().CountryIn("SG", "CN").CategoryIdGreater(1)
  blogs.IdxXXXX().CountryIn("SG", "CN").CategoryIdRange(1, 10).Count()
  blogs.IdxXXXX().CountryIn("SG", "CN").CategoryIdRange(1, 10).All[blogs]()

  query := blogs.IdxCountryCategoryId().CountryEQ("SG").CategoryIdIn(1, 2)

  blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).PrimayKeys()

  blogs.CountryEQ("SG").CategoryIdIn(1, 2).All[blogs]()
  blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).FindPrimayKeys()
  blogs.Query().CountryEqual("SG").Count()
  blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).Count()
  blogs.Query().CountryEqual("SG").CategoryIdIn(1, 2).OrderBy(
    blogs.IdAsc,
  ).Limit(limit, offset)
}
```
