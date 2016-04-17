-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
  `id`                INT UNSIGNED NOT NULL AUTO_INCREMENT
  COMMENT 'ID',
  `name`              VARCHAR(255)
  COMMENT 'ユーザー名',
  `email`             VARCHAR(255) NOT NULL
  COMMENT 'メールアドレス',
  `session_token`     VARCHAR(255) NOT NULL
  COMMENT 'セッショントークン',
  `token_expire`      TIMESTAMP    NOT NULL
  COMMENT 'トークン有効期限',
  `last_login_at`     TIMESTAMP    NOT NULL
  COMMENT '最終ログイン日時',
  `created_at`        TIMESTAMP    NOT NULL
  COMMENT '作成日時',
  `updated_at`        TIMESTAMP    NOT NULL
  COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE (`session_token`)
)
  ENGINE = InnoDB
  CHARSET = utf8;

CREATE TABLE auths (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT
  COMMENT 'ID',
  `user_id`    INT UNSIGNED NOT NULL
  COMMENT 'ユーザーID',
  `source`     VARCHAR(255) NOT NULL
  COMMENT '認証元',
  `source_id`  VARCHAR(255) NOT NULL
  COMMENT '認証元ユーザーID',
  `email`      VARCHAR(255) NOT NULL
  COMMENT 'メールアドレス',
  `created_at` TIMESTAMP    NOT NULL
  COMMENT '作成日時',
  `updated_at` TIMESTAMP    NOT NULL
  COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE (`user_id`, `source`)
)
  ENGINE = InnoDB
  CHARSET = utf8;


-- INSERT INTO `users`
-- (`created_at`, `updated_at`, `name`)
-- VALUES
--   (NOW(), NOW(), '田中'),
--   (NOW(), NOW(), '鈴木');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `auths`;
