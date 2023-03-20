CREATE TABLE alamat(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
alamat varchar(255)
);
CREATE TABLE user(
id INT NOT NULL AUTO_INCREMENT,
alamat_id INT ,
name VARCHAR(100) ,
birth_date DATE,
gender CHAR(1),
status SMALLINT,
created_at DATETIME,
updated_at DATETIME,
PRIMARY KEY(id),
FOREIGN KEY(alamat_id) REFERENCES alamat(id)
);
CREATE TABLE product_type(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(100) 
);
CREATE TABLE operator(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(100) ,
address VARCHAR(255),
phone_number VARCHAR(15)
);
CREATE TABLE product_description(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
description TEXT
);
CREATE TABLE payment_method(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(100) ,
payment_number VARCHAR(20) 
);
CREATE TABLE product(
id INT NOT NULL AUTO_INCREMENT,
product_type_id INT ,
operator_id INT ,
product_description_id INT ,
name VARCHAR(100) ,
status SMALLINT ,
created_at DATETIME ,
updated_at DATETIME ,
PRIMARY KEY(id),
FOREIGN key(product_type_id) REFERENCES product_type(id),
FOREIGN KEY(operator_id) REFERENCES operator(id),
FOREIGN KEY(product_description_id) REFERENCES product_description(id)
);
CREATE TABLE transaction(
id INT NOT NULL AUTO_INCREMENT,
user_id INT ,
payment_method_id INT ,
qty INT ,
total_price NUMERIC(21,2),
status SMALLINT,
created_at DATETIME ,
updated_at DATETIME ,
PRIMARY KEY(id),
FOREIGN KEY(user_id) REFERENCES user(id),
FOREIGN KEY(payment_method_id) REFERENCES payment_method(id)
);
CREATE TABLE transaction_detail(
id INT NOT NULL AUTO_INCREMENT,
product_id INT ,
transaction_id INT ,
qty INT ,
price NUMERIC(21,2),
PRIMARY KEY(id),
FOREIGN KEY(product_id) REFERENCES product(id),
FOREIGN KEY(transaction_id) REFERENCES transaction(id)
);