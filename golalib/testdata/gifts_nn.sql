CREATE TABLE `gifts_nn` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT 'gift name',
  `is_free` tinyint(1) NOT NULL COMMENT 'is free gift',
  `gift_count` smallint NOT NULL,
  `gift_type` enum('', 'freebie', 'sovenir', 'membership') NOT NULL,
  `create_time` bigint NOT NULL,
  `discount` float unsigned NOT NULL,
  `price` double unsigned NOT NULL,
  `remark` varchar(128) NOT NULL,
  `manifest` varbinary(255) NOT NULL,
  `description` text NOT NULL,
  `update_time` timestamp NOT NULL,
  `branches` set('orchard','vivo','sentosa','changi') NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = COMPACT COMMENT = 'Song list';
