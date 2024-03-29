CREATE TABLE `profile` (
  `user_id` int NOT NULL,
  `level` tinyint(4) NOT NULL DEFAULT 1,
  `nick_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'Nick Name',
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
