package Basic
import (  
	"fmt"  
	"strconv"  
	"reflect"
)  
func Conversion(){
	var a int32=35
	var b int64=int64(a)
    //int 32 to int 64
	fmt.Printf("%T %v\n",b,b);

	fmt.Println("----------------------------------------------------------------------");

   //float to int conversion
	c:=2.5
	fmt.Printf("%T %v\n",c,c);
	var d int64=int64(c)
	fmt.Printf("%T %v\n",d,d);


	fmt.Println("----------------------------------------------------------------------");
	//int to String conversion
	var str int=111
	fmt.Printf("%T %v\n",str,str);
	var r string=strconv.Itoa(str)
	fmt.Printf("%T %q\n",r,r);
	fmt.Println("----------------------------------------------------------------------");



 //String to integer conversion

var strr string ="111"
intVar, err := strconv.Atoi(strr)
fmt.Println(intVar, err, reflect.TypeOf(intVar))
}