-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pg_stat_statements;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION pg_stat_statements;
-- +goose StatementEnd
