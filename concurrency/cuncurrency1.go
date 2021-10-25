package main

import(
	"fmt"
	"time"
	"sync"
)

func print(s string){
	for i:=0;i<5;i++{
		fmt.Println(i,s)
		time.Sleep(1*time.Second)
	}
}

func main(){
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		print("Hello")
		wg.Done()
	}()

	wg.Wait()
}