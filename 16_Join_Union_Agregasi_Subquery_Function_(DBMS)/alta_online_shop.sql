CREATE DATABASE  IF NOT EXISTS `alta_online_shop` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `alta_online_shop`;
-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: alta_online_shop
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `alamat`
--

DROP TABLE IF EXISTS `alamat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `alamat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alamat` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `alamat`
--

LOCK TABLES `alamat` WRITE;
/*!40000 ALTER TABLE `alamat` DISABLE KEYS */;
INSERT INTO `alamat` VALUES (1,'Jakarta'),(2,'Bali'),(3,'Surakarta');
/*!40000 ALTER TABLE `alamat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operator`
--

DROP TABLE IF EXISTS `operator`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `operator` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone_number` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operator`
--

LOCK TABLES `operator` WRITE;
/*!40000 ALTER TABLE `operator` DISABLE KEYS */;
INSERT INTO `operator` VALUES (1,'ela','Bandung','081234567892'),(2,'robi','Bandung','085484529594'),(3,'bryan','jakarta','0854844052565'),(4,'adam','malang','0854845258152'),(5,'elis','jombang','087548522595');
/*!40000 ALTER TABLE `operator` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_method`
--

DROP TABLE IF EXISTS `payment_method`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment_method` (
  `id` int NOT NULL AUTO_INCREMENT,
  `payment_method_description_id` int NOT NULL,
  `name` varchar(100) NOT NULL,
  `payment_number` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `payment_method_description_id` (`payment_method_description_id`),
  CONSTRAINT `payment_method_ibfk_1` FOREIGN KEY (`payment_method_description_id`) REFERENCES `payment_method_description` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_method`
--

