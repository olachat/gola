CREATE TABLE `gifts_nn_with_default` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL default 'gift for you' COMMENT 'gift name',
  `is_free` tinyint(1) NOT NULL default 1 COMMENT 'is free gift',
  `gift_count` smallint NOT NULL default 1,
  `gift_type` enum('', 'freebie', 'sovenir', 'membership') NOT NULL default 'membership',
  `create_time` bigint NOT NULL default 999,
  `discount` float unsigned NOT NULL default 0.1,
  `price` double unsigned NOT NULL default 5.0,
  `remark` varchar(128) NOT NULL default 'hope you like it',
  `manifest` varbinary(255) NOT NULL default 'manifest data',
  `description` text NOT NULL,
  `update_time` timestamp NOT NULL default '2023-01-19 03:14:07.0',
  `update_time2` timestamp NOT NULL default CURRENT_TIMESTAMP,
  `branches` set('orchard','vivo','sentosa','changi') NOT NULL default 'sentosa,changi',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = COMPACT COMMENT = 'gifts_nn_with_default';
