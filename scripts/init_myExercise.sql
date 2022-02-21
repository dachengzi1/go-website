
-- Create collection my_exercise

CREATE TABLE IF NOT EXISTS my_exercise (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`exercise_id` BIGINT(20) NOT NULL COMMENT '题集',
`user_id` BIGINT(20) NOT NULL COMMENT '用户',
`score` INT DEFAULT 0 COMMENT '分数',
`correct_count` INT DEFAULT 0 COMMENT '答对数',
`wrong_count` INT DEFAULT 0 COMMENT '错题数',
`status` VARCHAR(50) DEFAULT 'processing' COMMENT '状态',
`created_at` timestamp not null default current_timestamp COMMENT '创建时间',
`updated_at` timestamp not null default current_timestamp COMMENT '更新时间',
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;