// index dimulai dari 0
// index 0 data deven
// 0 = deven dan seterusnya

package main

import "fmt"

func main(){

	var nama [3] string
	nama[0]= "deven"
	nama[1]= "kezia"
	nama[2]= "eveline"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])
	

	var values = [3]int{
		90,
		85,
		80,
	}

		fmt.Println(values)

		fmt.Println(len(nama))
		fmt.Println(len(nama[0]))
		fmt.Println(len(nama[1]))
		fmt.Println(len(nama[2]))
		fmt.Println(len(values))
		fmt.Println(len(values))
		fmt.Println(len(values))
		
	
}

//ngambil datanya [] index brp
