package Basic

import "fmt"

func StringIndex(){
	s:="yash Verma"
	s2:="is a boy"
	fmt.Printf("%v\n",string(s[1]))
	fmt.Printf("%v\n",s+s2)
	b:=[]byte(s)

	fmt.Printf("%v\n",b)
}


func Rune(){
var s rune='a'
fmt.Printf("%d, %T\n",s,s)
}

func Runeexp(){
	r:='A'
	fmt.Printf("%d, %T\n",r,r)			// 97,int32

	a:="A"
	fmt.Printf("%s, %T\n",a,a)		


}


