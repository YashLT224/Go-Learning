package Basic

import "fmt"

func CallByReference(){
	var a int=100
	var b int=200
	fmt.Println("Value before swap")
	fmt.Println(a)
	fmt.Println(b)
	swap(&a,&b)
	fmt.Println("Value after swap")
	fmt.Println(a)
	fmt.Println(b)

}

func swap(x *int,y *int){
	var temp int=*x
	*x=*y
	*y=temp
}


func CallByValue(){
	var a int=100
	var b int=200
	fmt.Println("Value before swap")
	fmt.Println(a)
	fmt.Println(b)
	a,b=swapbyval(a,b)
	fmt.Println("Value after swap")
	fmt.Println(a)
	fmt.Println(b)

}

func swapbyval(x int,y int) (int ,int){
	var temp int=x
	x=y
	y=temp

	return x,y
}