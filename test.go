package main

import (
	"fmt"
)

func main(){
	temp:="12"
	fmt.Println(temp[:len(temp)-2]+temp[len(temp)-1:]+temp[len(temp)-2:len(temp)-1])
}