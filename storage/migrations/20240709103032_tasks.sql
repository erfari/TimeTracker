-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks
(
    id               BIGSERIAL PRIMARY KEY,
    task_name        VARCHAR(255) NOT NULL,
    assignee_name    VARCHAR(255),
    assignee_user_id BIGINT      ,
    created_at       TIMESTAMP    NOT NULL DEFAULT NOW(),
    start_task       TIMESTAMP,
    end_task         TIMESTAMP
);

ALTER TABLE tasks
    ADD CONSTRAINT fk_users
        FOREIGN KEY (assignee_user_id)
            REFERENCES users (id)
            ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd