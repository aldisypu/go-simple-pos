CREATE TABLE products
(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(255),
    price REAL NOT NULL CHECK (price > 0),
    stock INTEGER NOT NULL CHECK (stock >= 0),
    category_id VARCHAR(100) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);