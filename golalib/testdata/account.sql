CREATE TABLE `account` (
  `user_id` int NOT NULL,
  `type` enum ('free', 'vip') NOT NULL DEFAULT 'free' COMMENT 'user account type',
  `country_code` mediumint(6) unsigned NOT NULL DEFAULT 0 COMMENT 'user country code',
  `money` int(8) NOT NULL DEFAULT 0 COMMENT 'Account money',
  UNIQUE KEY `user_id_country_code` (`user_id`, `country_code`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
