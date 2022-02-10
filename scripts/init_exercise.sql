-- Create collection admin

-- CREATE TABLE IF NOT EXISTS question_group (
-- `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
-- `title` VARCHAR(500) DEFAULT '',
-- `description` VARCHAR(1000) DEFAULT '',
-- `type` VARCHAR(100) DEFAULT '',
-- `deleted` BOOLEAN NOT NULL DEFAULT 0,
-- `question_id` BIGINT(20) NOT NULL,
-- `created_at` timestamp not null default current_timestamp,
-- `updated_at` timestamp not null default current_timestamp,
-- PRIMARY KEY(`id`) ,
-- FOREIGN KEY(`question_id`) REFERENCES `question` (`id`)
-- )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS exercise (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`title` VARCHAR(500) DEFAULT '',
`description` VARCHAR(1000) DEFAULT '',
`type` VARCHAR(100) DEFAULT '',
`deleted` BOOLEAN NOT NULL DEFAULT 0,
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

