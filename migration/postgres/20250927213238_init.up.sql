CREATE TABLE IF NOT EXISTS component
(
    id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS inventory
(
    count        INT NOT NULL CHECK (count >= 0),
    component_id INT NOT NULL,
    CONSTRAINT fk_inventory_component FOREIGN KEY (component_id) REFERENCES component (id)
);