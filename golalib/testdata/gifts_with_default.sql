CREATE TABLE `gifts_with_default` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) default 'gift for you' COMMENT 'gift name',
  `is_free` tinyint(1) default 1 COMMENT 'is free gift',
  `gift_count` smallint default 1,
  `gift_type` enum('', 'freebie', 'sovenir', 'membership') default 'membership',
  `create_time` bigint default 999,
  `discount` float unsigned default 0.1,
  `price` double unsigned default 5.0,
  `remark` varchar(128) default 'hope you like it',
  `manifest` varbinary(255) default 'manifest data',
  `description` text default ('default gift'),
  `update_time` timestamp default '2023-01-19 03:14:07.999999',
  `update_time2` timestamp default CURRENT_TIMESTAMP,
  `branches` set('orchard','vivo','sentosa','changi') default 'sentosa,changi',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = COMPACT COMMENT = 'Song list';
