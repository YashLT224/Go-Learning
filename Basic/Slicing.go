package Basic

import "fmt"
func Slicing(){
	var arr[] int=[]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Println(len(arr))

	//capacity
	fmt.Println(cap(arr))


	//ex
	var arr2[] int=make([]int,2,4)
	fmt.Println(len(arr2))			//2
	fmt.Println(cap(arr2))			//4

}

func Slicing_append(){
	var arr2[] int =[]int{1,2,3,4}
	fmt.Println(arr2)			//[1 2 3 4]
	arr2 =append(arr2,8)
	fmt.Println(arr2)			//[1 2 3 4 8]

	arr21:=append(arr2,10)
	fmt.Println(arr21)			//[1 2 3 4 8 10]


	arr211:=append(arr2,arr21...)
	fmt.Println(arr211)				//[1 2 3 4 8 1 2 3 4 8 10]
}

func SlicingTrap(){
	var arr[] int=[]int{1,2,3,4,5}
	 
	 arr2:=arr
	 fmt.Println(arr2)
	 arr[0]=9
	 fmt.Println(arr)		//[9 2 3 4 5]
	 fmt.Println(arr2)  	//[9 2 3 4 5]

}