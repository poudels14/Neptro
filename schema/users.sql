USE neptro;

CREATE TABLE `users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(25) NOT NULL,
  `country_code` VARCHAR(10) NOT NULL,
  `created_at` BIGINT NOT NULL,
  `last_signed_in` BIGINT DEFAULT NULL,
  `last_signed_in_ip` VARCHAR(255) DEFAULT NULL,
  `archived_at` BIGINT DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;
