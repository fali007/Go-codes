/*
bottom up approcah of programming

basic:
	1- dont reuse a variable in a function

*/

package main

import(
	"fmt"
)

// func name(){
// 	firstName:=""
// 	lastName:=""
// 	fullname:=firstName+ " "+lastName
// }

func add(x int) func(y int) int{
	return func(y int) int{
		return x+y
	}
}

func factorial(n int)int{
	if n==1{
		return 1
	}
	return n*factorial(n-1)
}

func main(){
	add1:=add(1)
	res:=add1(3)
	fmt.Printf("\n %+v, %+v \n", add1, res)
	fmt.Println(factorial(5))
}