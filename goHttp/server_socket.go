package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Endpoint struct{
	Name           string        `json:"name"`
	Url            string        `json:"url"`
	Comment        string        `json:"comment"`
}

type Request struct{
	Conn *websocket.Conn
	R *Endpoint
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
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("upgrade failed: ", err)
        return
    }
	decoder := json.NewDecoder(r.Body)
	err=decoder.Decode(&req)
	if err!=nil{
		fmt.Println(err)
	}
	reqChannel<-Request{conn,&req}
}

func Index(req Request){
    for{
    	mt, message, _ := req.Conn.ReadMessage()
    	err:=req.Conn.WriteMessage(mt,[]byte("Reading message"))
    	fmt.Println(string(message))
    	err=req.Conn.WriteMessage(mt,message)
    	err=req.Conn.WriteMessage(mt,[]byte("End-closing connection"))
    	req.Conn.Close()
    	if err!=nil{
    		break
    	}
    }
}

func StartChannel(){
	for req :=range reqChannel{
		go Index(req)
	}
}

func main(){
	go StartChannel()
	r:=mux.NewRouter()
	r.HandleFunc("/index",Process)
	http.ListenAndServe(":8080",r)
}