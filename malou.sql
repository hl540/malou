CREATE TABLE `ml_runner`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `code`       char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci    NOT NULL,
    `secret`     char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci    NOT NULL,
    `name`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `status`     int unsigned NOT NULL DEFAULT '0',
    `created_at` int unsigned NOT NULL,
    `updated_at` int unsigned NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
CREATE TABLE `ml_runner_label`
(
    `id`        int unsigned NOT NULL AUTO_INCREMENT,
    `runner_id` int unsigned NOT NULL,
    `value`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `runner_id` (`runner_id`),
    CONSTRAINT `ml_runner_label_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
CREATE TABLE `ml_runner_env`
(
    `id`        int unsigned NOT NULL AUTO_INCREMENT,
    `runner_id` int unsigned NOT NULL,
    `name`      varchar(50) NOT NULL,
    `value`     varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY         `runner_id` (`runner_id`),
    CONSTRAINT `ml_runner_env_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `ml_runner_health`
(
    `runner_id`           int unsigned NOT NULL,
    `cpu_percent`         float unsigned NOT NULL,
    `memory_total`        float unsigned NOT NULL,
    `memory_used`         float unsigned NOT NULL,
    `memory_free`         float unsigned NOT NULL,
    `memory_used_percent` float unsigned NOT NULL,
    `disk_total`          float unsigned NOT NULL,
    `disk_used`           float unsigned NOT NULL,
    `disk_free`           float unsigned NOT NULL,
    `disk_used_percent`   float unsigned NOT NULL,
    `worker_status`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `created_at`          int unsigned NOT NULL,
    PRIMARY KEY (`runner_id`, `created_at`) USING BTREE,
    CONSTRAINT `ml_runner_health_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `ml_pipeline`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(50) DEFAULT NULL,
    `created_at` int unsigned NOT NULL,
    `updated_at` int unsigned NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `ml_pipeline_step`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT,
    `pipeline_id` int unsigned NOT NULL,
    `name`        varchar(50)  NOT NULL,
    `image`       varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    KEY           `pipeline_id` (`pipeline_id`),
    CONSTRAINT `ml_pipeline_step_ibfk_1` FOREIGN KEY (`pipeline_id`) REFERENCES `ml_pipeline` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `ml_pipeline_step_cmd`
(
    `id`               int unsigned NOT NULL AUTO_INCREMENT,
    `pipeline_id`      int unsigned NOT NULL,
    `pipeline_step_id` int unsigned NOT NULL,
    `cmd`              varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    KEY                `ml_pipeline_step_cmd_ibfk_1` (`pipeline_id`),
    KEY                `ml_pipeline_step_cmd_ibfk_2` (`pipeline_step_id`),
    CONSTRAINT `ml_pipeline_step_cmd_ibfk_1` FOREIGN KEY (`pipeline_id`) REFERENCES `ml_pipeline` (`id`),
    CONSTRAINT `ml_pipeline_step_cmd_ibfk_2` FOREIGN KEY (`pipeline_step_id`) REFERENCES `ml_pipeline_step` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `ml_pipeline_instance_log`
(
    `pipeline_instance_id` char(36)    NOT NULL,
    `step_name`            varchar(50)  DEFAULT NULL,
    `cmd`                  varchar(255) DEFAULT NULL,
    `result`               text,
    `type`                 varchar(50) NOT NULL,
    `timestamp`            int         NOT NULL,
    `duration`             int         NOT NULL,
    PRIMARY KEY (`pipeline_instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;