package main

import (
	"fmt"
)
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func min_platforms(arr, dep []float32)int{
	platforms:=0
	l:=len(arr)
	for i:=0;i<l;i++{
		temp:=1
		fmt.Printf("arrival : %f",arr[i])
		for j:=i;j<l;j++{
			if dep[i]>dep[j]{
				temp+=1
				fmt.Printf("\ttrain in platform %f, platforms required - %d", dep[j], max(platforms,temp))
			}
		}
		fmt.Printf("\n max platforms %d \n", platforms)
		platforms=max(platforms,temp)
	}
	return platforms
}

func main(){
	arr:=[]float32{9.00, 9.40, 9.50, 11.00, 15.00, 18.00}
	dep:=[]float32{19.10, 12.00, 11.20, 11.30, 19.00, 20.00}
	fmt.Println(min_platforms(arr,dep))
}