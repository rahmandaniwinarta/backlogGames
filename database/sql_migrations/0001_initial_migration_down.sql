-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS user CASCADE;
DROP TABLE IF EXISTS games CASCADE;
DROP TABLE IF EXISTS transaction CASCADE;
DROP TABLE IF EXISTS genre CASCADE;

-- +migrate StatementEnd