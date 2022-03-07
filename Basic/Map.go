package Basic

import "fmt"

func Map(){
	empSal:=make(map[string]int)  //declaration
	empSal=map[string]int{   //initialization
		"Neha":20000,
		"Yash":10000,
		"Tarun":150000,
		"Vamika":12000,

	}
	fmt.Println(empSal)

	empSal1:=map[string]int{   //declaraton with initialization
		"Neha":20000,
		"Yash":10000,
		"Tarun":150000,
		"Vamika":12000,

	}
	fmt.Println(empSal1)
	fmt.Println(len(empSal1))  		//length
	fmt.Println(empSal1["Yash"])		//getvalue

	empSal1["Vamika"]=150					//reassign value
	fmt.Println(empSal1)					//	map[Neha:20000 Tarun:150000 Vamika:150 Yash:10000]
}

func Put_Delete_in_map(){
	empSal1:=map[string]int{   //declaraton with initialization
		"Neha":20000,
		"Yash":10000,
		"Tarun":150000,
		"Vamika":12000,

	}
	fmt.Println(empSal1)				//map[Neha:20000 Tarun:150000 Vamika:12000 Yash:10000]
	empSal1["japneet"]=500 				//add in existing map
	fmt.Println(empSal1)			//map[Neha:20000 Tarun:150000 Vamika:12000 Yash:10000 japneet:500]

	delete(empSal1,"Neha")
	fmt.Println(empSal1)

}
func ContainsKey_thenPrint(){
	empSal:=map[string]int{
		"yash":1000,
		"dev":2000,
	}
	 ankit,flag:=empSal["ankit"]
	fmt.Println(ankit,flag)  //0 false

	//I dont want to use ankit variable  how can i skip
	_,isPresent:=empSal["ankit"]
	fmt.Println(isPresent)        // false
}

func Copy_Map_Trap(){
	empSal1:=map[string]int{   //declaraton with initialization
		"Neha":20000,
		"Yash":10000,
		"Tarun":150000,
		"Vamika":12000,

	}
	emp:=empSal1
	fmt.Println(empSal1)		
	fmt.Println(emp)
	empSal1["Rajesh"]=10     //I added rajesh in empSal1 but it also reflects in emp
	fmt.Println(empSal1)		//map[Neha:20000 Rajesh:10 Tarun:150000 Vamika:12000 Yash:10000]
	fmt.Println(emp)			//map[Neha:20000 Rajesh:10 Tarun:150000 Vamika:12000 Yash:10000]
	

}