package Basic

import "fmt"

func ScanInput(){
	fmt.Println("ENter any no")
	var num int
	fmt.Scanln(&num)
	fmt.Println(num)

}

func ArrayInput(){
fmt.Println("Enter the sie of Array")
var size int
fmt.Scanln(&size)
a := make([]int, size)

for i:=0;i<size;i++{
 
	fmt.Scanln(&a[i])
}

fmt.Println(a)

 
   }
