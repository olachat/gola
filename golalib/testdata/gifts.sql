CREATE TABLE `gifts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COMMENT 'gift name',
  `is_free` tinyint(1) COMMENT 'is free gift',
  `gift_count` smallint,
  `gift_type` enum('', 'freebie', 'sovenir', 'membership'),
  `create_time` bigint,
  `discount` float unsigned,
  `price` double unsigned,
  `remark` varchar(128),
  `manifest` varbinary(255),
  `description` text,
  `update_time` timestamp,
  `branches` set('orchard','vivo','sentosa','changi'),
  PRIMARY KEY (`id`),
  KEY `idx`(`price`,`remark`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = COMPACT COMMENT = 'Song list';
