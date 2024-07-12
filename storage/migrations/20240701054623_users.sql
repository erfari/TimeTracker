-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id                     BIGSERIAL PRIMARY KEY,
    name                   VARCHAR(255) NOT NULL,
    surname                VARCHAR(255) NOT NULL,
    patronymic             VARCHAR(255)          DEFAULT NULL,
    address                VARCHAR(255) NOT NULL,
    created_at             TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_documents;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
