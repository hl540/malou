/*
 Navicat Premium Dump SQL

 Source Server         : 127.0.0.1_3306
 Source Server Type    : MySQL
 Source Server Version : 90001 (9.0.1)
 Source Host           : 127.0.0.1:3306
 Source Schema         : malou

 Target Server Type    : MySQL
 Target Server Version : 90001 (9.0.1)
 File Encoding         : 65001

 Date: 29/07/2024 00:18:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ml_pipeline
-- ----------------------------
DROP TABLE IF EXISTS `ml_pipeline`;
CREATE TABLE `ml_pipeline`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` int UNSIGNED NOT NULL,
  `updated_at` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_pipeline
-- ----------------------------
INSERT INTO `ml_pipeline` VALUES (5, 'pipeline测试', 1722175821, 1722176523);

-- ----------------------------
-- Table structure for ml_pipeline_instance_log
-- ----------------------------
DROP TABLE IF EXISTS `ml_pipeline_instance_log`;
CREATE TABLE `ml_pipeline_instance_log`  (
  `pipeline_instance_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `step_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cmd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `timestamp` int NOT NULL,
  `duration` int NOT NULL,
  PRIMARY KEY (`pipeline_instance_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_pipeline_instance_log
-- ----------------------------
INSERT INTO `ml_pipeline_instance_log` VALUES ('asd', 'asd', 'asd', 'ads', 'ad', 123, 12);

-- ----------------------------
-- Table structure for ml_pipeline_step
-- ----------------------------
DROP TABLE IF EXISTS `ml_pipeline_step`;
CREATE TABLE `ml_pipeline_step`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pipeline_id` int UNSIGNED NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pipeline_id`(`pipeline_id` ASC) USING BTREE,
  CONSTRAINT `ml_pipeline_step_ibfk_1` FOREIGN KEY (`pipeline_id`) REFERENCES `ml_pipeline` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_pipeline_step
-- ----------------------------
INSERT INTO `ml_pipeline_step` VALUES (7, 5, 'checkout', 'alpine:3.18');
INSERT INTO `ml_pipeline_step` VALUES (8, 5, 'build', 'alpine:3.18');

-- ----------------------------
-- Table structure for ml_pipeline_step_cmd
-- ----------------------------
DROP TABLE IF EXISTS `ml_pipeline_step_cmd`;
CREATE TABLE `ml_pipeline_step_cmd`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pipeline_id` int UNSIGNED NOT NULL,
  `pipeline_step_id` int UNSIGNED NOT NULL,
  `cmd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ml_pipeline_step_cmd_ibfk_1`(`pipeline_id` ASC) USING BTREE,
  INDEX `ml_pipeline_step_cmd_ibfk_2`(`pipeline_step_id` ASC) USING BTREE,
  CONSTRAINT `ml_pipeline_step_cmd_ibfk_1` FOREIGN KEY (`pipeline_id`) REFERENCES `ml_pipeline` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ml_pipeline_step_cmd_ibfk_2` FOREIGN KEY (`pipeline_step_id`) REFERENCES `ml_pipeline_step` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_pipeline_step_cmd
-- ----------------------------
INSERT INTO `ml_pipeline_step_cmd` VALUES (15, 5, 7, 'ls -l -a');
INSERT INTO `ml_pipeline_step_cmd` VALUES (16, 5, 7, 'echo $(pwd)');
INSERT INTO `ml_pipeline_step_cmd` VALUES (17, 5, 7, 'echo $(uname -a) > log.txt');
INSERT INTO `ml_pipeline_step_cmd` VALUES (18, 5, 8, 'echo $(pwd)');
INSERT INTO `ml_pipeline_step_cmd` VALUES (19, 5, 8, 'ls -l -a');
INSERT INTO `ml_pipeline_step_cmd` VALUES (20, 5, 8, 'ping baidu.com');

-- ----------------------------
-- Table structure for ml_runner
-- ----------------------------
DROP TABLE IF EXISTS `ml_runner`;
CREATE TABLE `ml_runner`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `secret` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int UNSIGNED NOT NULL DEFAULT 0,
  `created_at` int UNSIGNED NOT NULL,
  `updated_at` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ml_runner
-- ----------------------------
INSERT INTO `ml_runner` VALUES (19, 'f063f1f1-8170-4d87-9807-8bae7c991394', '2p157p8v88tzrsqx5fqb', '测试runner1_更新后', 0, 1722179001, 1722179248);

-- ----------------------------
-- Table structure for ml_runner_env
-- ----------------------------
DROP TABLE IF EXISTS `ml_runner_env`;
CREATE TABLE `ml_runner_env`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `runner_id` int UNSIGNED NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `value` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `runner_id`(`runner_id` ASC) USING BTREE,
  CONSTRAINT `ml_runner_env_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_runner_env
-- ----------------------------
INSERT INTO `ml_runner_env` VALUES (10, 19, 'k3', 'v3');
INSERT INTO `ml_runner_env` VALUES (11, 19, 'k4', 'v4');
INSERT INTO `ml_runner_env` VALUES (12, 19, 'k1', 'v1');
INSERT INTO `ml_runner_env` VALUES (13, 19, 'k2', 'v2');

-- ----------------------------
-- Table structure for ml_runner_health
-- ----------------------------
DROP TABLE IF EXISTS `ml_runner_health`;
CREATE TABLE `ml_runner_health`  (
  `runner_id` int UNSIGNED NOT NULL,
  `cpu_percent` float UNSIGNED NOT NULL,
  `memory_total` float UNSIGNED NOT NULL,
  `memory_used` float UNSIGNED NOT NULL,
  `memory_free` float UNSIGNED NOT NULL,
  `memory_used_percent` float UNSIGNED NOT NULL,
  `disk_total` float UNSIGNED NOT NULL,
  `disk_used` float UNSIGNED NOT NULL,
  `disk_free` float UNSIGNED NOT NULL,
  `disk_used_percent` float UNSIGNED NOT NULL,
  `worker_status` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` int UNSIGNED NOT NULL,
  PRIMARY KEY (`runner_id`, `created_at`) USING BTREE,
  CONSTRAINT `ml_runner_health_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ml_runner_health
-- ----------------------------
INSERT INTO `ml_runner_health` VALUES (19, 10.26, 15.8, 14.78, 1.01, 93, 220.23, 113.09, 107.14, 51.3515, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183506);
INSERT INTO `ml_runner_health` VALUES (19, 10.3, 15.8, 14.8, 0.99, 93, 220.23, 113.09, 107.14, 51.3515, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183507);
INSERT INTO `ml_runner_health` VALUES (19, 17.14, 15.8, 14.8, 1, 93, 220.23, 113.09, 107.14, 51.3515, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183508);
INSERT INTO `ml_runner_health` VALUES (19, 6.92, 15.8, 14.8, 0.99, 93, 220.23, 113.09, 107.14, 51.3515, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183509);
INSERT INTO `ml_runner_health` VALUES (19, 9.87, 15.8, 14.79, 1.01, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183510);
INSERT INTO `ml_runner_health` VALUES (19, 5.38, 15.8, 14.77, 1.03, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183511);
INSERT INTO `ml_runner_health` VALUES (19, 11.33, 15.8, 14.77, 1.03, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183512);
INSERT INTO `ml_runner_health` VALUES (19, 5.51, 15.8, 14.78, 1.02, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183513);
INSERT INTO `ml_runner_health` VALUES (19, 11.83, 15.8, 14.77, 1.03, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183514);
INSERT INTO `ml_runner_health` VALUES (19, 6.73, 15.8, 14.78, 1.02, 93, 220.23, 113.09, 107.14, 51.3516, '{\"5P3NBX0035\":\"\",\"Q2M1N7UGL2\":\"\"}', 1722183515);

-- ----------------------------
-- Table structure for ml_runner_label
-- ----------------------------
DROP TABLE IF EXISTS `ml_runner_label`;
CREATE TABLE `ml_runner_label`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `runner_id` int UNSIGNED NOT NULL,
  `value` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `runner_id`(`runner_id` ASC) USING BTREE,
  CONSTRAINT `ml_runner_label_ibfk_1` FOREIGN KEY (`runner_id`) REFERENCES `ml_runner` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ml_runner_label
-- ----------------------------
INSERT INTO `ml_runner_label` VALUES (13, 19, 'local');
INSERT INTO `ml_runner_label` VALUES (14, 19, '标签1');
INSERT INTO `ml_runner_label` VALUES (15, 19, '标签2');
INSERT INTO `ml_runner_label` VALUES (16, 19, '标签3');

SET FOREIGN_KEY_CHECKS = 1;
