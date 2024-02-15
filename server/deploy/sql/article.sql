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
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3) DEFAULT NULL,
    `updated_at`     datetime(3) DEFAULT NULL,
    `title`          varchar(255)  NOT NULL DEFAULT '' COMMENT '标题',
    `subtitle`       varchar(255)  NOT NULL DEFAULT '' COMMENT '副标题',
    `keywords`       varchar(255)  NOT NULL DEFAULT '' COMMENT '关键字',
    `description`    varchar(1000) NOT NULL DEFAULT '' COMMENT '文章描述',
    `author`         varchar(60)   NOT NULL DEFAULT '' COMMENT '文章作者',
    `company`        varchar(60)   NOT NULL DEFAULT '' COMMENT '参赛单位',
    `phone`          varchar(11)   NOT NULL DEFAULT '' COMMENT '手机号',
    `source`         varchar(50)   NOT NULL DEFAULT '' COMMENT '文章来源',
    `content`        text          NOT NULL COMMENT '文章内容',
    `extend`         json                   DEFAULT NULL COMMENT '附加属性',
    `thumb`          varchar(255)  NOT NULL DEFAULT '' COMMENT '封面',
    `is_link`        tinyint(4) NOT NULL DEFAULT '2' COMMENT '是否外链 2:否 1:是',
    `link`           varchar(255)  NOT NULL DEFAULT '' COMMENT 'Link',
    `type`           enum('article','page','video') NOT NULL DEFAULT 'article' COMMENT '类型：article/page/video',
    `sort_order`     tinyint(4) NOT NULL DEFAULT '0' COMMENT '排序',
    `status`         tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1:待审核 2:审核通过',
    `publish`        tinyint(4) NOT NULL DEFAULT '2' COMMENT '前台显示 1:是 2:否',
    `view_count`     smallint(6) NOT NULL DEFAULT '0' COMMENT '展示次数',
    `ticket_count`   smallint(6) NOT NULL DEFAULT '0' COMMENT '投票次数',
    `awards`         tinyint(4) NOT NULL DEFAULT '5' COMMENT '获取奖项 5:无奖项 1:一等奖 2:二等奖 3:三等奖 4:特等奖',
    `top`            tinyint(4) NOT NULL DEFAULT '2' COMMENT '置顶 1:是 2:否',
    `auditor_id`     mediumint(9) NOT NULL DEFAULT '0' COMMENT '审核人',
    `auditor_failed` varchar(1000) NOT NULL DEFAULT '' COMMENT '审核结果',
    `created_by`     bigint(20) unsigned DEFAULT NULL COMMENT '创建者',
    `updated_by`     bigint(20) unsigned DEFAULT NULL COMMENT '更新者',
    `deleted_by`     bigint(20) unsigned DEFAULT NULL COMMENT '删除者',
    `category_id`    smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '所属分类',
    `deleted_at`     datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY              `idx_article_deleted_at` (`deleted_at`) USING BTREE,
    KEY              `created_by` (`created_by`),
    KEY              `type` (`type`),
    KEY              `publish` (`publish`),
    KEY              `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
1. 创意视频类
2. 非遗文创类
3. 文艺表演类
4. 宣传阵地类
投票数量


特等奖  四类视频各1人
一等奖  四类视频各3人
二等奖  四类视频各10人
三等奖  四类视频各20人
优秀奖
，宣传阵地类奖项
、优秀作品奖

