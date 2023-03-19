## JOIN

Join adalah sebuah perintah untuk mengkombinasikan record dari dua atau lebih tabel. ada beberapa perintah join seperti standar SQL ANSI yaitu :

- **INNER JOIN** yaitu menampilkan data record yang saling berhubungan antara tabel a dan tabel b atau lebih banyak.
- **LEFT JOIN** yaitu menampilkan data record dari tabel a dan dari tabel b yang saling berhubungan dengan tabel a.
- **RIGHT JOIN** yaitu menampilkan data record dari tabel b dan dari tabel a yang saling berhubungan dengan tabel b.

---

## UNION

UNION adalah perintah untuk menggabungkan hasil dari perintah SELECT tabel 1 dan perintah SELECT tabel 2, juga bisa untuk lebih dari 2 perintah SELECT.

**namun** ada poin yang perlu diperhatikan sebelum menggabungkan hasil perintah select yaitu filed pada perintah SELECT 1 dan lainya harus sama.

Contohnya:

```Sql
SELECT * FROM table1
UNION / UNION ALL
SELECT * FROM table2
```

UNION digunakan untuk menampilkan data dari 2 atau lebih perintah select namum akan membuat **DISTINC**/menampilkan hanya sekali data yang memiliki kesamaan,
sedangkan untuk UNION ALL menampilkan semuanya termasuk data yang sama akan ditampilkan.

---

### FUNGSI AGREGASI SQL

Fungsi agregasi digunakan untuk melakukan perhitungan terhadap nilai-nilai hasil suatu query menggunakan SQL.

- **MIN** digunakan untuk mencari data terendah dari sekumpulan record data dalam database
- **MAX** digunakan untuk mencari data tertinggi dari sekumpulan record data dalam database
- **SUM** digunakan untuk total penjumlahan nilai dari data record tabel.
- **AVG** digunakan untuk mencari nilai rata-rata dari data record tabel.
- **COUNT** digunakan untuk mencari berapa banyak data yang ada didalam record tabel.
- **HAVING** digunakan sebagai logic seperti **WHERE** namun hanya digunakan jika mempunyai FUNGSI AGREGASI dalam suatu perintah.
