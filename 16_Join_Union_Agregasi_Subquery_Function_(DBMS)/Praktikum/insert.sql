INSERT INTO operator(name, address, phone_number) 
VALUES 
('ela','Bandung','081234567892'),
('robi','Bandung','085484529594'),
('bryan','jakarta','0854844052565'),
('adam','malang','0854845258152'),
('elis','jombang','087548522595');

INSERT INTO product_type(name)
 VALUES 
 ('elektronik'),
 ('perabotan rumah tangga'),
 ('fashion');
 
 INSERT INTO product_description(description)
VALUES
("Cocok digunakan untuk semua iPhone yg sudah support penggunaan charger fast charging 20W 8+ / X / XS / XS MAX / 11 / 11 PRO / 11 PRO MAX / 12 / 12 MINI / 12 PRO / 12 PRO MAX / 13 / 13 MINI / 13 PRO / 13 PRO MAX / 14 / 14 PLUS / 14 PRO / 14 PRO MAX

Spesifikasi Adaptor :
- Model : CD137
- Input : 100-240V~50/60Hz 500mA Max
- USB C Output : 5V-3A / 9V-2.22A / 12V-1.67A / 3.3-5.9V-3A / 3.3-11V-1.8A
- Total Output : 20W Max
- Warna: PUTIH / HITAM"),
("Cocok digunakan untuk semua iPhone yg sudah support penggunaan charger fast charging 20W 8+ / X / XS / XS MAX / 11 / 11 PRO / 11 PRO MAX / 12 / 12 MINI / 12 PRO / 12 PRO MAX / 13 / 13 MINI / 13 PRO / 13 PRO MAX / 14 / 14 PLUS / 14 PRO / 14 PRO MAX

Spesifikasi Kabel :
- Length : 1m
- Current : 3A
- MFi Certificate
- Warna : PUTIH / HITAM"),
("Milk Pot with Steamer 
Diameter 20 cm
Panci susu serba guna yang bisa anda gunakan untuk memasak makanan seperti mie , soup , dan sekaligus bisa dipakai untuk mengukus makanan.
Terbuat dari bahan stainless steel berkualitas
tutup kaca"),
("Penggorengan bahan teflon
Ukuran pan 10X16cm
Anti lengket
Anti gores
Bisa untuk kompor konvensional
Handle tahan panas
Nyaman digenggam dan mudah dibersihkan
Type fry pan ini cocok untuk menggoreng telur (Tamagoyaki)"),
("Panci Mie / Panci masak 18 cm pegangan kayu
Bahan Almunium Tebal 2mm
Gagang kayu sehingga tidak panas ketika dipegang."),
("Black 20s cotton long sleeve T-shirts, open-end yarn, 100% cotton, tubular fit, seamless double needle 2cm collar, taped neck and shoulders, satin & cotton label,
rib cuffs, double needle bottom hem, olive green discharge ink screen print."),
("Black slide sandals with white ink embossed screen printing

38:24,5cm | 39:25cm | 40:25,5cm | 41:26,5cm | 42:27cm | 43:28cm | 44:28,5cm"),
("Black nylon 1000D mini shoulder / sling bag, with rubber patch at front.

Width:17cm, Height:15cm, Depth:5cm, Volume:1,2lt");

INSERT INTO product(product_type_id, operator_id, product_description_id, name, status, created_at, updated_at) VALUES 
(1,3,1,'USB C to C Ugreen',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(1,3,2,'ADAPTOR C to C Ugreen',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(2,1,3,'Milk pot steamer/ panci susu steamer/panci gagang serbaguna',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(2,1,4,'Yoshikawa Egg Pan Teflon Omelette Pan Wajan Telur Tamagoyaki',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(2,1,5,'panci mie stainless tebal gagang kayu 18 cm / panci susu',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(3,4,6,'Maternal Disaster-Tshirt-Rhizanthes',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(3,4,7,'Maternal Disaster-Footwear-Ss 39',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),
(3,4,8,'Maternal Disaster-Bags-Xars',1,'2023-03-01 11:15:00','2023-03-01 11:15:00');

INSERT INTO payment_method_description(description)
VALUES
("Bagi pengguna BRImo, berikut cara bayar Shopee lewat BRImo dengan praktis. 
Pengguna selesaikan pesanan dan muncul kode BRIVA. 
Buka BRImo. 
Login dengan PIN. 
Buka menu Pembayaran. 
Klik BRIVA. 
Masukkan nomor BRIVA. 
Masukkan PIN saat konfirmasi. 
Klik Send. 
Tunggu notifikasi SMS untuk bukti transaksi."),
("Bagi pengguna Mobile banking, berikut cara bayar Shopee lewat BCA dengan praktis. 
Pengguna menyelesaikan pesanan hingga muncul kode VA. 
Buka BCA Mobile. 
Login dengan PIN. 
Buka menu Transaksi. 
Klik Transfer. 
Klik BCA Virtual Account. 
Masukkan kode 126 + Nomor HP Terdaftar. 
Masukkan PIN saat konfirmasi. 
Klik Kirim. 
Tunggu hingga notifikasi muncul."), 
("Pada bagian Metode Pembayaran saat checkout, silahkan pilih Transfer Bank.
Kemudian pilih salah satu bank, misalnya BNI (Dicek Otomatis).
Tap Konfirmasi.
Tunggu sampai muncul tab baru berisi Nomor Virtual Account yang memiliki batas waktu tertentu.
Silahkan salin atau copy kode VA yang ditampilkan.
Buka aplikasi Gojek yang terinstall di HP kamu lalu login.
Pada halaman utama, pilih opsi Bayar.
Selanjutnya pilih Transfer Instan Gopay ke Rekening Baru.
Tap Bank Transfer lalu BNI (sesuai yang dipilih tadi).
Paste nomor Virtual Account yang sudah disalin melalui langkah di atas.
Berikutnya pilih Konfirmasi.
Akan muncul tab baru berisi detail transfer berupa total pembayaran Shopee beserta nama akun kamu.
Pastikan datanya sudah benar lalu pilih Lanjut.
Ketikkan nominal transfer, harus persis dengan yang ada di tagihan Shopee di langkah awal.
Tap Lanjut.
Saat memasukkan nominal, kamu tidak perlu menambah jumlah nilai transfer dengan biaya admin Gopay. Semua itu sudah terhitung secara otomatis di aplikasi Shopee.
Masukkan PIN Gopay untuk mengkonfirmasi tindakan transfer.
Jika sudah, akan muncul keterangan bahwa transfer rekening BNI sudah berhasil.");

INSERT INTO payment_method(payment_method_description_id, name, payment_number) 
VALUES 
(1,'BRI','00010101822534'),
(2,'BCA','5220304312'),
(3,'GOPAY','08154856255539');

INSERT INTO alamat(alamat)
VALUES 
('Jakarta'),
('Bali'),
('Surakarta');

INSERT INTO user(alamat_id, name, birth_date, gender, status, created_at, updated_at) 
VALUES 
(1,'dyas','2007-06-15','L',1,'2023-03-08 11:15:00','2023-03-10 09:15:00'),
(1,'nindy','2004-07-08','P',1,'2023-03-14 11:15:00','2023-03-15 09:15:00'),
(2,'ardi','2002-05-12','L',1,'2023-03-15 11:15:00','2023-03-16 09:15:00'),
(2,'louis','2003-05-14','L',1,'2023-03-16 11:15:00','2023-03-17 09:15:00'),
(3,'angga','2001-06-22','L',1,'2023-03-17 11:15:00','2023-03-18 09:15:00');

INSERT INTO transaction(user_id, payment_method_id, qty, total_price, status, created_at, updated_at) 
VALUES 
(1,1,3,500000.00,1,'2023-03-08 13:15:00','2023-03-08 13:15:00'),
(1,1,3,200000.00,1,'2023-03-08 14:15:00','2023-03-08 14:15:00'),
(1,1,3,400000.00,1,'2023-03-08 17:15:00','2023-03-08 17:15:00'),
(2,2,3,500000.00,1,'2023-03-14 13:15:00','2023-03-14 13:15:00'),
(2,2,3,400000.00,1,'2023-03-14 14:15:00','2023-03-14 14:15:00'),
(2,2,3,600000.00,1,'2023-03-14 15:15:00','2023-03-14 15:15:00'),
(3,2,3,400000.00,1,'2023-03-15 13:15:00','2023-03-15 13:15:00'),
(3,2,3,500000.00,1,'2023-03-15 14:15:00','2023-03-15 14:15:00'),
(3,2,3,600000.00,1,'2023-03-15 15:15:00','2023-03-15 15:15:00'),
(4,2,3,400000.00,1,'2023-03-16 13:15:00','2023-03-16 13:15:00'),
(4,2,3,500000.00,1,'2023-03-16 14:15:00','2023-03-16 14:15:00'),
(4,2,3,600000.00,1,'2023-03-16 15:15:00','2023-03-16 15:15:00'),
(5,3,3,400000.00,1,'2023-03-17 13:15:00','2023-03-17 13:15:00'),
(5,3,3,500000.00,1,'2023-03-17 14:15:00','2023-03-17 14:15:00'),
(5,3,3,200000.00,1,'2023-03-17 15:15:00','2023-03-17 15:15:00');

INSERT INTO transaction_detail(product_id, transaction_id, qty, price) 
VALUES 
(6,1,1,300000.00),
(2,1,1,150000.00),
(5,1,1,50000.00),
(1,2,1,50000.00),
(5,2,1,50000.00),
(7,2,1,100000.00),
(8,3,1,200000.00),
(2,3,1,150000.00),
(1,3,1,50000.00),
(6,4,1,300000.00),
(2,4,1,150000.00),
(5,4,1,50000.00),
(8,5,1,200000.00),
(2,5,1,150000.00),
(1,5,1,50000.00),
(6,6,1,200000.00),
(8,6,1,150000.00),
(7,6,1,50000.00),
(8,7,1,200000.00),
(2,7,1,150000.00),
(1,7,1,50000.00),
(6,8,1,300000.00),
(2,8,1,150000.00),
(5,8,1,50000.00),
(6,9,1,200000.00),
(8,9,1,150000.00),
(7,9,1,50000.00),
(8,10,1,200000.00),
(2,10,1,150000.00),
(1,10,1,50000.00),
(6,11,1,300000.00),
(2,11,1,150000.00),
(5,11,1,50000.00),
(6,12,1,200000.00),
(8,12,1,150000.00),
(7,12,1,50000.00),
(8,13,1,200000.00),
(2,13,1,150000.00),
(1,13,1,50000.00),
(6,14,1,300000.00),
(2,14,1,150000.00),
(5,14,1,50000.00),
(1,15,1,50000.00),
(5,15,1,50000.00),
(7,15,1,100000.00);