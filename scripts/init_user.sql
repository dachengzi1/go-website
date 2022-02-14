-- Create collection admin

CREATE TABLE `user` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`mobile` VARCHAR(20) NOT NULL unique ,
`nick_name` VARCHAR(20) NOT NULL unique,
`password` VARCHAR(20) DEFAULT '',
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
