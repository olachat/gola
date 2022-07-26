CREATE TABLE `account` (
  `user_id` int NOT NULL,
  `type` enum ('free', 'vip') NOT NULL DEFAULT 'free' COMMENT 'user account type',
  `money` int(8) NOT NULL DEFAULT 0 COMMENT 'Account money',
  UNIQUE KEY `user_id_type` (`user_id`, `type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
