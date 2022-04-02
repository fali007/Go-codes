package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Request struct{
	Name 	string	`json:"name"`
	Url 	string 	`json:"url"`
	Place 	string	`json:"comment"`
}

type Resp struct{
	Status string	`json:"message"`
	Id 	   int 		`json:"id"`
}

func GetJsonEncoding(o interface{})[]byte{
	json,err:=json.Marshal(o)
	if err!=nil{
		fmt.Println("Error marshalling json :",err)
	}
	return json
}

func worker(inp chan Request, out chan Resp, id int){
	for i :=range inp{
		resp, _ := http.Post("http://localhost:8080/index", "application/json",bytes.NewBuffer(GetJsonEncoding(i)))
		r,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("err")
		}else{
			out<-Resp{string(r),id}
		}
	}
}

func main(){
	number:=10
	inp:=make(chan Request,number)
	out:=make(chan Resp,number)

	threads:=8

	for i:=0;i<threads;i++{
		go worker(inp,out,i)
	}

	for i:=0;i<number;i++{
		inp<-Request{"feliz","abcd.com","kannur"}
	}
	close(inp)

	for i:=range out{
		fmt.Printf("%+v\n", i)
	}
}