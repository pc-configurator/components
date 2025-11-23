BEGIN;

CREATE TABLE IF NOT EXISTS category
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS component
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    price       INT  NOT NULL,
    category_id INT  NOT NULL REFERENCES category(id),
    description TEXT NOT NULL
);

COMMIT;