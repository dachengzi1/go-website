

-- Create collection my_question

CREATE TABLE IF NOT EXISTS my_question (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`question_id` BIGINT(20) NOT NULL,
`user_id` BIGINT(20) NOT NULL,
`user_score` INT DEFAULT 0 ,
`status` INT DEFAULT 0 ,
`user_answer` VARCHAR(20) DEFAULT '',
`created_at` timestamp not null default current_timestamp,
`updated_at` timestamp not null default current_timestamp,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- `exercise_id` BIGINT(20) NOT NULL,

-- CONSTRAINT `user_ibfk_1` FOREIGN KEY(`user_id`) REFERENCES `user` (`id`),
-- CONSTRAINT `my_exercise_ibfk_1`  FOREIGN KEY(`exercise_id`) REFERENCES `my_exercise` (`id`)


