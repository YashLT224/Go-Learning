package Basic

import "fmt"
import "time"

func Goroutine1(){
	fmt.Println("Main Starts")
	go Hello(1)
	fmt.Println("Main Ends")
	
}

func Hello(n int){
	for i:=0;i<5;i++{
		fmt.Println("yash")
	}
}
//main goroutime pehle hi finish hogya hello wle goroutine se
/*output
Main Starts
Main Ends
*/


func Goroutine2(){
	fmt.Println("Main Starts")
	go HelloG(1)
	fmt.Println("Main Ends")
	time.Sleep(100 * time.Millisecond)
	
}

func HelloG(n int){
	for i:=0;i<5;i++{
		fmt.Println(n,"yashG")
	}
}
//main goroutime ko sleep pr dal diya to execute  hello wle goroutine 
/* output
Main Starts
Main Ends
yashG
yashG
yashG
yashG
yashG
*/




func Goroutine3(){
	fmt.Println("Main Starts")
	go HelloG1(1)
	go HelloG1(2)
	go HelloG1(3)
	go HelloG1(4)
	go HelloG1(5)
	fmt.Println("Main Ends")
	time.Sleep(100 * time.Millisecond)
	
}

func HelloG1(n int){
	for i:=0;i<5;i++{
		fmt.Println(n,"yashG1")
	}
}

/*output

5 yashG1
5 yashG1
5 yashG1
5 yashG1
5 yashG1
3 yashG1
3 yashG1
3 yashG1
3 yashG1
3 yashG1
4 yashG1
4 yashG1
4 yashG1
4 yashG1
4 yashG1
2 yashG1
1 yashG1
2 yashG1
2 yashG1
2 yashG1
1 yashG1
1 yashG1
1 yashG1
1 yashG1
2 yashG1

these goroutines are run parallely
*/



//race condition
func Goroutine4(){
	fmt.Println("Main Starts")

	msg :="Hello"
	go func(){
		fmt.Println(msg,"Goroutine")
	}()
	msg="world"
	fmt.Println("Main Ends")
	time.Sleep(100 * time.Millisecond)
}

//resolving race conditions
func Goroutine5(){
	fmt.Println("Main Starts")

	msg :="Hello"
	go func(m string){
		fmt.Println(m)
	}(msg)
	msg="world"
	fmt.Println("Main Ends")
	time.Sleep(100 * time.Millisecond)
}