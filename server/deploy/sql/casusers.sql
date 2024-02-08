CREATE TABLE `casusers`
(
    `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`      datetime(3) DEFAULT NULL,
    `updated_at`      datetime(3) DEFAULT NULL,
    `deleted_at`      datetime(3) DEFAULT NULL,
    `uuid`            varchar(191) DEFAULT NULL COMMENT '用户UUID',
    `username`        varchar(255) NOT NULL COMMENT '用户名',
    `nickname`        varchar(255) DEFAULT '' COMMENT '昵称',
    `email`           varchar(255) NOT NULL COMMENT 'Email',
    `phone`           varchar(11)  NOT NULL COMMENT '手机号',
    `enable`          int(1) DEFAULT '1' COMMENT '用户是否被冻结 1:正常 2:冻结',
    `password`        varchar(255) DEFAULT '' COMMENT '用户登录密码',
    `sign_in_count`   int(11) DEFAULT '0' COMMENT '登录次数',
    `last_sign_in_at` datetime     DEFAULT NULL COMMENT '最近登录时间',
    `last_sign_in_ip` varchar(255) DEFAULT '' COMMENT '最近登录IP',
    `description`     varchar(255) DEFAULT '' COMMENT '用户说明',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casusers_username` (`username`),
    UNIQUE KEY `idx_casusers_email` (`email`),
    UNIQUE KEY `idx_casusers_phone` (`phone`),
    UNIQUE KEY `idx_casusers_uuid` (`uuid`),
    KEY               `idx_casusers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;