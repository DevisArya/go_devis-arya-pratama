## DATABASE RELATIONSHIP

1. **ONE TO ONE** Adalah relasi tabel dalam database dimana data dari tabel a mempunyai hubungan dengan tabel b sebanyak 1
   Contohnya : Tabel product dan product description
2. **ONE TO MANY** Adalah relasi tabel dalam database dimana data dari tabel a mempunyai hubungan dengan tabel b sebanyak x data
   Contohnya : Tabel transaction dan transaction_detail
3. **MANY TO MANY** Adalah banyak data pada sebuah tabel memiliki relasi ke banyak data juga pada tabel yang lainnya.
   Contohnya : Tabel product dan transaction.

---

## DDL

DDL (Data Definition Languange) merupakan sekumpulan set perintah yang bertujuan
untuk mendefinisikan atribut – atribut database, tabel, atribut kolom (field), maupun
batasan – batasan terhadap suatu atribut dan relasi/hubungan antar tabel. Yang termasuk
dalam kelompok perintah DDL adalah : CREATE, ALTER, dan DROP.
CREATE merupakan perintah DDL yang digunakan untuk membuat database maupun
tabel. Nama database maupun tabel tidak boleh mengandung spasi (space). Nama
database tidak boleh sama antar database.
ALTER merupakan perintah DDL yang digunakan untuk mengubah nama/struktur tabel.
DROP merupakan perintah DDL yang digunakan untuk menghapus database ataupun
tabel

---

## DML

Merupakan bentuk bahasa basis data yang berguna untuk melakukan manipulasi data
dan pengambilan data pada suatu basis data. Manipulasi data dapat berupa :

- penyisipan/penambahan data baru ke suatu basis data
- penghapusan data dari suatu basis data
- pengubahan data di suatu basis data

Contoh :

```Sql
--Perintah untuk menambahkan data kedalam table
INSERT INTO namatabel(field1,field2, field_n) VALUES
(nilai1, nilai2, nilai_n);

--perintah untuk menghapus baris nilai dalam table
DELETE FROM namatabel WHERE [kondisi];

--perintah untuk mengubah nilai dalam suatu table
UPDATE namatable SET namafield = nilai1
WHERE [kondisi];

--perintah untuk menampilkan data dari table

SELECT * FROM namatable;
SELECT * FROM namatable WHERE[kondisi];
SELECT namafield1,namafiled2 FROM namatable WHERE[kondisi];
```
