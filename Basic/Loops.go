package Basic

import "fmt"

func Forloop(){

	for a:=0;a<10;a++{
		fmt.Println("Yash")
	}
}


func Continue(){

	for a:=0;a<10;a++{
		if(a==2){
			continue
		}else{
			fmt.Println(a)
		}
		
	}
}



func ForRange(){
	nums:=[]int{1,2,3}
	sum:=0
	for _,value := range nums {// "_ " is to ignore the index  
	sum += value  
 }  
 fmt.Println("sum:", sum)  
}