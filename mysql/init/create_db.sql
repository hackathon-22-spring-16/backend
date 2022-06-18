DROP DATABASE IF EXISTS 22hack16;
CREATE DATABASE 22hack16;
USE 22hack16;

CREATE TABLE IF NOT EXISTS `codes` (
  `user_name` varchar(32) NOT NULL,
  `hash` char(8) NOT NULL,
  `plain_code` text NOT NULL,
  `stdin` text NOT NULL,
  `title` varchar(64) NOT NULL,
  `options` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`user_name`,`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
