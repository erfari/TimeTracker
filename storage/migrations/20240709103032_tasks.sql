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

INSERT INTO tasks(task_name, assignee_name, assignee_user_id, start_task, end_task)
values ('task1', 'test1', 1, '2024-07-09 11:13:10.000000', '2024-07-09 13:13:10.000000');

INSERT INTO tasks(task_name, assignee_name, assignee_user_id, start_task, end_task)
values ('task1', 'test1', 1, '2024-07-09 11:13:10.000000', '2024-07-09 13:02:10.000000');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd