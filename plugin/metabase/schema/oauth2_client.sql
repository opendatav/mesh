-- @CreateClient
CREATE TABLE IF NOT EXISTS oauth2_client
(
    `id`     VARCHAR(255) NOT NULL COMMENT '客户端ID',
    `name`   VARCHAR(255) NOT NULL COMMENT '客户端名称',
    `secret` VARCHAR(255) NOT NULL COMMENT '客户端密钥',
    `domain` VARCHAR(255) NOT NULL COMMENT '客户端域名',
    `data`   TEXT         NOT NULL COMMENT '补充数据',
    PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8 COMMENT ='OAuth2客户端表';

-- @DeleteClient
DELETE
FROM oauth2_client
WHERE `id` = '?.id';

-- @IndexClient#many
SELECT *
FROM `oauth2_client`
ORDER BY `id` ASC LIMIT '?.index', '?.limit';