-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_tasks
(
    user_id    bigint REFERENCES users (id) ON DELETE CASCADE,
    task_id    bigint REFERENCES tasks (id) ON DELETE CASCADE,
    start_time timestamp,
    PRIMARY KEY (user_id, task_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_tasks
-- +goose StatementEnd
