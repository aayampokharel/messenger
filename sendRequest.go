package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)
type RequestSent struct{
	Sender_id int `json:"Sender_id"`
	Receiver_id int `json:"Receiver_id"`
	//Latest_message_timing time.Time `json:"Latest_message_timing"` 
	 // this is not required 
}
var chat_connectionsTableLock =&sync.Mutex{};

func sendRequest(db *sql.DB,_ http.ResponseWriter, r *http.Request) {
	
 //@ request sent after "send  request" button is pressed in frontend,  sends the request:{
    //@ Sender_id:1
	//@ Receiver_id:2
    //@(FUTURE IMPLEMENT)Latest_mesasge_timing:DateTime.now() (from flutter)
//@}
//@ THIS f(x) INPUTS THE DATA IN chat_connections table.and its time keeps getting updated with latest message around. 

var requestsent RequestSent;
response,err:=io.ReadAll(r.Body);
if(err!=nil){
	fmt.Print("error while reading the request sent to sendRequest function");
	return;
}

err=json.Unmarshal(response,&requestsent);
if err!=nil{
	fmt.Print("error while unmarshalling the request sent in sendRequest function");
	return;
}
chat_connectionsTableLock.Lock() 
defer chat_connectionsTableLock.Unlock();
var count int;

errors:=db.QueryRow("SELECT count(*) FROM chat_connections WHERE (sender_id=? AND receiver_id=?) OR (sender_id=? AND receiver_id=?) ",requestsent.Sender_id,requestsent.Receiver_id,requestsent.Receiver_id,requestsent.Sender_id).Scan(&count);
if errors!=nil{


	fmt.Print("error while reading count in sendrequest.go");
}



//currentTime:=time.Now(); //! THIS IS TO BE EDITED AS i NEED THIS FROM FLUTTER.
//latest_message_timing := currentTime.Format("2006-01-02 15:04:05")

// room_id:=generateRandomNumber();

if (count!=0 ||  requestsent.Sender_id == requestsent.Receiver_id){
	fmt.Print("invalid insertion while count !=0 or same value ")
return;
}
sender_id:=requestsent.Sender_id;
request_id:=requestsent.Receiver_id;
room_id:=generateRandomNumber();


query:="INSERT INTO chat_connections(room_id,sender_id,receiver_id) VALUES (?,?,?)";
_,er:= db.Exec(query,room_id,sender_id,request_id);
if er != nil {
	fmt.Print("cant insert as error occured in sendrequest. ")
	return;
}
 
}