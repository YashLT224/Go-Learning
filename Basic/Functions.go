package Basic

import "fmt"

func BasicFunc(){
	show()
}

func show(){
	fmt.Println("basic function")
}

//------------------------------/------------------------------/------------------------------/------------------------------------------------------------
func ReturnDataTypeFunc(){
	a:=IntFunc(4)
	fmt.Println(a)
	fmt.Println(StringFunc("Hi"))
}
func IntFunc(n int) int{
	return (n*n)
}
func StringFunc(s string )string{
return "hi"
}

//------------------------------/-------------------------------/-----------------------------/------------------------------/------------------------------

func VariadicFunc(){
	fmt.Println(summation(1,2,3,4,5))
}
func summation(vals...int)int{
	sum:=0
	for _,n:=range vals{
		sum+=n
	}
	return sum
}

//------------------------------/-------------------------------/-----------------------------/------------------------------/------------------------------
//In go you can return mutiple values from function

func MultipleReturnFunc(){
sum,sub:=multiple(2,5)
	fmt.Println(sum)
	fmt.Println(sub)
}
func multiple(x int,y int)(int,int){
 
	return (x+y),(x-y)
}
//------------------------------/-------------------------------/-----------------------------/------------------------------/------------------------------
//koi function struct type ka banana hai toh

type EmployeeData struct {  
	fname string  
	lname string  
 }  
 func StructTypeFunc() {  
	e1 := EmployeeData{ "John","Ponting"}  
	fmt.Println(e1.fullname())
 } 

 func (emp EmployeeData) fullname() string{  
	return emp.fname+" "+emp.lname 
 }  
 
 //------------------------------/-------------------------------/-----------------------------/------------------------------/------------------------------
//anonymous function function within a function

func Anonymounsfunc(){

	f:=func()string{
		return "define here call later"
	}
	fmt.Println(f())   //here we have call that function

	 func() {
		fmt.Println("immediately invoke/call")
	}()
	 



}