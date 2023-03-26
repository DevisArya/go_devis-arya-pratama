## Jika mengambil semua data dari tabel user pada database SQL adalah `SELECT * FROM user`

---

Pada database **redis** bisa menggunakan perintah `GET.JSON user` alasanya karena menurut saya pada database redis ini berbasis key dan value. Jadi akan sangat menyusahkan jika tidak dalam bentuk json karena akan memanggil berulang dari key atau regex untuk mendapatkan valuenya.

Pada database **neo4j** bisa menggunakan perintah berikut :
`MATCH (u:user) RETURN u;`

sedangkan pada database **cassandra** perintah untuk menampilkan semua data pada tabel user sama dengan perintah pada database MySql yaitu menggunakan perintah berikut : `SELECT * FROM user`
