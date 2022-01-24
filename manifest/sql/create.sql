DROP TABLE IF EXISTS `task`;

CREATE TABLE `task` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  `delete_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'Deleted Time',

  `title` varchar(45) NOT NULL COMMENT 'title',
  `is_done` boolean NOT NULL COMMENT 'is done',
  
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;