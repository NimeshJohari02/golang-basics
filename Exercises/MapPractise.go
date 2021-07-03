package main

import "fmt"
func main() {
	myMap:=map[string][]string{
		// The below given also works it's just that it's redundant because you already specified the type in 
		// The map Declaration 
		"a-B":[]string{"this","is","sample"},
		"b-B":{"this","is","sample1"},
		"c-B":{"this","is","sample2"},
		"d-B":{"this","is","sample3"},
		"e-B":{"this","is","sample4"},
		"f-B":{"this","is","sample5"},
	}	
	//! Add 
	myMap["g-B"]=[]string{"This", "is","added ","later"}
	for i,v:= range myMap{
		fmt.Println("Index is ",i);
		for _,v2 := range v{
			fmt.Println(v2)
		}
	}
	//Delete A Record From 
	delete(myMap,"a-B");
	fmt.Println(myMap)
}