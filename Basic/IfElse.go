package Basic


import "fmt"

func IfElse(){
	for a:=0;a<=5;a++{
		if(a%2==0){
			fmt.Println("Even no",a)
		}else{
			fmt.Println("odd no",a)
		}
	}
}


func IfElseIf() {  
	fmt.Print("Enter text: ")  
	var input int  
	fmt.Scanln(&input)  
	if (input < 0 || input > 100) {  
	   fmt.Println("Please enter valid no")  
	} else if (input >= 0 && input < 50  ) {  
	   fmt.Println(" Fail")  
	} else if (input >= 50 && input < 60) {  
	   fmt.Println(" D Grade")  
	} else if (input >= 60 && input < 70  ) {  
	   fmt.Println(" C Grade")  
	} else if (input >= 70 && input < 80) {  
	   fmt.Println(" B Grade")  
	} else if (input >= 80 && input < 90  ) {  
	   fmt.Println(" A Grade")  
	} else if (input >= 90 && input <= 100) {  
	   fmt.Println(" A+ Grade")  
	}
}  


func Switch(){
	fmt.Println("Enter the Input")
	var input int
	fmt.Scanln(&input)
	fmt.Println(input)
	switch(input){
	case 1:
		fmt.Println("you press 1")
	case 2:
		fmt.Println("you press 2")
	case 3:
		fmt.Println("you press 3")
	case 4:
		fmt.Println("you press 4")
	case 5:
		fmt.Println("you press 5")
	default:
		fmt.Println("invalid input")

	}

}

func FallThrough(){
	input:=4
	 
	switch(input){
	case 1:
		fmt.Println("you press 1")
	case 2:
		fmt.Println("you press 2")
	case 3:
		fmt.Println("you press 3")
	case 4:
		fmt.Println("you press 4")
		fallthrough 					//just next case bhi run ho jayega
	case 5:
		fmt.Println("you press 5")
	default:
		fmt.Println("invalid input")

}
}


func Break(){
	 
		var  a int = 1  
		for a < 10{  
		   fmt.Print("Value of a is ",a,"\n")  
		   a++;  
		   if a > 5{  
			  /* terminate the loop using break statement */  
			  break;  
		   }  
		} 
}