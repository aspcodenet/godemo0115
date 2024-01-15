package main

import "fmt"

func main() {
	//var l = new List<Player>();
	var name string
	var age int

	fmt.Print("Vad heter du?")
	fmt.Scanln(&name)

	fmt.Print("Hur ganmmal är du")
	_, err := fmt.Scanln(&age)
	if err == nil {
		fmt.Printf("hej %s du är %d år\n", name, age)
	} else {
		fmt.Printf("Ngt gick fel")
	}

	// var name string = "Stefan Holmberg"

	// var i uint8 = 254
	// //var i int = 12
	// // var i int
	// // i  = 12
	// //i := 12
	// var country, code = "Sweden", 46
	// fmt.Printf("Hejsan %d s\n", code, country)
}
