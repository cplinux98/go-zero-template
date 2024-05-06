CREATE TABLE `user` (
                            `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户id',
                            `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
                            `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
                            `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
                            `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
                            `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态，1启用、2禁用',
                            `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '头像',
                            `comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '描述',
                            `org_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所属组织',
                            `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `login_time` timestamp NULL DEFAULT NULL COMMENT '上次登录时间',
                            `password` char(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';