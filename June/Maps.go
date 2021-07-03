//Maps are an unordered Collection
package main

//map[key_data_type]value_data_type
import "fmt"
func main() {
	x :=make(map[int]string)
	x[0]="Nimesh";
	x[1]="Kumar";
	x[2]="Johari";
	// fmt.Print(x);

	// fmt.Print(x[0]);
	myMap:=map[string]int{
		"Nimesh":3232,
		"Shruti":12121,
		"Johari":908990,
	}
	if _,ok1:=myMap["Shruti"];ok1{
		fmt.Println("Shruti found ");
	}
	for k,v:=range myMap{
		fmt.Println(k,v);

	}
}