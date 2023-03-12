## Goroutine

Goroutine adalah fungsi atau method yang dijalankan secara independent dengan fungsi atau method yang lainya. goroutine berjalan di atas tread dengan ukuran yang sangat kecil yaitu 2kb. Mekanisme dalam goroutine ketika ada suatu antrian suatu pekerjaan maka dia akan menjalankan yang cepat dan akan mengembalikan pekerjaan yang dinilai lama untuk di proses di akhir. Jika dalam suatu tread antrian pekerjaan sudah tidak ada maka akan mengambil pekerjaan dari tread lainya untuk dikerjakan.

contoh Goroutine :

```Go

printName{
    fmt.Println("Devis Arya")
}
function main() {
    go hello()
    fmt.Println("main function")
    time.Sleep(time.Second)
}
```

---

## Channel

Channel digunakan untuk berkomunikasi antara gorourine satu dengan goroutine yang lainya. ada 2 macam channel yaitu buffered chanel dan unbufered channel

unbufered channel yaitu chanel yang hanya akan menerima satu data. Jika channel sudah terisi data maka tidak bisa di isikan lagi dan jika suatu data sudah dikeluarkan dari channel maka channel siap untuk diisikan data lagi begitu seterusnya.

contoh unbuffered channel :

```Go

func isiData(ch chan string){
    data := <-ch
}
func main() {
    ch := make(chan string)

    go isiData(ch)

    c <- "Devis"
}
```

contoh buffered channel:

```Go

func main() {
    ch := make(chan string)

    ch <- "Devis"
    ch <- "Arya"

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

---

## Select

Select digunakan untuk mempermudah kontrol data dari satu atau banyak channel.

Contoh Select

```Go
func main() {
    number := []int{1,2,3,4,5}

    var ch1 = make(chan float64)
    go getAverage(number, ch1)

    var ch2 = make(chan int)
    go getMax(number, ch2)

   for i := 0 ; i < 2 ; i++{
    select{
        case avg := <-ch1:
            fmt.Println(avg)
        case max := <-ch2:
            fmt.Println(avg)
    }
   }
}
```
