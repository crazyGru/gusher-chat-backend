alter table startup
    add protocol int not null default 0;
alter table role
    add protocol_role varchar(300) default "" not null;

CREATE TABLE `protocol`
(
    `id`   int NOT NULL AUTO_INCREMENT,
    `name` varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `protocol_task`
(
    `id`          int         NOT NULL AUTO_INCREMENT,
    `protocol_id` int         NOT NULL,
    `role`        varchar(30)  DEFAULT NULL,
    `complete_fn` varchar(30) NOT NULL,
    `depends_on`  varchar(200) DEFAULT NULL,
    `title`       varchar(100) DEFAULT NULL,
    `content`     mediumtext,
    PRIMARY KEY (`id`),
    KEY `protocol_task_protocol_id_index` (`protocol_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `protocol_task_startup`
(
    `startup_id` int NOT NULL,
    `task_id`    int NOT NULL,
    PRIMARY KEY (`startup_id`, `task_id`)
) ENGINE = InnoDB;







