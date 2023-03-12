## Apa Itu Clean Code

Clean Code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah oleh programmer.

---

## Karakteristik Clean Code

1. Penamaan mudah dipahami
1. Mudah dieja dan dipahami
1. singkat namun mendeskripsikan konteks
1. Konsisten
1. Hindari penambahan konteks yang tidak perlu
1. komentar secukupnya saja, terletak di atas function
1. Function
   - jangan terlalu banyak argumen
   - bedakan penamaan variabel global untuk menghindari data terubah / rusak
1. Gunakan konvesi
   menggunakan style guide dari perusahaan besar seperti airbnb, google dan lainya
1. Formatting
   - lebar baris code 80 - 120 karakter
   - satu class 300-500 kata
   - baris code yang berhubungan saling berdekatan
   - dekatkan fungsi dengan pemanggilnya
   - deklarasi variabel berdekatan dengan penggunanya
   - perhatikan indentasi
   - menggunakan prettier atau formatter

---

## Prinsip Dalam Clean Code

**KISS**
Keep It So Simple

Hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dst.

**DRY**
Dont Repeat Yourself

Duplikasi code sering terjadi karena copy paste. untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

**Refactoring**

refactoring adalah proses restrukturusasu kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal Prinsip KISS dan DRY bisa dicapai dengan cara refactor
