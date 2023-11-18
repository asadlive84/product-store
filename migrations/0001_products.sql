-- +goose Up
-- SQL in this section is executed when the migration is applied.


CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    status_id INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    parent_id INT,
    sequence INT,
    status_id INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE suppliers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20),
    status_id INT DEFAULT 1,
    is_verified_supplier BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    brand_id INT REFERENCES brands(id),
    category_id INT REFERENCES categories(id),
    supplier_id INT REFERENCES suppliers(id),
    unit_price DECIMAL(10, 2) NOT NULL,
    discount_price DECIMAL(10, 2),
    tags VARCHAR(255),
    status_id INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_stocks (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    stock_quantity INT CHECK (stock_quantity >= 0),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE brands;
DROP TABLE categories;
DROP TABLE suppliers;
DROP TABLE products;
DROP TABLE product_stocks;
