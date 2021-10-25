package main

import (
	"fmt"
	"time"
)

/*

go routine wont run past main
Everything ends when instruction exits after main
It continues to run in background

*/

func main(){
	go print("Hello")
	go print("World")
	time.Sleep(10*time.Second)
	// fmt.Scanln()
}

func print(value string){
	for true{
		fmt.Println(value)
		time.Sleep(2*time.Second)
	}
}