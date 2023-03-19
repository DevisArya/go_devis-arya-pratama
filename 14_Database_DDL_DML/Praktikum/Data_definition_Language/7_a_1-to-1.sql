CREATE TABLE payment_method_description(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
description TEXT
);
ALTER TABLE payment_method
ADD payment_method_description_id INT 
AFTER id;
