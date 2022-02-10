-- Create collection admin

CREATE TABLE IF NOT EXISTS question_group (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`exercise_id` BIGINT(20) NOT NULL,
`question_id` BIGINT(20) NOT NULL,
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`) ,
CONSTRAINT `question_ibfk_1` FOREIGN KEY(`question_id`) REFERENCES `question` (`id`),
CONSTRAINT `exercise_ibfk_1`  FOREIGN KEY(`exercise_id`) REFERENCES `exercise` (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;