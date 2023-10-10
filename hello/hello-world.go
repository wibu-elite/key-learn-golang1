package main

import "fmt"

func main() {
	var anggotaSatu string = "Key"

	// var anggotaDua string
	// anggotaDua = "Ken"

	anggotaTiga := "Joy"

	timSatu, timDua, timTiga := "satu", "dua", "tiga"

	var inputSatu string

	myName := new(string)

	var try bool = true

	fmt.Printf("try? %t", try)

	fmt.Println(*myName)

	fmt.Print("Input Masukan Pertama:")
	fmt.Scan(&inputSatu)
	fmt.Println("Your First Input is : ", inputSatu)

	fmt.Println("Hello World")
	fmt.Println(anggotaSatu, anggotaTiga)
	fmt.Println(timSatu, timDua, timTiga)

}
