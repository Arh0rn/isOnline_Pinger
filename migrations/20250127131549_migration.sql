-- +goose Up
-- +goose StatementBegin

CREATE TABLE urls (
                      id SERIAL PRIMARY KEY,
                      url TEXT NOT NULL
);
CREATE TABLE parameters (
                            id SERIAL PRIMARY KEY,
                            timeout INTEGER NOT NULL,
                            interval INTEGER NOT NULL,
                            workers INTEGER NOT NULL
);


INSERT INTO parameters (timeout, interval, workers) VALUES (10, 5, 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS parameters;
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
