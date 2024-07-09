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

WITH ids AS(
    INSERT INTO users(name, surname, patronymic, address)
        VALUES('test1', 'test1', 'test1', 'test1 street')
        RETURNING id
)
INSERT INTO user_documents(user_id, passport_number, passport_serial_number)
SELECT id, '5555', '777777' FROM ids;

WITH ids AS(
    INSERT INTO users(name, surname, patronymic, address)
        VALUES('test2', 'test2', 'test2', 'test2 street')
        RETURNING id
)
INSERT INTO user_documents(user_id, passport_number, passport_serial_number)
SELECT id, '1432', '555777' FROM ids;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_documents;
-- +goose StatementEnd
