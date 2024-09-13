CREATE TABLE sales
(
    id VARCHAR(100) NOT NULL,
    sale_date BIGINT NOT NULL,
    total REAL NOT NULL CHECK (total >= 0),
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    PRIMARY KEY (id)
);