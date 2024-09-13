CREATE TABLE categories
(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    PRIMARY KEY (id)
);