-- Your SQL goes here
CREATE TABLE `stat` (
  `id` int(11) NOT NULL,
  `succeed` int(11) unsigned NOT NULL DEFAULT '0',
  `failed` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;