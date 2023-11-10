-- Active: 1698714583417@@127.0.0.1@5436@postgres@public
CREATE TABLE role
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE status 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE order_status 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

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

CREATE TABLE manufacturer 
(
    name VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE product
(
    article INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    manufacturer VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    image VARCHAR(255),
    description VARCHAR(255) NOT NULL,
    FOREIGN KEY (manufacturer) REFERENCES manufacturer(name),
    FOREIGN KEY (category) REFERENCES category(name) 
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

CREATE TABLE warehouse
(
    id SERIAL NOT NULL PRIMARY KEY,
    address INTEGER NOT NULL,
    FOREIGN KEY (address) REFERENCES address(id)
);

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
    order_date TIMESTAMP NOT NULL,
    account_id INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL,
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