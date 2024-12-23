CREATE TABLE `blogs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'User Id',
  `slug` varchar(255) NOT NULL DEFAULT '' COMMENT 'Slug',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'Title',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Category Id',
  `is_pinned` boolean NOT NULL DEFAULT '0' COMMENT 'Is pinned to top',
  `is_vip` boolean NOT NULL DEFAULT '0' COMMENT 'Is VIP reader only',
  `country` varchar(255) NOT NULL DEFAULT '' COMMENT 'Country of the blog user',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Created Timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Updated Timestamp',
  `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'count',
  PRIMARY KEY (`id`),
  KEY `user` (`user_id`),
  KEY `user_vip` (`user_id`, `is_vip`),
  KEY `country_cate` (`country`, `category_id`, `is_vip`),
  KEY `country_vip` (`country`, `is_vip`),
  KEY `cate_pinned` (`category_id`, `is_pinned`, `is_vip`),
  KEY `user_pinned_cate` (`user_id`, `is_pinned`, `category_id`),
  KEY `user_id_count` (`user_id`, `count`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
