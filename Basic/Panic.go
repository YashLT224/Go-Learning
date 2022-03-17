package Basic

import "fmt"
import "strconv"

func Panic(){
	s:="100"
	n,err:=strconv.ParseInt(s,10,64)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(n)
}

func Panic2(){
	fmt.Println("pre panic")
	Testpanic()
	fmt.Println("post panic")
}
func Testpanic(){
	fmt.Println("testpanic")
	panic("code breakdown")
	fmt.Println("return back")
}

/*output
pre panic
testpanic
panic: code breakdown
*/





func Recover(){
	fmt.Println("recover from panic")
	Recoverpanic()
	fmt.Println("post panic")
}
func Recoverpanic(){
	defer tryToRecover()
	fmt.Println("testpanic")
	panic("code breakdown")
	fmt.Println("I just panicked")
}

func tryToRecover(){
	if err:= recover();err!=nil{
		fmt.Println("Recovered from error",err)
	}
}
