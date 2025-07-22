-- @Create
CREATE TABLE IF NOT EXISTS `mesh_kv`
(
    `key`       varchar(255) NOT NULL COMMENT '配置KEY',
    `value`     text         NOT NULL COMMENT '配置内容',
    `create_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建人',
    `update_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新人',
    PRIMARY KEY (`key`)
    ) DEFAULT CHARSET = utf8mb4 COMMENT ='KV表';

-- @Index#page
SELECT *
FROM `mesh_kv`
ORDER BY `key` ASC
    LIMIT '?.index', '?.limit';

-- @SelectKeys#many
SELECT `key`
FROM `mesh_kv`
WHERE `key` LIKE '%?.key%';