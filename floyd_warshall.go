package main

import(
	"fmt"
	"math"
)

func Get2DArry(l,b int, val interface{})[][]interface{}{
	dyn:=[][]interface{}{}
	for i:=0;i<l;i++{
		temp:=[]interface{}{}
		for j:=0;j<b;j++{
			temp=append(temp,val)
		}
		dyn=append(dyn,temp)
	}
	return dyn
}

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func floydWarshall(a [][]int)[][]int{
	k:=len(a)
	for i:=0;i<k;i++{
		for j:=0;j<k;j++{
			for l:=0;l<k;l++{
				a[j][l]=min(a[j][l],a[j][i]+a[i][l])
			}
		}
		fmt.Printf("\nOutput after %d iteration : %v \n", i+1, a)
	}
	return a
}

func main(){
	INF:=math.MaxInt32
	a:=[][]int{{0,   5,  INF, 10},
        {INF,  0,  3,  INF},
        {INF, INF, 0,   1},
        {INF, INF, INF, 0}}
    fmt.Println("final output - ",floydWarshall(a))
    a=[][]int{{0,   3,  INF, 7},
        	  {8,  0,  2,  INF},
        	  {5, INF, 0,   1},
        	  {2, INF, INF, 0}}
    fmt.Println("final output - ",floydWarshall(a))
}