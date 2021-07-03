package main

//! Values of Same time in slice
import (
	"fmt"
	"os"
)
func main(){
	//if you put size more than needed it will be filled by usign spaces which is not acceptable 
	//x:= type{values} // composite literal
	x:=[]int{1,2,334,32,1}
	fmt.Println(x)
 s:=make([]string,2);
 fmt.Println("Empty String " ,s);
 s[0]="N"
 s[1]="i"
s = append(s, "m")
s = append(s, "e")
s = append(s, "s")
s = append(s, "h")
c:=make([]string, len(s));
copy(c,s);
fmt.Println(s);
l:=s[1:5];
fmt.Println("Value of [1:5] is",l);

l=s[2:5];
fmt.Println("Value of [2:5] is",l);
//!! IN STRING SLICE METHODS THE LAST VALUE IS NOT TAKEN 
//! matlab agar 4 tak limit di hai toh 3 tak hi karna chaiye 
l=s[:5];
fmt.Println("Value of [:5] is",l);
// Making a dynamic array such that ith row has i+1 elements
myArr:= make([][]int,3)
for i:=0;i<3;i++{
	l2:=i+1;
	myArr[i] =make([]int, l2)
	for j:=0;j<l2;j++{
		myArr[i][j]=i+j;
	}
}
fmt.Println("Sum is ", getSum(1,2,3,4,5,6,8,96,34))
os.Exit(0);//This Exits the program before the next line is executed 
fmt.Println("Two Dimensional Array " , myArr);
}
func getSum(x... int) int {
	sum:=0;
	for _,v:=range x{
		sum+=v	
}
return sum
}