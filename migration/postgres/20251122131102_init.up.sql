BEGIN;

CREATE TABLE IF NOT EXISTS category
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS component
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    price       INT  NOT NULL,
    category_id INT  NOT NULL REFERENCES category(id),
    description TEXT NOT NULL
);

COMMIT;