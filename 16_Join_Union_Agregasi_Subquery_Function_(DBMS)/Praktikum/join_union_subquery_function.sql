SELECT * FROM transaction WHERE user_id = 1
UNION ALL
SELECT * FROM transaction WHERE user_id = 2;

SELECT *
FROM 
(SELECT user_id, SUM(total_price) jumlah_harga_transaksi FROM transaction GROUP BY user_id ) AS total_transaksi
WHERE user_id = 1;

SELECT COUNT(1) total_transaksi FROM (SELECT td.id, td.product_id, td.transaction_id,  p.product_type_id FROM product p INNER JOIN transaction_detail td ON p.id = td.product_id) ptd 
WHERE product_type_id = 2 GROUP BY product_type_id;

SELECT p.*, pt.name product_type_name  FROM product p INNER JOIN product_type pt ON p.product_type_id = pt.id;

SELECT a.*, b.user_name FROM (SELECT td.*,  p.name product_name FROM product p INNER JOIN transaction_detail td ON p.id = td.product_id) a 
INNER JOIN 
(SELECT t.*, u.name user_name FROM transaction t INNER JOIN user u ON u.id = t.user_id) b 
ON 
a.transaction_id = b.id;

DELIMITER $$
CREATE TRIGGER delete_transaction_detail 
    AFTER DELETE 
    ON transaction
    FOR EACH ROW 
BEGIN
    DELETE FROM transaction_detail WHERE transaction_id = old.id;
END$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER update_qty 
    AFTER DELETE 
    ON transaction_detail
    FOR EACH ROW 
BEGIN
    UPDATE transaction 
    SET qty = qty - old.qty
    WHERE old.transaction_id = id;
END$$
DELIMITER ;

SELECT * FROM product
WHERE
id NOT IN 
(SELECT DISTINCT transaction_id
FROM transaction_detail);

