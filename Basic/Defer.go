package Basic

import "fmt"

func Defer(){
	defer fmt.Println("i")
	fmt.Println("Love")
	fmt.Println("Programming")

	 

}

//sabse pehla defer sabse last me execute hoga

func Defercase(){
	defer fmt.Println("i")
	defer fmt.Println("love")
	defer fmt.Println("programming")      
	//output
	//programming
	//love
	//i


}