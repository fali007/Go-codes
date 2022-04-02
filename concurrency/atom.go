package main

import(
	"fmt"
	"math/rand"
	"sync"
)

type doc struct{
	Name string
	Age int
	Place string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func genArray(s int)*[]doc{
	a:=make([]doc,s)
	for i:=0;i<s;i++{
		a[i]=doc{RandStringBytes(5),i,RandStringBytes(6)}
	}
	return &a
}
func writeMap(m map[string]doc, k string, v doc,lock *sync.RWMutex){
	lock.Lock()
	defer lock.Unlock()
	m[k]=v
}

func dis(n int, s string, m map[string]doc,wg *sync.WaitGroup, lock *sync.RWMutex){
	defer wg.Done()
	d:=genArray(n)
	for _,i:=range *d{
		fmt.Printf("%+v %s \n", i, s)
		writeMap(m,i.Name,i, lock)
	}
}

func main(){
	wg:= &sync.WaitGroup{}
	lock:=&sync.RWMutex{}
	m:=make(map[string]doc)

	wg.Add(2)
	go dis(4000, "1",m, wg, lock)
	go dis(5000, "2",m, wg, lock)
	wg.Wait()
	fmt.Printf("\n\n\n%+v\n\n",m)
}