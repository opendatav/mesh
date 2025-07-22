-- @Create
CREATE TABLE IF NOT EXISTS `route`
(
    `name`     varchar(64) NOT NULL COMMENT '名称',
    `listen`   varchar(2048) NOT NULL COMMENT '监听',
    `matcher`   varchar(2048) NOT NULL COMMENT '正则',
    `backend`   varchar(2048) NOT NULL COMMENT '服务',
    `priority` int NOT NULL DEFAULT 0 COMMENT '优先级',
    `pass_host_header` int NOT NULL DEFAULT 0 COMMENT '是否透传头',
    `create_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`name`)
    ) DEFAULT CHARSET = utf8mb4 COMMENT ='路由表';

-- @Index#page
SELECT *
FROM `route`
ORDER BY `name` ASC
    LIMIT '?.index', '?.limit';

-- @All#many
SELECT *
FROM `route`;
