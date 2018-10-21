-- Your SQL goes here
CREATE TABLE `token` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `plain_text` varchar(64) NOT NULL DEFAULT '',
  `hash_data` varchar(128) NOT NULL DEFAULT '',
  `trusty` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `trusty_at` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_plain_text` (`plain_text`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;