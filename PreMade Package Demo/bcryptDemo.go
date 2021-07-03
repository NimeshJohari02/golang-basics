package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
func main() {
	s:="MySuperComplicatedPassword"
	bs,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err==nil{
		fmt.Println(bs)
	}else {
		fmt.Println(err)
	}
}