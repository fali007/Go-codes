package main

import (
	"fmt"
	// "io/ioutil"
	"encoding/json"
	"github.com/gorilla/websocket"
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
		conn,_, err := websocket.DefaultDialer.Dial("ws://localhost:8080/index", nil)
		if err!=nil{
			fmt.Println("error")
		}
		conn.WriteMessage(websocket.TextMessage,GetJsonEncoding(i))
		for{
			_,resp,err:=conn.ReadMessage()
			if err!=nil{
				conn.Close()
				break
			}else{
				out<-Resp{string(resp),id}
			}
		}
		
	}
}

func main(){
	number:=100
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