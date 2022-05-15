-- TODO разобравться с несоответствием uint64 моделек в ID и типом BIGSERIAL в БД
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
