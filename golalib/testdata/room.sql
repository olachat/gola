CREATE TABLE `room`
(
    id                 int(11) unsigned auto_increment primary key,
    `group`              int(11) unsigned default 0 not null comment '',
    lang               varchar(5)       default 'en' not null comment '',
    priority           double           default 0 not null comment '',
    deleted            tinyint(1)       default 0 not null comment '',
    KEY `lang_deleted` (`lang`, `deleted`),
    KEY `theme_group` (`group`),
    KEY `priority` (`priority`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='list of room';
