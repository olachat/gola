CREATE TABLE `songs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT 'Song title',
  `rank` mediumint(8) NOT NULL DEFAULT 0 COMMENT 'Song Ranking',
  `type` enum('', '101', '1+9', '%1', '0.9') DEFAULT '',
  `hash` varchar(128) NOT NULL DEFAULT '' COMMENT 'Song file hash checksum',
  `remark` varchar(128),
  `manifest` varbinary(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash` (`hash`),
  KEY `remark` (`remark`),
  KEY `remark2` (`rank`, `remark`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = COMPACT COMMENT = 'Song list';
