CREATE TABLE `blogs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'User Id',
  `slug` varchar(255) NOT NULL DEFAULT '' COMMENT 'Slug',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'Title',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Category Id',
  `country` varchar(255) NOT NULL DEFAULT '' COMMENT 'Country of the blog user',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Created Timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Updated Timestamp',
  PRIMARY KEY (`id`),
  KEY `user` (`user_id`),
  KEY `country_cate` (`country`, `category_id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
