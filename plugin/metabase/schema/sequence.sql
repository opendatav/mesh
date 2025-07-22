-- @Create
CREATE TABLE IF NOT EXISTS `mesh_seq`
(
    `kind`      varchar(255) NOT NULL DEFAULT '' COMMENT '序列号类型',
    `min`       bigint(20)   NOT NULL DEFAULT '0' COMMENT '当前范围最小值',
    `max`       bigint(20)   NOT NULL DEFAULT '0' COMMENT '当前范围最大值',
    `size`      int(11)      NOT NULL DEFAULT '0' COMMENT '每次取号段大小',
    `length`    int(11)      NOT NULL DEFAULT '0' COMMENT '序列号长度不足补零',
    `status`    int(11)      NOT NULL DEFAULT '0' COMMENT '状态',
    `version`   int(11)      NOT NULL DEFAULT '0' COMMENT '乐观锁版本',
    `create_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`kind`)
    ) DEFAULT CHARSET = utf8mb4 COMMENT ='序列号表';

-- @SetMin
UPDATE `mesh_seq`
SET `min`     = '?.min',
    `version` = `version` + 1
WHERE `kind` = '?.kind'
  AND `version` = '?.version';

-- @SelectByKindForUpdate#one
SELECT *
FROM `mesh_seq`
WHERE `kind` = '?.kind' FOR
    UPDATE;