-- Create collection admin

CREATE TABLE `question` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`context` VARCHAR(500) DEFAULT '',
`question` VARCHAR(1000) DEFAULT '',
`option1` VARCHAR(500) DEFAULT '',
`option2` VARCHAR(500) DEFAULT '',
`option3` VARCHAR(500) DEFAULT '',
`option4` VARCHAR(500) DEFAULT '',
`score` INT DEFAULT 0,
`source` VARCHAR(100) DEFAULT '',
`year` VARCHAR(100) DEFAULT '',
`section` VARCHAR(100) DEFAULT '',
`no` INT DEFAULT 0,
`sn` VARCHAR(100) DEFAULT '',
`order` INt DEFAULT 0,
`pay_free` VARCHAR(20) DEFAULT '',
`deleted` BOOLEAN NOT NULL DEFAULT 0,
`details` VARCHAR(1000) DEFAULT '',
`answer` VARCHAR(100) DEFAULT '',
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

