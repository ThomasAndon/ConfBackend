/*
 Navicat Premium Data Transfer

 Source Server         : current db
 Source Server Type    : MySQL
 Source Server Version : 80023 (8.0.23)
 Source Host           : localhost:3306
 Source Schema         : terrain

 Target Server Type    : MySQL
 Target Server Version : 80023 (8.0.23)
 File Encoding         : 65001

 Date: 14/12/2023 15:08:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_hero_pcd_uoload
-- ----------------------------
DROP TABLE IF EXISTS `t_hero_pcd_uoload`;
CREATE TABLE `t_hero_pcd_uoload` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `file_uuid` varchar(255) DEFAULT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `original_uploaded_filename` varchar(255) DEFAULT NULL,
  `saved_filename` varchar(255) DEFAULT NULL,
  `file_size` bigint DEFAULT NULL,
  `save_duration` int DEFAULT NULL,
  `pcd_file_type` varchar(255) DEFAULT NULL COMMENT '2d or 3d',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_im_cache_inbox
-- ----------------------------
DROP TABLE IF EXISTS `t_im_cache_inbox`;
CREATE TABLE `t_im_cache_inbox` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `cache_msg_uuid` varchar(255) DEFAULT NULL,
  `inbox_user_uuid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_im_group_info
-- ----------------------------
DROP TABLE IF EXISTS `t_im_group_info`;
CREATE TABLE `t_im_group_info` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `group_uuid` varchar(255) DEFAULT NULL,
  `group_name` varchar(255) DEFAULT NULL,
  `created_at` varchar(255) DEFAULT NULL,
  `deleted_at` varchar(255) DEFAULT NULL,
  `member` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_im_group_member
-- ----------------------------
DROP TABLE IF EXISTS `t_im_group_member`;
CREATE TABLE `t_im_group_member` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `group_uuid` varchar(255) DEFAULT NULL,
  `member_uuid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_im_message
-- ----------------------------
DROP TABLE IF EXISTS `t_im_message`;
CREATE TABLE `t_im_message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) DEFAULT NULL,
  `msg_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `text_type_text` text,
  `file_type_uri` varchar(500) DEFAULT NULL,
  `from_user_uuid` varchar(255) DEFAULT NULL,
  `to_entity_uuid` varchar(255) DEFAULT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniquuid` (`uuid`),
  KEY `sorted_by_time` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_member
-- ----------------------------
DROP TABLE IF EXISTS `t_member`;
CREATE TABLE `t_member` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(300) DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `login_id` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_sensor_stat
-- ----------------------------
DROP TABLE IF EXISTS `t_sensor_stat`;
CREATE TABLE `t_sensor_stat` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `time` datetime DEFAULT NULL,
  `deleted_at` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_vibration_info
-- ----------------------------
DROP TABLE IF EXISTS `t_vibration_info`;
CREATE TABLE `t_vibration_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `detected_time` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `value` double DEFAULT NULL,
  `location_in_meter` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `time_index` (`created_at` DESC) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
