package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailReceived struct{
	Email string `json:"Email"`
}

type HomeHistory struct{
	RoomId int 
	SenderId int
	ReceiverId int
	DisplayName string
}

func homeHistory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//@request :{Email:"aayam.pokharel@gmail.com"}
	//@loads the home history of the messenger user. 
var emailReceived EmailReceived;
	var homeHistoryToBeSent HomeHistory ;
	var listOfHomeHistoryMap []HomeHistory;

	err:=json.NewDecoder(r.Body).Decode(&emailReceived);
	defer r.Body.Close();
	if err!=nil{
		fmt.Print("cant decode the json email in homehistory.go");
return ;
	}
	
	query:=`  

	select 
	c1.room_id ,
	c1.sender_id, 
	c1.receiver_id,
	l2.display_name
	 from 
	   chat_connections c1
	inner join 
	 loginCredential_table l1
	on c1.sender_id=l1.registration_id or
	 c1.receiver_id=l1.registration_id
	inner join 
	 loginCredential_table l2 
	 on l2.registration_id=c1.sender_id or  l2.registration_id=c1.receiver_id
	where l1.email=? and  l2.email!=?
`

rows,err:=db.Query(query,emailReceived.Email,emailReceived.Email); 

if err!=nil{
	fmt.Print("problem while querying from homeHistory page.");
	return ;
}
 defer func(){
	if rows !=nil{

		rows.Close();
	}
 }();
for rows.Next(){

	
	if err=rows.Scan(&homeHistoryToBeSent.RoomId,&homeHistoryToBeSent.SenderId,&homeHistoryToBeSent.ReceiverId,&homeHistoryToBeSent.DisplayName);err!=nil{
		fmt.Print("problem while scanning in homehistroy.go")
		return;
	}
	listOfHomeHistoryMap=append(listOfHomeHistoryMap,homeHistoryToBeSent );
	}
	if(listOfHomeHistoryMap==nil){
		fmt.Print("empty map because of no data");
		if err=json.NewEncoder(w).Encode("NODATA");err!=nil{
			fmt.Print("cant send the json file to the frontend.");
		}
	}else{
		
		if err=json.NewEncoder(w).Encode(listOfHomeHistoryMap);err!=nil{
			fmt.Print("cant send the json file to the frontend.");
		}
	}
}
