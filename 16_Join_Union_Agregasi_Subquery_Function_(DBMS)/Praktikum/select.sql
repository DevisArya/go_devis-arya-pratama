SELECT * FROM user WHERE gender = "L";

SELECT * FROM product WHERE id = 3;

SELECT * FROM user WHERE created_at >= DATE(NOW()-INTERVAL 7 DAY);

SELECT COUNT(1) jumlah_user_perempuan FROM user WHERE gender="P";

SELECT * FROM user ORDER BY name ASC;

SELECT * FROM product LIMIT 5;