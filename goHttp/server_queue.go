package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Request struct{
	W http.ResponseWriter
	R *Endpoint
}

type Endpoint struct{
	Name           string        `json:"name"`
	Url            string        `json:"url"`
	Comment        string        `json:"comment"`
}

func GetJsonEncoding(o interface{})[]byte{
	json,err:=json.Marshal(o)
	if err!=nil{
		fmt.Println("Error marshalling json :",err)
	}
	return json
}

var reqChannel = make(chan Request)

func Process(w http.ResponseWriter, r *http.Request){
	var req Endpoint
	decoder := json.NewDecoder(r.Body)
	err:=decoder.Decode(&req)
	if err!=nil{
		fmt.Println(err)
	}
	// Index(Request{w,&req})
	reqChannel<-Request{w,&req}
}

func Index(inp Request){
	time.Sleep(time.Second*2)
	inp.W.WriteHeader(http.StatusOK)
	inp.W.Header().Set("Content-Type", "application/json")
	inp.W.Write(GetJsonEncoding(inp.R))
}

func StartChannel(){
	for req :=range reqChannel{
		go Index(req)
	}
}

func main(){
	go StartChannel()
	fmt.Println("Server started")
	r:=mux.NewRouter()
	r.HandleFunc("/index",Process).Methods("POST")
	http.ListenAndServe(":8080",r)
}