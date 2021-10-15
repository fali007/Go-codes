package main

import (
	"fmt"
	"math"
)

func posDiff(a,b int)int{
	if a>b{
		return a-b
	}
	return b-a
}

func closestSumTwo(arr []int, x int)(int, int){
	s:=0
	e:=len(arr)-1
	val:=math.MaxInt32
	var a,b int
	for e>s{
		if posDiff(arr[s]+arr[e],x)<val{
			a,b=arr[s],arr[e]
			val=posDiff(arr[s]+arr[e],x)
		}
		if arr[s]+arr[e]>x{
			e--
		}else{
			s++
		}
	}
	return a,b
}

func closestSum(arr []int, x int) (int,int){
	l:=len(arr)
	var a,b int
	val:=math.MaxInt32
	for i:=0;i<l;i++{
		for j:=i;j<l;j++{
			if posDiff(arr[i]+arr[j],x)<val{
				val=posDiff(arr[i]+arr[j],x)
				a,b=arr[i],arr[j]
			}
		}
	}
	return a,b
}

func main(){
	arr:=[]int{1, 3, 4, 7, 10}
	fmt.Println(closestSum(arr,15))
	fmt.Println(closestSumTwo(arr,15))
}