package Basic

import "fmt"

func Bitwise(){
	a:=10
	b:=3
	fmt.Println(a&b)	//0010
	fmt.Println(a | b)//1011
	fmt.Println(a^b)//1001
	fmt.Println(a&^b)//0100
}

func BitShifting(){
	a:=7		//2^3
	fmt.Println(a<<3)   //shift left 3 places     2^3 *2^3 = 64
	fmt.Println(a>>3)	//shift right 2 places   2^3/2^3  =1
}