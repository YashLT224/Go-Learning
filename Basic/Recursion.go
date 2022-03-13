package Basic

import "fmt"

func Recursion(){
var a= Factorial(5)
fmt.Println(a)
}
func Factorial( i int)int{
if(i<=1){
	return 1
}
return i*Factorial(i-1)
}

 