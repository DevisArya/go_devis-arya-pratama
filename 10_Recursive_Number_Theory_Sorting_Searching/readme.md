## Recursive

recursive adalah suatu cara untuk menyelesaikan suatu masalah dengen fungsi yang memanggil ulang fungsi itu sendiri

Contohnya

```Go

func factorial(n int)int{
    if (n == 1) {
        return 1
    }
    return n * factorial(n-1)
}
```

fungsi factorial diatas akan memanggil fungsinya lagi sampai n kali jadi misalnya ingin mencari nilai 3! yang terjadi adalah 1 x 2 x 3 = 6

---

## Searching

Searching adalah proses mencari suatu data dari tumpukan data.

Contohnya
**Linear Search**

```Go
func search(list []int, nilai int)bool{
    for _, item := range list {
		if item == nilai {
			return true
		}
	}
	return false
}
```

---

## Sorting

Sorting adalah proses untuk mengurutkan suatu data, bisa dari yang terkecil ke yang terbesar atau sebaliknya. Sorting juga bisa digunakan untuk mengurutkan data numeric, kata, pair dan lainya

contohnya
**Counting Sort**

```Go
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	max := 0
	min := 0
	for _, num := range nums {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	l := max - min + 1
	countArr := make([]int, l)
	resultArr := make([]int, len(nums))
	for _, num := range nums {
		countArr[num-min] ++
	}
	for i := 1; i < l; i++ {
		countArr[i] += countArr[i-1]
	}
	for i := len(resultArr) - 1; i > -1; i-- {
		resultArr[countArr[nums[i]-min]-1] = nums[i]
		countArr[nums[i]-min] --
	}
	return resultArr
}
```
