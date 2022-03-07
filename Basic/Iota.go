package Basic

import "fmt"

const (
	x=iota 
	y=iota
	z=iota+2
	a=iota
	b=iota
	c=iota

)

const (
	d=iota
	e=iota
)

func Iota(){
	 
	fmt.Println(x)	//0
	fmt.Println(y)	//1
	fmt.Println(z)	//4
	fmt.Println(a)	//3
	fmt.Println(b)	//5
	fmt.Println(c)	//5
	fmt.Println(d)  //0
	fmt.Println(e)	//1
	 
	 
}




//to skip value in iota

const (
	f=iota
	_
	g=iota
	h=iota
	_
	k=iota

)
func IotaSkipValues(){
	 
	fmt.Println(f)	//0
	fmt.Println(g)	//2
	fmt.Println(h)	//3
	fmt.Println(k)	//5
	 
	 
}
