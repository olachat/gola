CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'Name',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'Email address',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Created Timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'Updated Timestamp',
  `float_type` float NOT NULL DEFAULT 0.0 COMMENT 'float',
  `double_type` double NOT NULL DEFAULT 0.0 COMMENT 'double',
  `hobby` enum ('swimming', 'running', 'singing') NOT NULL DEFAULT 'swimming' COMMENT 'user hobby',
  `hobby_no_default` enum ('swimming', 'running', 'singing') NOT NULL COMMENT 'user hobby',
  `sports` SET('swim', 'tennis', 'basketball', 'football', 'squash', 'badminton') NOT NULL DEFAULT "swim,football" COMMENT 'user sports',
  `sports2` SET('swim', 'tennis', 'basketball', 'football', 'squash', 'badminton') NOT NULL DEFAULT "swim,football" COMMENT 'user sports',
  `sports_no_default` SET('swim', 'tennis', 'basketball', 'football', 'squash', 'badminton') NOT NULL COMMENT 'user sports',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
