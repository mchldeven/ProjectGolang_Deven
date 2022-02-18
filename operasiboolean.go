//opearsi &&  true false = true sisanya false
//operasi || or/atau false/false false sisanya true
// operasi!  kebalikan nilai true hasil false
// nilai false - true
package main

import "fmt"

 func main(){
	 var ujian = 90
	 var absensi = 80

	 var lulusUjian = ujian >= 80
	 var lulusAbsensi = absensi >= 8

	 fmt.Println(ujian)
	 fmt.Println(absensi)

	 var lulus = lulusUjian&&lulusAbsensi
	 fmt.Println(lulus)
	 
 }