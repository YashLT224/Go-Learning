package Basic

import "fmt"
import 	"reflect"

type Employee struct{
	EmpId int 
	EmpName string
	DeptName string
	Phone []string
}

func Struct(){
	emp:=Employee{1,"Yash","CSE",[]string{"12345678,1234564"},}
	fmt.Println(emp)						//{1 Yash CSE [12345678,1234564]}
	fmt.Println(emp.EmpName,emp.EmpId,emp.DeptName,emp.Phone)

	emp1:=Employee{
		EmpId:1,
		EmpName:"Yash",
		DeptName:"CSE",
		Phone:[]string{"12345678,1234564"},}		// ypu can change order in that as well
		fmt.Println(emp1)  						//{1 Yash CSE [12345678,1234564]}

		emp1.EmpId=2
		fmt.Println(emp1)  



		emp2:=emp1
		 

		//only cop data
		emp1.EmpId=3
		fmt.Println(emp1)		//{3 Yash CSE [12345678,1234564]}
		fmt.Println(emp2) 		//{2 Yash CSE [12345678,1234564]}


		//if want to put same address

		emp3:=&emp1
		fmt.Println(emp2)

		//points to same reference change  in one reflects in second as well
		emp1.EmpId=4
		fmt.Println(emp1)		//{4 Yash CSE [12345678,1234564]}
		fmt.Println(emp3) 		//&{4 Yash CSE [12345678,1234564]}
}


type Books struct{
 	name string
	author string
}

func Structargpass(){
	var book1 Books
	var book2 Books

	book1.name="csgeeks"
	book1.author="yash"

	book2.name="javatpoint"
	book2.author="verma"

	Printdetails(book1)
	Printdetails(book2)
}

func Printdetails(book Books){
	   
	fmt.Printf("Book name is  %s  and Book author is %s\n",book.name,book.author)
}



// ----------------nested struct

type Animal struct{
	name string
	origin string
}

type Bird struct{
	Animal
	speedKPH float32
	fly bool
}

func StructInheritance(){
	b:=Bird{}
	b.name="Emu"
	b.origin="australia"
	b.speedKPH=45.4
	b.fly=false
	fmt.Println(b)

}

// -------------struct tags
type Animals struct{
	Name string `required max:100`
	origin string
}

func StructTags(){
	t:=reflect.TypeOf(Animals{})
	field,_:=t.FieldByName("Name")
	fmt.Println(field.Tag)
}
