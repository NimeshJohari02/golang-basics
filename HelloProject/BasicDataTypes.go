package main

import "fmt"

var global float64

func main() {
	fmt.Println("Hello ")
	intest()
}
/*+If you store using int8 then -129 will give you an error
saying that constant -129 overflows int8 */
func intest() {
	x := 12
	y := 12.122212
	fmt.Println(x)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	global=12.122121212121212
	fmt.Println(global)
	fmt.Printf("%T\n",global)
}
func testbool() {
	fmt.Println(42 >= 23)
	fmt.Println(42 <= 23)
	fmt.Println(42 == 42)
}
