package Basic

import "fmt"
import "time"

func receiver(ch chan int){
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func Channel(){
	fmt.Println("Channel")
	ch:=make(chan int)
	go receiver(ch)
	ch <- 42	
	ch <- 42
	close(ch)
	time.Sleep(100 * time.Millisecond)
}


//iterating over channel
func receiver1(ch chan int){
	for{
		i,ok:=<-ch
		if ok==false{
			fmt.Println(i,ok,"loop break")
			break
		}else{
			fmt.Println(i,ok)
		}
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func Channel1(){
	fmt.Println("Channel")
	ch:=make(chan int)
	go receiver1(ch)
	ch <- 42	
	ch <- 43
	ch <- 44
	ch <- 45
	close(ch)
	time.Sleep(100 * time.Millisecond)
}



func receiver2(ch chan int){
	for val:=range ch{
		fmt.Println(val)
	}
		 
	
	fmt.Println(<-ch)
	fmt.Println("Channel has been closed by sender")
}
func Channel2(){
	fmt.Println("Channel")
	ch:=make(chan int)
	go receiver2(ch)			//udhar go routine parallel chal rha hai aur yha channel me new new values ja rhi hai jo go routine read kar raha hai 
	ch <- 42	
	ch <- 43
	ch <- 44
	ch <- 45
	close(ch)
	time.Sleep(100 * time.Millisecond)
}





//bidirectional channel example

func receiver3(ch chan int){
		fmt.Println(<-ch)
		ch <- 12345
}
func Channel3(){
	fmt.Println("Channel")
	ch:=make(chan int)
	go receiver3(ch)			//udhar go routine parallel chal rha hai aur yha channel me new new values ja rhi hai jo go routine read kar raha hai 
	ch <- 42	
	fmt.Println(<-ch)
	close(ch)
	time.Sleep(100 * time.Millisecond)
}


//unidirectional channel    sirf sender bheje aur rceiver receive kare receive send na kare


func receiver4(ch <-chan int){    //this signifies that you can only read the channel not write the channel
	fmt.Println(<-ch)
	//ch <- 12345				//error aagya
}
func Channel4(){
fmt.Println("Channel")
ch:=make(chan int)
go receiver4(ch)			 
ch <- 42	
fmt.Println(<-ch)
close(ch)
time.Sleep(100 * time.Millisecond)
}




func receiver5(ch chan<- int){    //this signifies that you can only write the channel not read the channel
	// fmt.Println(<-ch)	//error aagya
	ch <- 12345			
}
func Channel5(){
fmt.Println("Channel")
ch:=make(chan int)
go receiver5(ch)			 
ch <- 42	
fmt.Println(<-ch)
close(ch)
time.Sleep(100 * time.Millisecond)
}






func receive72(ch chan int){
	fmt.Println(<-ch)
	 
}
func BufferedChannel(){
	fmt.Println("Channel")
	ch:=make(chan int,2)
	go receive72(ch)			 
	ch <- 42			//42 is read by goroutine 43 is in buffer  .this is not a  nonblocking call  error nai aayega no deadlock situation
	ch <- 43
	 
	close(ch)
	time.Sleep(100 * time.Millisecond)
}







func Channel6(){
	fmt.Println("Go Channels")
	ch:=make(chan string)

	go routine("nonblocking 1",ch)
	go routine("nonblocking 2",ch)
 
	msg:=<- ch
	fmt.Println(msg)

	msg2:=<- ch
	fmt.Println(msg2)

}

func routine( s string , c chan string){
	for i:=0;i<2;i++{
		fmt.Println(i,s)
	}
	c <-s+"done"
	
}