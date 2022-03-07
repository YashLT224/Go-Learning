package Basic

import "fmt"

func ComplexDatattype(){
	var t complex64 =66+3i 
	fmt.Printf("%v %T \n",t,t)
	var t1 complex64 =4i 
	fmt.Printf("%v %T \n",t1,t1)

	var t2 complex64 =4 
	fmt.Printf("%v %T \n",t2,t2)


	 
	var a1 complex64=complex(5,20)
	fmt.Printf("%v %T \n",a1,a1)

	//fetch real and imaginary part from value
	var t3 complex64 =66+3i 
	fmt.Printf("%v %v \n",real(t3),imag(t3))

}


func ConstantDatatype(){
	const iw int =145
 
	fmt.Printf("%v %T \n",iw,iw)
}