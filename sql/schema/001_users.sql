-- +goose Up
CREATE TABLE users(
                      id int PRIMARY KEY NOT NULL,
                      created_at TIMESTAMP DEFAULT NOW(),
                      updated_at TIMESTAMP DEFAULT NOW(),
    slug TEXT NOT NULL,
    name TEXT NOT NULL,
    agent_id TEXT,
    visit_count TEXT,
    has_called BOOL
);

-- +goose Down
DROP TABLE users;