-- Active: 1698714583417@@127.0.0.1@5436@postgres@public
CREATE TABLE role
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

INSERT INTO role(name)
VALUES('admin'), ('manager'), ('client');

CREATE TABLE status 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

INSERT INTO status
VALUES('bronze'),('silver'),('gold');

CREATE TABLE order_status 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

INSERT INTO order_status 
VALUES('new'),('in progress'),('packing'),('paid'),('canceled'),('delivered'),('returned');

CREATE TABLE account
(
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    patronumic VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    FOREIGN KEY (status) REFERENCES status(name),
    FOREIGN KEY (role) REFERENCES role(name)
);

CREATE TABLE category 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

INSERT INTO category 
VALUES('cake_packaging'),
('sushi_packaging'),
('drinks_packaging'),
('bakery_packaging'),
('other_packaging');

CREATE TABLE manufacturer 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

INSERT INTO manufacturer 
VALUES('Komus'), ('Protek');

CREATE TABLE product
(
    article INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    image VARCHAR(255),
    description VARCHAR(255) NOT NULL,
    FOREIGN KEY (manufacturer) REFERENCES manufacturer(name)
    ON DELETE CASCADE,
    FOREIGN KEY (category) REFERENCES category(name)
    ON DELETE CASCADE
);

CREATE TABLE product_price 
(
    date TIMESTAMP NOT NULL,
    product_article INTEGER NOT NULL,
    price INTEGER NOT NULL,
    FOREIGN KEY (product_article) REFERENCES product(article),
    PRIMARY KEY(date, product_article)
);

CREATE TABLE address 
(
    id SERIAL PRIMARY KEY NOT NULL,
    street VARCHAR(50) NOT NULL,
    house VARCHAR(10) NOT NULL,
    apartment VARCHAR(10) NOT NULL,
    postal_code VARCHAR(10) NOT NULL,
    city VARCHAR(50) NOT NULL
);

INSERT INTO address(street, house, apartment, postal_code, city) 
VALUES('Машиностроителей', '30/2', '101', 64000, 'Курган');

CREATE TABLE warehouse
(
    id SERIAL NOT NULL PRIMARY KEY,
    address INTEGER NOT NULL UNIQUE,
    FOREIGN KEY (address) REFERENCES address(id)
);

INSERT INTO warehouse(address) VALUES(1);

CREATE TABLE product_stack 
(
    product_article INTEGER NOT NULL,
    warehouse_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    FOREIGN KEY (product_article) REFERENCES  product(article),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id),
    PRIMARY KEY(product_article, warehouse_id)
);

CREATE TABLE "order"
(
    id SERIAL NOT NULL PRIMARY KEY,
    address INTEGER NOT NULL,
    order_date TIMESTAMP NOT NULL DEFAULT current_timestamp,
    account_id INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'new',
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (address) REFERENCES address(id),
    FOREIGN KEY (status) REFERENCES order_status(name)
);

CREATE TABLE purchase 
(
    product_article INTEGER NOT NULL,
    order_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    FOREIGN KEY (product_article) REFERENCES product(article),
    FOREIGN KEY (order_id) REFERENCES "order"(id),
    PRIMARY KEY(product_article, order_id)
);

CREATE TABLE cart
(
    id INTEGER NOT NULL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id)
);

CREATE TABLE product_cart
(
    product_article INTEGER NOT NULL,
    cart_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    FOREIGN KEY (product_article) REFERENCES product(article),
    FOREIGN KEY (cart_id) REFERENCES cart(id),
    PRIMARY KEY(product_article, cart_id)
);

CREATE TABLE favorite
(
    account_id INTEGER NOT NULL,
    product_article INTEGER NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (product_article) REFERENCES product(article)
    ON DELETE CASCADE,
    PRIMARY KEY(account_id, product_article)
);
