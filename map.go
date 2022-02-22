package main

import "fmt"
func main(){

	person := map[string]string{

		"nama": "Michael",
		"alamat": "Viladago",
	}
	person["status"] = "single"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	

	var book map[string]string = make(map[string]string)
	book["title"] = "Universal BPR"
	book["Author"] = "Deven"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))
	
	delete(book,"ups")

	fmt.Println(book)
	
}