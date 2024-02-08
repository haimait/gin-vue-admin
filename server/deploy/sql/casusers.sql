/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : fofa_cert_init

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 08/02/2024 11:23:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casusers
-- ----------------------------
DROP TABLE IF EXISTS `casusers`;
CREATE TABLE `casusers`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名称',
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
  `phone_number` char(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户email',
`state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '用户状态，1-启用，2-关闭',
  `encrypted_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '登陆密码',
  `sign_in_count` int(11) NULL DEFAULT NULL COMMENT '登陆次数',
  `last_sign_in_at` datetime(0) NULL DEFAULT NULL COMMENT '登后登陆时间',
  `last_sign_in_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '最后登陆ip',
  `description` varchar(1000) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_casusers_on_username_and_key`(`username`) USING BTREE,
  INDEX `index_casusers_on_email_and_key`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 131 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = DYNAMIC;

