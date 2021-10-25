package main

import(
	"fmt"
	"time"
)

func print(s string, c chan string){
	for i:=0;i<5;i++{
		c<-s
		time.Sleep(1*time.Second)
	}
	close(c)
	// only sender blocks when its finnished sending all the messages to channel
}

func main(){
	c:=make(chan string)
	go print("Hello", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

// in a non buffered channel we can receive and send only in different go routines
// a channel will send only when there is a receiver ready to receive the value
// In such cases to read from the same channel in same go routine, we need to create a buffered channel with size