package main

//! If you get 123 125 as output then capitalize first letter of class so as to make in Accessible to other packages
//! Also capitalize varibale names insside struct
//! Return of unmarshall ie 2nd param should be a pointer
import (
	"encoding/json"
	"fmt"
	"os"
)
type Student struct{
	Name string
	Regno string
}
func main() {
	st1:=Student{
		Name: "Nimesh",
		Regno: "19BIT0080",
	}
	fmt.Println(st1)
	data, err := json.Marshal(st1)
	if err!=nil{
	fmt.Println(err)
	}else{
	os.Stdout.Write(data)
	}
	var Student1 Student
	err2:=json.Unmarshal([]byte(data) , &Student1)
	if err2==nil{
	fmt.Println(Student1.Name)
	}else{
		fmt.Println(err2)
	}
}
