package main
import "fmt"

func main() {
	fmt.Print("Hello World ")
	rand()
	fmt.Println("Out of the function ")
	/*The println function returns an error and the number of bytes written here n is the number of bytes written and _ denotes the error but writing _ means that we wont access it Because if You are declaring a variable it is imperative that you use it other wise you get an error saying declared but not used  */
	n, _ := fmt.Println("Hello", 12+12, true)
	fmt.Println(n)
	printEven(200)
}

func rand() {
	fmt.Println("This is inside the rand function ")
}
func printEven(a int) {

	for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
