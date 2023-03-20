CREATE TABLE user_payment_method_detail(
id INT NOT NULL AUTO_INCREMENT,
user_id INT ,
payment_method_id INT ,
PRIMARY KEY(id),
FOREIGN KEY(user_id) REFERENCES user(id),
FOREIGN KEY(payment_method_id) REFERENCES payment_method(id)
);