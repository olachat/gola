CREATE TABLE `xs_pay_change_new` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `dateline` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '记录时间',
  `money` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '金额变化，都是正数',
  `op` enum('pay','consume','income','cash','return','income-lock','income-unlock','income-back','back','change','collect','confiscate','cash-back','punish','punish-back','give','refund','subsidy') NOT NULL DEFAULT 'pay' COMMENT '金额变化方式(pay 充值, consume 消费 income 收入 cash 提现 confiscate 官方没收 refund 原路退款 subsidy 补贴 give 赠送)',
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '序列化数据，关联数据',
  `subject` varchar(64) NOT NULL DEFAULT '',
  `deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户删除',
  PRIMARY KEY (`id`,`dateline`),
  KEY `uid_2` (`uid`,`deleted`,`dateline`,`op`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