LOCK TABLES `payment_method` WRITE;
/*!40000 ALTER TABLE `payment_method` DISABLE KEYS */;
INSERT INTO `payment_method` VALUES (1,1,'BRI','00010101822534'),(2,2,'BCA','5220304312'),(3,3,'GOPAY','08154856255539');
/*!40000 ALTER TABLE `payment_method` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_method_description`
--

DROP TABLE IF EXISTS `payment_method_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payment_method_description` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_method_description`
--

LOCK TABLES `payment_method_description` WRITE;
/*!40000 ALTER TABLE `payment_method_description` DISABLE KEYS */;
INSERT INTO `payment_method_description` VALUES (1,'Bagi pengguna BRImo, berikut cara bayar Shopee lewat BRImo dengan praktis. \nPengguna selesaikan pesanan dan muncul kode BRIVA. \nBuka BRImo. \nLogin dengan PIN. \nBuka menu Pembayaran. \nKlik BRIVA. \nMasukkan nomor BRIVA. \nMasukkan PIN saat konfirmasi. \nKlik Send. \nTunggu notifikasi SMS untuk bukti transaksi.'),(2,'Bagi pengguna Mobile banking, berikut cara bayar Shopee lewat BCA dengan praktis. \nPengguna menyelesaikan pesanan hingga muncul kode VA. \nBuka BCA Mobile. \nLogin dengan PIN. \nBuka menu Transaksi. \nKlik Transfer. \nKlik BCA Virtual Account. \nMasukkan kode 126 + Nomor HP Terdaftar. \nMasukkan PIN saat konfirmasi. \nKlik Kirim. \nTunggu hingga notifikasi muncul.'),(3,'Pada bagian Metode Pembayaran saat checkout, silahkan pilih Transfer Bank.\nKemudian pilih salah satu bank, misalnya BNI (Dicek Otomatis).\nTap Konfirmasi.\nTunggu sampai muncul tab baru berisi Nomor Virtual Account yang memiliki batas waktu tertentu.\nSilahkan salin atau copy kode VA yang ditampilkan.\nBuka aplikasi Gojek yang terinstall di HP kamu lalu login.\nPada halaman utama, pilih opsi Bayar.\nSelanjutnya pilih Transfer Instan Gopay ke Rekening Baru.\nTap Bank Transfer lalu BNI (sesuai yang dipilih tadi).\nPaste nomor Virtual Account yang sudah disalin melalui langkah di atas.\nBerikutnya pilih Konfirmasi.\nAkan muncul tab baru berisi detail transfer berupa total pembayaran Shopee beserta nama akun kamu.\nPastikan datanya sudah benar lalu pilih Lanjut.\nKetikkan nominal transfer, harus persis dengan yang ada di tagihan Shopee di langkah awal.\nTap Lanjut.\nSaat memasukkan nominal, kamu tidak perlu menambah jumlah nilai transfer dengan biaya admin Gopay. Semua itu sudah terhitung secara otomatis di aplikasi Shopee.\nMasukkan PIN Gopay untuk mengkonfirmasi tindakan transfer.\nJika sudah, akan muncul keterangan bahwa transfer rekening BNI sudah berhasil.');
/*!40000 ALTER TABLE `payment_method_description` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_type_id` int NOT NULL,
  `operator_id` int NOT NULL,
  `product_description_id` int NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_type_id` (`product_type_id`),
  KEY `operator_id` (`operator_id`),
  KEY `product_description_id` (`product_description_id`),
  CONSTRAINT `product_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_type` (`id`),
  CONSTRAINT `product_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operator` (`id`),
  CONSTRAINT `product_ibfk_3` FOREIGN KEY (`product_description_id`) REFERENCES `product_description` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1,1,3,1,'product dummy',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(2,1,3,2,'ADAPTOR C to C Ugreen',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(3,2,1,3,'Milk pot steamer/ panci susu steamer/panci gagang serbaguna',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(4,2,1,4,'Yoshikawa Egg Pan Teflon Omelette Pan Wajan Telur Tamagoyaki',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(5,2,1,5,'panci mie stainless tebal gagang kayu 18 cm / panci susu',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(6,3,4,6,'Maternal Disaster-Tshirt-Rhizanthes',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(7,3,4,7,'Maternal Disaster-Footwear-Ss 39',1,'2023-03-01 11:15:00','2023-03-01 11:15:00'),(8,3,4,8,'Maternal Disaster-Bags-Xars',1,'2023-03-01 11:15:00','2023-03-01 11:15:00');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_description`
--

DROP TABLE IF EXISTS `product_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_description` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_description`
--

LOCK TABLES `product_description` WRITE;
/*!40000 ALTER TABLE `product_description` DISABLE KEYS */;
INSERT INTO `product_description` VALUES (1,'Cocok digunakan untuk semua iPhone yg sudah support penggunaan charger fast charging 20W 8+ / X / XS / XS MAX / 11 / 11 PRO / 11 PRO MAX / 12 / 12 MINI / 12 PRO / 12 PRO MAX / 13 / 13 MINI / 13 PRO / 13 PRO MAX / 14 / 14 PLUS / 14 PRO / 14 PRO MAX\n\nSpesifikasi Adaptor :\n- Model : CD137\n- Input : 100-240V~50/60Hz 500mA Max\n- USB C Output : 5V-3A / 9V-2.22A / 12V-1.67A / 3.3-5.9V-3A / 3.3-11V-1.8A\n- Total Output : 20W Max\n- Warna: PUTIH / HITAM'),(2,'Cocok digunakan untuk semua iPhone yg sudah support penggunaan charger fast charging 20W 8+ / X / XS / XS MAX / 11 / 11 PRO / 11 PRO MAX / 12 / 12 MINI / 12 PRO / 12 PRO MAX / 13 / 13 MINI / 13 PRO / 13 PRO MAX / 14 / 14 PLUS / 14 PRO / 14 PRO MAX\n\nSpesifikasi Kabel :\n- Length : 1m\n- Current : 3A\n- MFi Certificate\n- Warna : PUTIH / HITAM'),(3,'Milk Pot with Steamer \nDiameter 20 cm\nPanci susu serba guna yang bisa anda gunakan untuk memasak makanan seperti mie , soup , dan sekaligus bisa dipakai untuk mengukus makanan.\nTerbuat dari bahan stainless steel berkualitas\ntutup kaca'),(4,'Penggorengan bahan teflon\nUkuran pan 10X16cm\nAnti lengket\nAnti gores\nBisa untuk kompor konvensional\nHandle tahan panas\nNyaman digenggam dan mudah dibersihkan\nType fry pan ini cocok untuk menggoreng telur (Tamagoyaki)'),(5,'Panci Mie / Panci masak 18 cm pegangan kayu\nBahan Almunium Tebal 2mm\nGagang kayu sehingga tidak panas ketika dipegang.'),(6,'Black 20s cotton long sleeve T-shirts, open-end yarn, 100% cotton, tubular fit, seamless double needle 2cm collar, taped neck and shoulders, satin & cotton label,\nrib cuffs, double needle bottom hem, olive green discharge ink screen print.'),(7,'Black slide sandals with white ink embossed screen printing\n\n38:24,5cm | 39:25cm | 40:25,5cm | 41:26,5cm | 42:27cm | 43:28cm | 44:28,5cm'),(8,'Black nylon 1000D mini shoulder / sling bag, with rubber patch at front.\n\nWidth:17cm, Height:15cm, Depth:5cm, Volume:1,2lt');
/*!40000 ALTER TABLE `product_description` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_type`
--

DROP TABLE IF EXISTS `product_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_type`
--

LOCK TABLES `product_type` WRITE;
/*!40000 ALTER TABLE `product_type` DISABLE KEYS */;
INSERT INTO `product_type` VALUES (1,'elektronik'),(2,'perabotan rumah tangga'),(3,'fashion');
/*!40000 ALTER TABLE `product_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `payment_method_id` int NOT NULL,
  `qty` int NOT NULL,
  `total_price` decimal(21,2) DEFAULT NULL,
  `status` smallint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `payment_method_id` (`payment_method_id`),
  CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `transaction_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` VALUES (2,1,1,3,200000.00,1,'2023-03-08 14:15:00','2023-03-08 14:15:00'),(3,1,1,3,400000.00,1,'2023-03-08 17:15:00','2023-03-08 17:15:00'),(4,2,2,3,500000.00,1,'2023-03-14 13:15:00','2023-03-14 13:15:00'),(5,2,2,3,400000.00,1,'2023-03-14 14:15:00','2023-03-14 14:15:00'),(6,2,2,3,600000.00,1,'2023-03-14 15:15:00','2023-03-14 15:15:00'),(7,3,2,3,400000.00,1,'2023-03-15 13:15:00','2023-03-15 13:15:00'),(8,3,2,3,500000.00,1,'2023-03-15 14:15:00','2023-03-15 14:15:00'),(9,3,2,3,600000.00,1,'2023-03-15 15:15:00','2023-03-15 15:15:00'),(10,4,2,3,400000.00,1,'2023-03-16 13:15:00','2023-03-16 13:15:00'),(11,4,2,3,500000.00,1,'2023-03-16 14:15:00','2023-03-16 14:15:00'),(12,4,2,3,600000.00,1,'2023-03-16 15:15:00','2023-03-16 15:15:00'),(13,5,3,3,400000.00,1,'2023-03-17 13:15:00','2023-03-17 13:15:00'),(14,5,3,3,500000.00,1,'2023-03-17 14:15:00','2023-03-17 14:15:00');
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_detail`
--

DROP TABLE IF EXISTS `transaction_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `transaction_id` int NOT NULL,
  `qty` int NOT NULL,
  `price` decimal(21,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `transaction_id` (`transaction_id`),
  CONSTRAINT `transaction_detail_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `transaction_detail_ibfk_2` FOREIGN KEY (`transaction_id`) REFERENCES `transaction` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_detail`
--

LOCK TABLES `transaction_detail` WRITE;
/*!40000 ALTER TABLE `transaction_detail` DISABLE KEYS */;
INSERT INTO `transaction_detail` VALUES (4,1,2,3,50000.00),(5,5,2,1,50000.00),(6,7,2,1,100000.00),(7,8,3,1,200000.00),(8,2,3,1,150000.00),(9,1,3,3,50000.00),(10,6,4,1,300000.00),(11,2,4,1,150000.00),(12,5,4,1,50000.00),(13,8,5,1,200000.00),(14,2,5,1,150000.00),(15,1,5,3,50000.00),(16,6,6,1,200000.00),(17,8,6,1,150000.00),(18,7,6,1,50000.00),(19,8,7,1,200000.00),(20,2,7,1,150000.00),(21,1,7,3,50000.00),(22,6,8,1,300000.00),(23,2,8,1,150000.00),(24,5,8,1,50000.00),(25,6,9,1,200000.00),(26,8,9,1,150000.00),(27,7,9,1,50000.00),(28,8,10,1,200000.00),(29,2,10,1,150000.00),(30,1,10,3,50000.00),(31,6,11,1,300000.00),(32,2,11,1,150000.00),(33,5,11,1,50000.00),(34,6,12,1,200000.00),(35,8,12,1,150000.00),(36,7,12,1,50000.00),(37,8,13,1,200000.00),(38,2,13,1,150000.00),(39,1,13,3,50000.00),(40,6,14,1,300000.00),(41,2,14,1,150000.00),(42,5,14,1,50000.00);
/*!40000 ALTER TABLE `transaction_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alamat_id` int DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `birth_date` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `status` smallint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `alamat_id` (`alamat_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`alamat_id`) REFERENCES `alamat` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,1,'dyas','2007-06-15','L',1,'2023-03-08 11:15:00','2023-03-10 09:15:00'),(2,1,'nindy','2004-07-08','P',1,'2023-03-14 11:15:00','2023-03-15 09:15:00'),(3,2,'ardi','2002-05-12','L',1,'2023-03-15 11:15:00','2023-03-16 09:15:00'),(4,2,'louis','2003-05-14','L',1,'2023-03-16 11:15:00','2023-03-17 09:15:00'),(5,3,'angga','2001-06-22','L',1,'2023-03-17 11:15:00','2023-03-18 09:15:00');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_payment_method_detail`
--

DROP TABLE IF EXISTS `user_payment_method_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_payment_method_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `payment_method_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `payment_method_id` (`payment_method_id`),
  CONSTRAINT `user_payment_method_detail_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_payment_method_detail_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_payment_method_detail`
--

LOCK TABLES `user_payment_method_detail` WRITE;
/*!40000 ALTER TABLE `user_payment_method_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_payment_method_detail` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-20 20:30:01
