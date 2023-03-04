package main

import (
	"fmt"
	"strings"
)

var ChiperText = map[string]string{
    "A" : "Q", "B" : "H", "C" : "S",
    "D" : "N", "E" : "Z", "F" : "W",
    "G" : "B", "H" : "T", "I" : "G",
    "J" : "L", "K" : "Y", "L" : "C",
    "M" : "E", "N" : "R", "O" : "X",
    "P" : "P", "Q" : "K", "R" : "I",
    "S" : "J", "T" : "D", "U" : "O",
    "V" : "A", "W" : "U", "X" : "F",
    "Y" : "M", "Z" : "V",
}

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	dataName := strings.Split(strings.ToUpper(s.name), "")

	for i, val := range dataName{
		dataName[i] = ChiperText[val]
	}

	nameEncode = strings.Join(dataName,"")
	
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	dataName := strings.Split(strings.ToUpper(s.name), "")
	decode := make([]string, len(dataName))

	for i, element := range dataName{
		for x, val := range ChiperText{
			if element == val {
				decode[i] = x
			}
		}
	}

	nameDecode = strings.Join(decode,"")

  return nameDecode
}

func main() {
	var menu int
   	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	}
}