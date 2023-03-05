## String

dalam sebuah string ada beberapa build in function yang bisa digunakan yaitu :

1. **Len** untuk mencari panjang suatu string
1. **Compare** untuk membandingkan string
1. **Contains** untuk mengecek apakah string merupakan bagian dari string lainya
1. **Substring** untuk mengambil bagian dari string berdasarkan index
1. **Replace** untuk untuk merubah bagian string yang di inginkan
1. **Insert** untuk menyisipkan / memasukan string ke string lainya

---

## Advance Function

#### Variadic Function

Variadic Function digunakan ketika suatu function akan digunan untuk menerima beberapa parameter yang tidak tentu.

Contohnya

```Go
func sum(numbers ...int)int{
    var total int = 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main(){
    result := sum(2, 3, 4, 5, 6)
    result1 := sum(1, 30)
}
```

#### Anonymous Function

Anonymous Function atau Literal Function adalah fungsi yang tidak memiliki nama. ini sangat berguna ketika ingin memebuat inline function

Contohnya

```Go
func main() {
  // Anonymous function
  func() {
    fmt.Println("Hello World")
  }()

  // memasukan kedalam suatu variabel
  value := func() {
    fmt.Println("Hello Worlds")
  }
  value()

  // Passing arguments in anonymous function
  func(sentence string) {
    fmt.Println(sentence)
  }("Alterra")
}

```

#### Defer Function

fungsi ini akan dijalankan ketika fungsi parent dari fungsi ini sudah melakukan return.

Contohnya

```Go
func main() {
  defer func() {
    fmt.Println("Later")
  }()

  fmt.Println("First")

}
```

---

## Pointer

Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.

Contohnya

```Go

func main() {
  var name string = "John"
  var nameAddress *string = &name
  fmt.Println("name (value)   :", name)                // John
  fmt.Println("name (address) :", &name)               // 0xc000010050
  fmt.Println("nameAddress (value)   :", *nameAddress) // John
  fmt.Println("nameAddress (address) :", nameAddress)  // 0xc000010050
}
```
