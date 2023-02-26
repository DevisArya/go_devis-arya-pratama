## Apa itu Versioning

Versioning merupakan cara untuk **mengatur versi** dari source code program.
Ketika membangun sebuah program pasti akan melakukan revisi karena tidak semua code yang kita buat sempurna.

---

## GIT

Git merupakan Salah satu **version control system** populer yang digunakan para developer untuk mengembangkan software secara bersama-bersama.

setiap local computer akan tersingkronisasi dengan remote server. Setiap perubahan akan terekam di remote server / repository. untuk membuat repository bisa menggunakan github. github merupakan salah satu git hosting service.

---

## Saving Changes

dibagi menjadi 3 fase yaitu **working directory, staging area, dan repository**

|        **Code**        |                       **Deskripsi**                       |
| :--------------------: | :-------------------------------------------------------: |
|  git add <directory>   |      Untuk memindahkan ke staging area file spesifik      |
|       git add .        | Untuk memindahkan ke staging area semua file yang berubah |
| git commit -m"message" |      Untuk melakukan commit disertai dengan message       |
| git push origin master |    Untuk melakukan push ke reposity dari branch master    |

---

## Branch

Bisa digunakan jika ragu dengan perubahan yang dilakukan apakah berhasil atau gagal. jika berhasil maka akan dilakukan merge untuk menyimpan perubahan di branch utama / master.

### branching

|         **Code**          |                        **Deskripsi**                        |
| :-----------------------: | :---------------------------------------------------------: |
|  git branch <namaBranch>  |             digunakan untuk membuat branch baru             |
| git checkout <namaBranch> |              digunakan untuk berpindah branch               |
|  git branch -D <branch>   |              digunakan untuk menghapus branch               |
|       git branch -a       | digunakan untuk melihat semua branch termasuk remote branch |

### merge

|         **Code**          |           **Deskripsi**            |
| :-----------------------: | :--------------------------------: |
|    git checkout master    |  untuk berpindah ke branch master  |
|   git merge new-feature   | digunakan untuk merge suatu branch |
| git branch -d new-feature |  digunakan untuk menghapus branch  |
