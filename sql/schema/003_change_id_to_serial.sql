-- +goose Up
CREATE SEQUENCE users_id_seq OWNED BY users.id;
SELECT setval('users_id_seq', COALESCE(MAX(id), 0) + 1) FROM users;
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('users_id_seq');

-- +goose Down
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE users_id_seq;
