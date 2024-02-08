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
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casusers
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT,
    `created_at`      datetime(0) NOT NULL COMMENT '创建时间',
    `updated_at`      datetime(0) NOT NULL COMMENT '更新时间',
    `title`           varchar(255) NOT NULL COMMENT '标题',
    `alias`         varchar(150) DEFAULT '' COMMENT '别名',
    `subtitle`      varchar(150) DEFAULT '' COMMENT '副标题',
    `keywords`      varchar(150) DEFAULT '' COMMENT '关键字',
    `description`   varchar(255) DEFAULT '' COMMENT '文章描述',
    `author`        varchar(60)  DEFAULT NULL COMMENT '文章作者',
    `company`        varchar(60)  DEFAULT NULL COMMENT '参赛单位',
    `phone`           varchar(11)  NOT NULL COMMENT '手机号',
    `source`        varchar(50)  DEFAULT NULL COMMENT '文章来源',
    `content`       text         NOT NULL COMMENT '文章内容',
    `attribute`     json         DEFAULT NULL COMMENT '附加属性',
    `thumb`         varchar(255) DEFAULT NULL COMMENT '封面',
    `is_link`       enum('1','2') NOT NULL DEFAULT '1' COMMENT 'isLink',
    `link`          varchar(255) DEFAULT NULL default '' COMMENT 'Link',
    `type`          enum('article','page','video') DEFAULT 'article' COMMENT '类型：article/page/video',
    `order`         int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `status`          int(1) DEFAULT '1' COMMENT '状态 1:待审核 2:审核通过',
    `publish`    int(1) DEFAULT '2' COMMENT '前台显示 1:是 2:否',
    `view_count`   int(11) DEFAULT '0' COMMENT '展示次数',
    `ticket_count`   int(11) DEFAULT '0' COMMENT '投票次数',
    `awards`    int(1) DEFAULT '0' COMMENT '获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖',
    `top`           enum('1','2') NOT NULL DEFAULT '2' COMMENT '置顶 1:是 2:否',
    `auditor_id`         int(11) NOT NULL DEFAULT '0' COMMENT '审核人',
    `auditor_result`  varchar(255) DEFAULT NULL default '' COMMENT '审核结果',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX             `index_casusers_on_username_and_key`(`username`) USING BTREE,
    INDEX             `index_casusers_on_email_and_key`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 131 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = DYNAMIC;


1. 创意视频类
2. 非遗文创类
3. 文艺表演类
4. 宣传阵地类
投票数量


特等奖  四类视频各1人
一等奖  四类视频各3人
二等奖  四类视频各10人
三等奖  四类视频各20人
优秀奖，宣传阵地类奖项、优秀作品奖

