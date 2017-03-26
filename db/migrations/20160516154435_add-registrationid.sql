
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `users` ADD `registration_id` TEXT COMMENT 'registration id' AFTER `token_expire`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `users` DROP `registration_id`;
