package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type Room struct {
	RoomId int `json:"RoomId"`

}
type LoadChat struct{
	ReceiverId int `json:"ReceiverId"`
	Chat string `json:"Chat"`

}
func loadChatHistory(db *sql.DB, w http.ResponseWriter,r *http.Request) {
	 //// This is used to load the chat history from the database after a user presseeds inside the other user. 
	 //@ no LOCK required . 
     //@request ={ roomId:1234}
	 //# response: 
	 requestByte,err:=io.ReadAll(r.Body);
	 if err!=nil{
		fmt.Print("error while reading body inside loadChatHistory");
		return ;
	 }
	 var room Room;
	 json.Unmarshal(requestByte,&room);////change the table name 
	rows,err:= db.Query("SELECT receiver_id,chat FROM private_chats_table WHERE private_room_id=?;",room.RoomId);
if err!=nil{
	fmt.Print("error while selecting rows inside the loadChatHistory.go");
	return ;
}
defer rows.Close();
var loadChat LoadChat;
var listOfChats []LoadChat; 
for rows.Next(){
	err:=rows.Scan(&loadChat.ReceiverId,&loadChat.Chat);
	if err!=nil{
	fmt.Print("error while scanning from loadChatHistory.go \n");
	return;
	}
listOfChats=append(listOfChats, loadChat);
}


 w.Header().Set("Content-Type","application/json");

 if err:=json.NewEncoder(w).Encode(listOfChats);err!=nil{
	fmt.Print("there is error while encoding inside loadchathistory.go");
 }
}