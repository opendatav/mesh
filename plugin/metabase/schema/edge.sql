-- @CreateEdge
CREATE TABLE IF NOT EXISTS `edge`
(
    `node_id`     varchar(255)  NOT NULL DEFAULT '' COMMENT '节点编号',
    `inst_id`     varchar(255)  NOT NULL DEFAULT '' COMMENT '机构编号',
    `inst_name`   varchar(255)  NOT NULL DEFAULT '' COMMENT '机构名称',
    `address`     varchar(4096) NOT NULL DEFAULT '' COMMENT '节点地址',
    `describe`    varchar(4096) NOT NULL DEFAULT '' COMMENT '节点说明',
    `certificate` text          NOT NULL COMMENT '节点证书',
    `status`      int(11)       NOT NULL DEFAULT '0' COMMENT '状态',
    `version`     int(11)       NOT NULL DEFAULT '0' COMMENT '乐观锁版本',
    `auth_code`   varchar(4096) NOT NULL DEFAULT '' COMMENT '授权码',
    `static_ip`   varchar(4096) NOT NULL DEFAULT '' COMMENT '静态IP',
    `public_ip`   varchar(4096) NOT NULL DEFAULT '' COMMENT '公网IP',
    `requests`   int NOT NULL DEFAULT '' COMMENT '请求并发限制',
    `extra`       varchar(4096) NOT NULL DEFAULT '' COMMENT '补充信息',
    `expire_at`   datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    `create_at`   datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`   datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `create_by`   varchar(255)  NOT NULL DEFAULT '' COMMENT '创建人',
    `update_by`   varchar(255)  NOT NULL DEFAULT '' COMMENT '更新人',
    `group`       varchar(255)  NOT NULL DEFAULT '' COMMENT '联盟中心节点机构id-多个用逗号分割',
    PRIMARY KEY `uk_node_id` (`node_id`),
    KEY           `idx_inst_id` (`inst_id`),
    KEY           `idx_inst_name` (`inst_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 10
  DEFAULT CHARSET = utf8 COMMENT ='网络节点联通表，两点之间最多一条边';

-- @DeleteEdge
DELETE
FROM `edge`
WHERE `node_id` = '?.node_id';

-- @IndexEdge#page
SELECT *
FROM `edge`
ORDER BY `node_id` ASC LIMIT '?.index', '?.limit';


-- @SelectByIds#one
SELECT *
FROM `edge`
WHERE `node_id` = '?.node_id' OR `inst_id` = '?.inst_id' LIMIT 1;

-- @SelectAll#many
SELECT *
FROM `edge`;
