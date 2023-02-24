## (5) Basic Programming

### Introduction Golang

Golang adalah general purpose programming language yang mempermudah dalam pembuatan perangkat lunak, sederhana, handal, dan efisien. Golang adalah bahasa yang bagus untuk menulis program tingkat rendah yang menyediakan layanan ke sistem lain, yang disebut system programming

### Package "Fmt"

Output :

- fmt.Printf()
- fmt.Prinln()
- fmt.Sprintf()
  Scanning
- fmt.Scanln()
  Format Verb
- %T, %v, %s, %q, etc.

### Variables, Types, Const, Zero value

#### Variabel

Variabel digunakan untuk **menyimpan informasi** dalam program komputer, juga untuk memberi label data dengan nama deskriptif yang memiliki tipe data (boolean, numerik, string).

#### Data Type

| **Data Type** |                    **Deskripsi**                     |
| :-----------: | :--------------------------------------------------: |
|    Boolean    |                     True / False                     |
|     unit8     |                       0 to 255                       |
|    uint16     |                      0 to 65535                      |
|    uint32     |                   0 to 4294967295                    |
|    uint64     |              0 to 18446744073709551615               |
|     int8      |                    (-128 to 128)                     |
|     int16     |                  (-32768 to 32767)                   |
|     int32     |             (-2147483648 to 2147483647)              |
|     int64     |    (-9223372036854775808 to 9223372036854775807)     |
|      int      |     sama seperti int32 / int64 berdasarkan value     |
|     rune      |                  sama seperti int32                  |
|     uint      |    sama seperti uint32 / uint64 berdasarkan value    |
|     byte      |                  sama seperti uint8                  |
|    float32    |         IEE-754 32-bit floating-poin-numbers         |
|    float64    |         IEE-754 64-bit floating-poin-numbers         |
|   complex64   | Complex number with float32 real and imaginary parts |
|  complex128   | Complex number with float64 real and imaginary parts |

#### Const & Var

Const dan Var digunakan untuk deklarasi. bedanya Const digunakan untuk nilai yang konstan atau tidak bisa dirubah nilainya sedangkan Var sebaliknya digunakan untuk data yang bisa berubah nilainya.

Contoh penggunaanya:

```go
const phi float64 = 3.14

var isLogin bool = false

var prime int = 7

var decimal float64 = 45.6

var hello string = "Hello World"
```

#### Zero Values

| boleans  | false |
| :------: | :---: |
|  floats  |  0.0  |
| integers |   0   |
| strings  |  ""   |

### Arithmetic Operators

Operator Aritmatika:
+|Penjumlahan
:-----:|:-----:
-|Pengurangan
\*|Perkalian
/|Pembagian
%|Modulo
++|Increment
--|Decrement

Operator Lainya:

- Comparasion ( ==, !=, >, <, >=, <= )
- Logical ( &&, || )
- Bitwise ( &, |, ^, <<, >> )
- Assignment ( =, +=, -=, \*=, /=, %=, <<=, >>=, &=, ^=, |= )
- Miscellaneous ( &, \*(Pointer) )

### Control Structures (Branching & Looping)

Hanya ada beberapa struktur kontrol di Go. untuk bercabang kita gunakan if dan switch, untuk menulis loop kami menggunakan kata kunci for.

Contoh IF, ELSE IF ELSE :

```go
hour := 15
if hour < 12 {
  fmt.Println("Selamat Pagi")
} else if hour < 18 {
  fmt.Println("Selamat Sore")
} else {
  fmt.Println("Selamat Malam")
}

```

Contoh Switch :

```go
var today int = 2
switch today {
    case 1:
        fmt.Printf("Today is Monday")
    case 2:
        fmt.Printf("Today is Tuesday")
    default:
        fmt.Printf("Invalid Date")
}
```

Contoh For :

```go
package main
import "fmt"
func main() {
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
}
```
