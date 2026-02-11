-- +goose Up
ALTER TABLE users
    ALTER COLUMN visit_count TYPE INTEGER USING visit_count::INTEGER,
    ALTER COLUMN visit_count SET DEFAULT 0;

-- +goose Down
ALTER TABLE users
    ALTER COLUMN visit_count TYPE TEXT,
    ALTER COLUMN visit_count DROP DEFAULT;
