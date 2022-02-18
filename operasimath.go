package main

import "fmt"

func main(){
	var a = 100
	var b = 1000
	var c = a*b 

	fmt.Println(c)

	var i = 30
	i*= 10000
	fmt.Println(i)

	i++
	fmt.Println(i)
}