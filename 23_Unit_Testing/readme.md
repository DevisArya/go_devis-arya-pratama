## Middleware

Middleware adalah sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di-proses. Middleware biasanya dibuat per-fungsi-nya, contohnya: middleware autentikasi, middleware untuk logging, middleware untuk gzip compression, dan lainnya.

## Echo Middleware

**Before Router**

Echo#Pre()dapat digunakan untuk mendaftarkan middleware yang dijalankan sebelum router memproses permintaan. Sangat membantu untuk membuat perubahan apa pun pada properti permintaan, misalnya, menambahkan atau menghapus garis miring dari jalur sehingga cocok dengan rute.

Middleware bawaan berikut harus didaftarkan pada level ini:

- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect
- NonWWWRedirect
- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride
- Rewrite

**After Router**

Middleware ini dijalankan setelah router memproses permintaan dan memiliki akses penuh ke echo.ContextAPI.
The following built-in middleware should be registered at this level:

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

## JWT

JWT merupakan salah satu standar JSON (RFC 7519) untuk keperluan akses token. Token dibentuk dari kombinasi beberapa informasi yang di-encode dan di-enkripsi. Informasi yang dimaksud adalah header, payload, dan signature.

**Header**, isinya adalah jenis algoritma yang digunakan untuk generate signature.
**Payload**, isinya adalah data penting untuk keperluan otentikasi, seperti grant, group, kapan login terjadi, dan atau lainnya. Data ini dalam konteks JWT biasa disebut dengan CLAIMS.
**Signature**, isinya adalah hasil dari enkripsi data menggunakan algoritma kriptografi. Data yang dimaksud adalah gabungan dari (encoded) header, (encoded) payload, dan secret key.
