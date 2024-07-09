package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func getNamesFromUserSearchHistoryTable(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var currentUserId UserId;
	err:=json.NewDecoder(r.Body).Decode(&currentUserId);
	if err!=nil{
		fmt.Print("error handling while decoding inside getNamesFromUserSearchHistory.go");
		return;
	}


	query:=`select l1.display_name from user_search_history s1
inner join loginCredential_table l1 on s1.other_user_id=l1.registration_id where s1.current_user_id=? `;//@select names after joining the logincredential table .  
rows,er:=db.Query(query,currentUserId.UserId);
if er!=nil{
	fmt.Print("cant select a problem in db has arrived ");

}
var temp string;
var tempList []string;

for rows.Next(){
rows.Scan(&temp);
tempList = append(tempList, temp);
}
err=json.NewEncoder(w).Encode(tempList);
if err!=nil{
	fmt.Print("error while encoding in getNamesFromUserSearchHistory.go")
	return;
}
}
