-- Create collection admin

CREATE TABLE `admin` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`nickname` VARCHAR(20) DEFAULT '',
`password` VARCHAR(20) DEFAULT '',
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
