package main

import (
	"fmt"
)

func fibnocci(a int) int{
	if a<2{
		return 1
	}
	return fibnocci(a-1)+fibnocci(a-2)
}

func worker(input <-chan int, output chan<- int){
	for num:=range input{
		output<-fibnocci(num)
	}
	close(output)
}

func main(){
	input:=make(chan int,50)
	output:=make(chan int,50)

	go worker(input,output)
	go worker(input,output)
	go worker(input,output)
	go worker(input,output)

	for i:=0;i<50;i++{
		input<-i
	}
	close(input)

	for i := range output{
		fmt.Println(i)
	}
}