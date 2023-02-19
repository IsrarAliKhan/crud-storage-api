CREATE TABLE item (
    id          SERIAL      PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    price       INT,
    created_at  TIMESTAMP   NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);
