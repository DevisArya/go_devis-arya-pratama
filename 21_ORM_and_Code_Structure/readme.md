## ORM

ORM (Object Relation Mapping) merupakan teknik yang merubah suatu table menjadi sebuah object yang nantinya mudah untuk digunakan. Object yang dibuat memiliki property yang sama dengan field — field yang ada pada table tersebut.
Kelebihan:

1. Terdapat banyak fitur seperti transactions, connection pooling migrations, seeds, streams, dan lain sebagainya.
2. perintah query memiliki kinerja yang lebih baik, daripada kita menulisnya secara manual.
3. Kita menulis model data hanya di satu tempat, sehingga lebih mudah untuk update, maintain, dan reuse the code.
4. Membuat akses data menjadi lebih abstrak(kita dapat mengubahnya kapanpun) dan portable.
5. Memungkinkan kita memanfaatkan OOP (object oriented programming) dengan baik

## Gorm

Gorm merupakan ORM yang dikembangkan untuk bahasa GO.
Berikut ini beberapa kelebihan GORM

- Full-Featured ORM
- Associations (has one, has many, belongs to, many to many, polymorphism, single-table inheritance)
- Hooks (before/after create/save/update/delete/find)
- Eager loading with Preload, Joins
- Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- Context, Prepared Statment Mode, DryRun Mode
- Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
- Composite Primary Key, Indexes, Constraints
- Auto Migrations
- Logger
- Extendable, flexible plugin API: Database Resolver (multiple databases, read/write splitting) / Prometheus…
- Every feature comes with tests
- Developer Friendly

Cara install:

> go get -u gorm.io/gorm
> go get -u gorm.io/driver/sqlite

## MVC

Model View Controller atau yang dapat disingkat MVC adalah sebuah pola arsitektur dalam membuat sebuah aplikasi dengan cara memisahkan kode menjadi tiga bagian yang terdiri dari:

**Model**
Bagian yang bertugas untuk menyiapkan, mengatur, memanipulasi, dan mengorganisasikan data yang ada di database.

**View**
Bagian yang bertugas untuk menampilkan informasi dalam bentuk Graphical User Interface (GUI).

**Controller**
Bagian yang bertugas untuk menghubungkan serta mengatur model dan view agar dapat saling terhubung.
