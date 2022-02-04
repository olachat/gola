CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'Name',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'Email address',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Created Timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Updated Timestamp',
  `float_type` float NOT NULL DEFAULT 0.0 COMMENT 'float',
  `double_type` double NOT NULL DEFAULT 0.0 COMMENT 'double',
  `hobby` enum ('swimming', 'running', 'singing') NOT NULL DEFAULT 'swimming' COMMENT 'user hobby',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
