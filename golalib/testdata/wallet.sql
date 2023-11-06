CREATE TABLE `wallet` (
  `user_id` int NOT NULL,
  `wallet_type` tinyint(2) NOT NULL,
  `wallet_name` varchar(50),
  `money` bigint NOT NULL DEFAULT 0 COMMENT 'money',
  PRIMARY KEY (`user_id`,`wallet_type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
