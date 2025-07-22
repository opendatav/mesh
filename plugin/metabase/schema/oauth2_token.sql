-- @CreateToken
CREATE TABLE IF NOT EXISTS oauth2_token
(
    `code`      VARCHAR(255) NOT NULL COMMENT '授权码',
    `access`    VARCHAR(255) NOT NULL COMMENT '准入TOKEN',
    `refresh`   VARCHAR(255) NOT NULL COMMENT '刷新TOKEN',
    `data`      TEXT         NOT NULL COMMENT '补充数据',
    `create_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `expire_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    PRIMARY KEY (`code`),
    KEY         `idx_access` (`access`),
    KEY         `idx_refresh` (`refresh`),
    KEY         `idx_expire_at` (`expire_at`)
)DEFAULT CHARSET = utf8  COMMENT ='OAuth2令牌表';

-- @DeleteToken
DELETE
FROM `oauth2_token`
WHERE `code` = '?.code';

-- @IndexToken#page
SELECT *
FROM `oauth2_token`
ORDER BY `code` ASC LIMIT '?.index', '?.limit';