CREATE TABLE `song_user_favourite` (
  `user_id` int(10) unsigned NOT NULL COMMENT 'User ID',
  `song_id` int(10) unsigned NOT NULL COMMENT 'Song ID',
  `is_favourite` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'Is favourite',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last Update Time',
  PRIMARY KEY (`user_id`,`song_id`),
  KEY `idx_uid_collection` (`user_id`,`is_favourite`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='User favourite song record';
