package main

import (
	"fmt"
)

func binarySearch(a []int,s,e,x int)int{
	m:=(s+e)/2
	fmt.Println(a,m,a[s:m],a[m:e])
	if(s<=e){
		fmt.Println(a)
		if a[m]==x{
			return a[m]
		}
		if a[m]>x{
			return binarySearch(a,s,m,x)
		}else{
			return binarySearch(a,m,e,x)
		}
	}else{
		return 0
	}
}

func getSubsetSum(a []int, x int) (int,int){
	m:=make(map[int]int)
	for _,val := range a{
		if m[val]!=0{
			return val,m[val]
		}else{
			m[x-val]=val
		}
	}
	return 0,0
}

func main(){
	arr:=[]int{3, 34, 4, 12, 5, 2}
	fmt.Println(getSubsetSum(arr,14))
	arr=[]int{2, 3, 4, 5, 12, 34,40}
	fmt.Println(binarySearch(arr,0,len(arr),3))
}