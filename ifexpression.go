package main

import "fmt"

func main(){
	
	var nama = "Michael"

	if nama == "eveline" {
		fmt.Println("Hello Eveline")
	} else if nama == "deven" {
		fmt.Println("hello Deven")
	} else if nama == "Dimas" {
		fmt.Println("Hellow Dimass")
	} else { 
		fmt.Println("MENU SELANJUTNYA")
	}

	
	if length :=len (nama) ;length > 5 {
		fmt.Println("Teralalu Panjang") 
	} else {
		fmt.Println("Nama Sudah BENAR")
	}
}