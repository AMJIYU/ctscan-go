CREATE TABLE `user_wxmini` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `openid` varchar(100) NOT NULL COMMENT '微信openid',
    `unionid` varchar(100) DEFAULT NULL COMMENT '微信unionid',
    `nickname` varchar(50) DEFAULT NULL COMMENT '微信昵称',
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
    `user_id` bigint(20) NOT NULL COMMENT '关联的用户ID',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信小程序用户表';