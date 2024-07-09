package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ForHistory struct{
	CurrentUserId int `json:"CurrentUser"`;
	OtherUserId int `json:"OtherUser"`;
}

func storeHistory(db *sql.DB, w http.ResponseWriter, r *http.Request){
//# stores people who have been searched for reference anbother time . while current is running i can maintain local as well instead of fetching .ani frontend ma fetch garna chahincha. ani tei list locally update garne after that. 
	//@ request {CurrentUserId,OtherUserId}
	//@ response: nothing .
	var forHistory ForHistory;
err:= json.NewDecoder(r.Body).Decode(&forHistory);
if err!=nil{
	fmt.Print("error while storeHistory is running from storehistory.go ");
}
query:=`insert into user_search_history(current_user,other_user) values (?, ?)`;

_,er:=db.Exec(query,forHistory.CurrentUserId,forHistory.OtherUserId);
if er!=nil{
	fmt.Print("error while inserting in storehistory.go");
}
}