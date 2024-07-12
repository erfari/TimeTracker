-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_documents
(
    id                     BIGSERIAL PRIMARY KEY,
    user_id                BIGINT NOT NULL,
    passport_serial_number VARCHAR(255) NOT NULL,
    passport_number        VARCHAR(255) NOT NULL
);

ALTER TABLE user_documents
ADD CONSTRAINT fk_users
FOREIGN KEY (user_id)
REFERENCES users (id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_documents;
-- +goose StatementEnd
