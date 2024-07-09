package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func getNamesFromUserSearchHistoryTable(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//@ this should be run when flag=true or 1 from client to get the history of names . 
	var currentUserId UserId;
	err:=json.NewDecoder(r.Body).Decode(&currentUserId);
	if err!=nil{
		fmt.Print("error handling while decoding inside getNamesFromUserSearchHistory.go");
		return;
	}


	query:=`select l1.display_name from user_search_history s1
inner join loginCredential_table l1 on s1.other_user_id=l1.registration_id where s1.current_user_id=? `;//@select names after joining the logincredential table .  
rows,er:=db.Query(query,currentUserId);//returning strung //@ returning string and receiving through nt ? 

if er!=nil{
	fmt.Print("cant select a problem in db has arrived ");
return; 
}
var temp string;
var tempList []string;

for rows.Next(){
er:=rows.Scan(&temp);
if er!=nil{
	fmt.Print("error while sdcanning from getNamesfromusersearchhistory.go");
	return;
}
tempList = append(tempList, temp);
}
err=json.NewEncoder(w).Encode(tempList);
if err!=nil{
	fmt.Print("error while encoding in getNamesFromUserSearchHistory.go")
	return;
}
}
