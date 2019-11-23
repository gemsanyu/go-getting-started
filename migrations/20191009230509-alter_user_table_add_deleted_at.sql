
-- +migrate Up
ALTER TABLE `users`
    ADD COLUMN `deleted_at` datetime default null;

-- +migrate Down
ALTER TABLE `users`
    DROP COLUMN `deleted_at`;
