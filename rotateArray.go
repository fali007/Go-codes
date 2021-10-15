package main

import (
	"fmt"
)

func rotateArray(ary []int, rot int) []int{
	l:=len(ary)
	return append(ary[l-rot:],ary[:l-rot]...)
}

func rotate(ary []int, k int){
	l:=len(ary)
	for i:=0;i<k;i++{
		temp:=ary[l-1]
		for j:=l-1;j>0;j--{
			ary[j]=ary[j-1]
		}
		ary[0]=temp
	}
}

func main(){
	ary:=[]int{1,2,3,4,5,6,7,8,9}
	// fmt.Println(rotateArray(ary,3))
	rotate(ary,3)
	fmt.Println(ary)
	point(&ary)
}

func point(ary *[]int){
	for index,e:=range *ary{
		fmt.Printf("Value - %d Address - %p\n",e,&e)
	}
}