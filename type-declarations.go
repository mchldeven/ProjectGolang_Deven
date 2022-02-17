package main

import "fmt"

func main(){
	type noKTP string
	type Menikah bool

	var noKtpDeven noKTP ="9998344411221"
	fmt.Println(noKtpDeven)
	var statusMenikah Menikah = true
	fmt.Println(statusMenikah)

}