package main

import "fmt"

type Student struct {
	name []string
	score []float64
}

func (a Student) GetAverage() float64 {

	var avg float64

	for _, val := range a.score{
		avg += val
	}
	return avg/5
}

func (n Student) GetMin() (float64, string) {
	
	index, Min := 0, n.score[0]

	for i, val := range n.score{
		if val < Min {
			Min = val
			index = i
		}
	}
	return Min, n.name[index]
}

func (x Student) GetMax() (float64, string){

	index, Max := 0, x.score[0]

	for i, val := range x.score{
		if val > Max {
			Max = val
			index = i
		}
	}
	return Max, x.name[index]
}

func main() {
	var nama []string
	var nilai []float64
	var n string
	var s float64


	for i := 0; i < 5; i++ {
		fmt.Printf("Input Students Name: ")
		fmt.Scan(&n)
		fmt.Printf("Input Students Score: ")
		fmt.Scan(&s)
		
		nama = append(nama, n)
		nilai = append(nilai, s)
	}

	student:= Student{
		name: nama,
		score: nilai,
	}

	scoreMin, nameMin := student.GetMin()
	scoreMax, nameMax := student.GetMax()

	fmt.Println("Output :")

	fmt.Println("Average Score: ", student.GetAverage())
	fmt.Println("Min Score of Students: ", nameMin, "(", scoreMin, ")" )
	fmt.Println("Max Score of Students: ", nameMax, "(", scoreMax, ")" )

}

