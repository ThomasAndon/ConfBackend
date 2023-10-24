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

 Date: 24/10/2023 15:32:49
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
-- Records of t_hero_pcd_uoload
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of t_im_cache_inbox
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of t_im_group_info
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of t_im_group_member
-- ----------------------------
BEGIN;
COMMIT;

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
-- Records of t_im_message
-- ----------------------------
BEGIN;
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (19, 'mid-f636a8ed-7114-417b-be79-aac99c3c61e6', 'text', '', '', '12345', '67890', '2023-05-29 20:49:34.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (20, 'mid-aa4a752d-680a-4466-b6dc-e823eeaa1cd4', 'text', '', '', '12345', '67890', '2023-05-29 20:55:28.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (21, 'mid-a7c7b1b5-5d48-4356-b4e9-ae0c9e196ace', 'text', '', '', '12345', '67890', '2023-05-29 20:55:39.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (22, 'mid-8d172d6f-e49b-43ec-85cb-3f670caebd69', 'text', '', '', '12345', '67890', '2023-05-29 20:55:40.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (23, 'mid-82cd7a49-8501-448e-a6b6-d0745e3a5345', 'text', '', '', '12345', '67890', '2023-05-29 20:55:41.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (24, 'mid-d0c53e2b-473b-420d-a06a-e6defbf39db7', 'text', '', '', '12345', '67890', '2023-05-29 20:55:42.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (25, 'mid-004439de-c265-4119-a6f8-0a76f4c2fc8f', 'text', '', '', '12345', '67890', '2023-05-29 20:55:43.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (26, 'mid-1b6be811-8267-425b-999c-745867dbf265', 'text', '', '', '12345', '67890', '2023-05-29 20:57:33.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (27, 'mid-9521a0c4-dbf3-49f3-8575-7bd3462631e3', 'text', '', '', '12345', '67890', '2023-05-29 20:57:34.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (28, 'mid-7a1abf2f-c63d-4d1b-a0fd-4531ab41b383', 'text', '', '', '12345', '67890', '2023-05-29 20:59:08.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (29, 'mid-f07852ac-df5c-42b0-926f-1aeb3f2ec728', 'text', '', '', '12345', '67890', '2023-05-29 21:09:13.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (30, 'mid-0bb864f8-1aeb-43e2-9fc3-d502b1c56205', 'text', '', '', '12345', '67890', '2023-05-29 21:36:27.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (31, 'mid-f5455ec0-0aa1-4169-bff6-b6c8303c1f64', 'text', '', '', '12345', '67890', '2023-05-29 21:36:48.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (32, 'mid-aa62af1e-411e-43f3-8e48-a967f1233f9a', 'text', '', '', '12345', '67890', '2023-05-29 21:56:38.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (33, 'mid-f28c6ce6-8101-4aee-b838-4e60a213d1af', 'text', '', '', '12345', '67890', '2023-05-29 21:57:17.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (34, 'mid-c0673f31-ac06-46d2-ae12-89f2644a7ada', 'text', '', '', '1234567890', '123456789', '2023-05-30 17:30:15.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (35, 'mid-a48e2867-f6b5-40e8-8383-47a546fe5b5d', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:02:50.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (36, 'mid-7cb7964f-383a-4d1a-bea6-88357aee8369', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:02:53.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (37, 'mid-ab0b7737-25bc-4733-b95b-ffd90b03e1a1', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:02:54.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (38, 'mid-91ce416a-21ad-4ad5-9a96-02b0363e90c5', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:05.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (39, 'mid-39b5e86d-ace5-4fd5-a0cc-c7aa3e40b74b', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:08.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (40, 'mid-21e90afe-71fa-4969-87ff-362fdbe839cd', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:11.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (41, 'mid-ba690f09-2a83-4657-9369-f4c4e1f16b07', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:14.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (42, 'mid-4a4a683b-8d2b-4908-aef2-f331d29de50f', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:17.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (43, 'mid-7ccea51a-6fcf-4129-b0f2-87d95eaab677', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:21.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (44, 'mid-5dcadb07-e02d-411f-bc1b-eee026504c53', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:25.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (45, 'mid-9064fd21-0ca8-4da2-bad3-da78b8d48ea7', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:28.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (46, 'mid-4c9bb9bd-3d61-49f8-baf2-3dcea57a2812', 'text', '', '', '1234567890', '123456789', '2023-05-30 18:03:32.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (47, 'mid-lalalalalalalalala', 'text', '', NULL, '123456789', '1234567890', '2023-05-30 20:52:56.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (48, 'mid-93dbda75-5345-4201-9c29-9a41698ae38d', 'text', '', '', '1234567890', '123456789', '2023-05-30 21:42:38.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (49, 'mid-09e66339-df24-4489-b6ee-c77f93641394', 'text', '', '', '1234567890', '123456789', '2023-05-30 21:43:08.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (63, 'mid-1c395a57-40a8-4529-83be-ab68b88aff3e', 'text', '', '', '1234567890', '123456789', '2023-06-01 11:38:22.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (64, 'mid-551dabe7-7698-498f-bd1f-ac91b426be2f', 'text', '', '', '1234567890', '123456789', '2023-06-01 11:38:23.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (65, 'mid-07121b44-8a3a-49c9-b783-66c0c6d7113d', 'text', '', '', '1234567890', '123456789', '2023-06-01 20:08:02.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (66, 'mid-8e910065-a752-4160-b688-a84f2bbe6fcf', 'image', '', '566f74b4-f790-4c9e-9241-6bf3384c36f6.jpg', '1234567890', '123456789', '2023-06-01 20:08:14.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (67, 'mid-ec10ca1f-9bb3-4f68-8af5-8dfa44cae205', 'text', '', '', '1234567890', '123456789', '2023-06-01 20:34:23.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (68, 'mid-8c9796ec-ce6c-4fe0-95e0-d12891e208e1', 'text', '', '', '1234567890', '123456789', '2023-06-01 20:49:04.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (69, 'mid-34ca8036-ae5f-4f85-bfd3-623b5933e64f', 'image', '', '8a3cbb78-c9bd-47e9-842c-c275b759486f.jpg', '1234567890', '123456789', '2023-06-01 20:51:05.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (70, 'mid-342423423-wdgwrgwew-23423523tg-2g', 'text', '', NULL, '1234567890', '123456789', '2023-06-02 12:44:05.000000');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (71, 'mid-4f4afc63-de86-4512-a1f5-769d359184fc', 'text', '', '', '1234567890', '123456789', '2023-06-02 12:44:26.428832');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (72, 'mid-885a4ab0-7343-40e6-84e9-c227e62a0465', 'text', '', '', '1234567890', '123456789', '2023-06-02 12:44:29.874620');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (73, 'mid-982aa2c8-bfc4-4a44-9a56-9592971b9ffb', 'text', '', '', '1234567890', '123456789', '2023-06-03 16:38:47.380554');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (74, 'mid-785e688e-70cf-4952-9859-d918c34651ed', 'text', '', '', '1234567890', '123456789', '2023-06-03 16:39:03.691122');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (75, 'mid-5102a389-9562-407e-a544-f3a1ad33fece', 'text', '', '', '1234567890', '123456789', '2023-06-03 16:40:31.946068');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (76, 'mid-b29d48e2-3bad-4c80-9356-59063766807c', 'image', '', '3a47d242-3573-455c-94dd-0a04cdcb980a.jpg', '1234567890', '123456789', '2023-06-08 16:31:56.394102');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (77, 'mid-15c6bda2-a690-4b56-af57-004d4c67a4ad', 'image', '', 'ae043f22-ddf3-491f-aef4-61692df00a8a.jpg', '1234567890', '123456789', '2023-06-12 16:04:01.553189');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (78, 'mid-8f50c27c-958e-4daa-9a6f-8fdd32ea3524', 'audio', '', '0b741fa2-cc2d-4f96-ba54-9ec15af17294.wav', 't01', 't02', '2023-06-14 19:59:30.929453');
INSERT INTO `t_im_message` (`id`, `uuid`, `msg_type`, `text_type_text`, `file_type_uri`, `from_user_uuid`, `to_entity_uuid`, `created_at`) VALUES (79, 'mid-2555ef87-ca43-4604-9367-1addecd2abf1', 'image', '', '205e2759-698e-469a-a6cc-df1784f6fb38.jpg', '1234567890', '123456789', '2023-08-09 11:40:38.853309');
COMMIT;

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
-- Records of t_member
-- ----------------------------
BEGIN;
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (1, '1234567890', 'a', NULL, '1234567890', '1234', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (2, '123456789', 'a', NULL, '123456789', '1234', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (3, '987654321', 'a', NULL, NULL, NULL, '2023-05-30 21:52:50');
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (4, '18bd3f28-f2ca-8522-fefd-25ef4f359c14', 'abc', NULL, NULL, NULL, NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (5, 't01', 'nickname-t01', NULL, 'test01', 'test01', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (6, 't02', 'nickname-t02', NULL, 'test02', 'test02', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (7, 't001', 'nickname-t001', NULL, 'test001', 'test001', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (8, 't002', 'nickname-t002', NULL, 'test002', 'test002', NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (9, '4026621b-7e77-6e09-a16f-ac269464c988', 'a', NULL, NULL, NULL, NULL);
INSERT INTO `t_member` (`id`, `uuid`, `nickname`, `created_at`, `login_id`, `password`, `deleted_at`) VALUES (10, '03d87a72-b133-eeec-ffd9-6e81bf8b5b88', 'a', NULL, NULL, NULL, NULL);
COMMIT;

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
-- Records of t_sensor_stat
-- ----------------------------
BEGIN;
COMMIT;

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of t_vibration_info
-- ----------------------------
BEGIN;
INSERT INTO `t_vibration_info` (`id`, `detected_time`, `created_at`, `value`, `location_in_meter`) VALUES (3, '2023-09-18 12:34:56', '2023-09-18 18:47:37', 12.3, 45.6);
INSERT INTO `t_vibration_info` (`id`, `detected_time`, `created_at`, `value`, `location_in_meter`) VALUES (4, '2023-09-18 12:34:56', '2023-09-18 18:47:37', 23.4, 56.7);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
