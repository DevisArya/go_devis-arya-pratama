## Array

Array merupakan sekumpulan data yang disimpan dalam sebuah variabel, di bahasa pemrograman Go array **tidak diperbolehkan menyimpan berbagai tipe data yang berbeda kedalam satu variabel array** dan juga kapasitas array ditentukan pada saat pembuatan. Data yang disimpan didalam array **tidak boleh melebihi kapasitas** yang sudah ditentukan.

```go
///penerapan Array

//deklarasi array
var name [2]string

//inisialisasi array
names[0] = "devis"
names[1] = "arya"

//shorthand dan inisialisasi
hobi := [2]string{"futsal","renang"}

//menggunakan keyword make
var kendaraan = make([]string, 2)
kendaraan[0] = "sepeda motor"
kendaraan[1] = "pesawat"

```

---

## Slice

Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

```go
///penerapan Slice

//inisialisasi slice
var name = []string{"devis", "arya", "pratama"}

//shorthand dan inisialisasi
hobi := []string{"futsal","renang"}

//fungsi append
hobi = append(hobi, "basketball")
```

---

## Map

Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value. Untuk setiap data (atau value) yang disimpan memiliki suatu unique key. key digunakan untuk pengaksesan value yang di inginkan.

```go
///penerapan Map

point = map[string]int{}

point["barcelona"] = 59
point["real madrid"] = 52

//menghapus item map
delete(point, "real madrid")

//cek item map
var value, isExist = point["sevila"]

if isExist {
    fmt.Println(value)
} else {
    fmt.Println("item tidak ada")
}
```
