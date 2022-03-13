package Basic

import "fmt"
func Pointer(){
	var x int=40
	var y *int =&x
	fmt.Println(x)		//40
	fmt.Println(y)		//0xc000012070      address of x stored
	//if you want value not address you have to derefer
	fmt.Println(*y)			//40

}

type Employees struct{
	EmpName string
}

func  StructPointer()  {
	var emp *Employees

	fmt.Println(emp)		//nil
	emp=new (Employees)
	fmt.Println(*emp)
}



func PointerExp(){
		x:=10  
		changeX(&x) 		//10 
		fmt.Println(x)  	//0
	 }  
func changeX(x *int){  
	*x=0  
 } 


 func PointerExp2() {  
	ptr := new(int)
	fmt.Println("Before change ptr",*ptr)  			//0
	changePtr(ptr)  
	fmt.Println("After change ptr",*ptr) 			//10 
 }  
 func changePtr(ptr *int)  {  
	*ptr = 10  
 }

 func PointertoPointer(){
	 var a int
	 var  ptr *int
	 var ptrr **int
	 a=100
	 ptr=&a
	 ptrr=&ptr
	 fmt.Println(a)
	 fmt.Println(*ptr)
	 fmt.Println(**ptrr)
 }