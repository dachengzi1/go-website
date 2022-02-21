

-- Create collection my_question_group

CREATE TABLE IF NOT EXISTS my_question_group (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`my_exercise_id` BIGINT(20) NOT NULL,
`my_question_id` BIGINT(20) NOT NULL,
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`) ,
CONSTRAINT `my_question_ibfk_1` FOREIGN KEY(`my_question_id`) REFERENCES `my_question` (`id`),
CONSTRAINT `my_exercise_ibfk_1`  FOREIGN KEY(`my_exercise_id`) REFERENCES `my_exercise` (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


