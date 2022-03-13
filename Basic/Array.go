package Basic

import "fmt"

func ArrrayDeclaration(){
	var arr[4]int=[4]int{1,2,3,4}
	fmt.Println(arr)
	fmt.Printf("Element at [%d] index is=%d\n",3,arr[1])
}

func DynamicArray(){
	var arr1 =[...]int{1,2,3,4,5,6}
	fmt.Println(arr1)
}

func ArrayLength(){
	var arr1 =[...]int{1,2,3,4,5,6}
	fmt.Println(len(arr1))
}

func ReplaceValueAtIndex(){
	var arr1 =[...]int{1,2,3,4,5,6}

	arr1[2]=20
	fmt.Println(arr1)

	//value at particular index
	fmt.Println(arr1[2])
}

func Copy_OneArray_TO_Another(){
	var arr1 =[...]int{1,2,3,4,5,6}
     arr2:=arr1
	fmt.Println(arr2)

	arr1[0]=10    			// see mene arr1 me change kiya but arr2 me change nahi hua
	fmt.Println(arr1) 
	fmt.Println(arr2)
}

func PrintSubarray(){
	  arr1 :=[...]int{1,2,3,4,5,6}
       
	fmt.Println(arr1[1:3])   //thrd index not considered
	fmt.Println(arr1[1:])   //from first inex to length
}



func MultidimensionalArray(){
	var matrix[3][3]int=[3][3] int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(matrix)
	fmt.Println(matrix[1][1])
	fmt.Println(len(matrix[0]))
matrix[0][0]=35
fmt.Println(matrix)

}

func MultiDemArray(){
	var matrix[3][3]int=[3][3] int{{1,2,3},{4,5,6},{7,8,9}}
var i,j int
	for  i=0;i<3;i++{
	for j=0;j<3;j++{
fmt.Print(matrix[i][j] )
fmt.Print(" ")
}
fmt.Println()
	}
	 

}