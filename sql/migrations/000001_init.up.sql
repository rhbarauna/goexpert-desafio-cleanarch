CREATE TABLE orders (
    id varchar(36) NOT NULL PRIMARY KEY,
    price FLOAT,
    tax FLOAT,
    final_price FLOAT
);
