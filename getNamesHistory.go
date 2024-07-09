package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)
type UserId struct{
	UserId int `json:"UserId"`  // id of the user 
}

type SearchHistory struct{
	currentUserId int `json:"currentUserId"`;
	OtherUserId int `json:"OtherUserId"`;

}
func getNamesHistory(db *sql.DB, w http.ResponseWriter,r *http.Request){
	
	//@ simply esle chai  user_search_history table bata chai fetch garcha for given user_id or registration id 
var userId UserId;
var searchHistory SearchHistory;
var listOfNamesHistory []SearchHistory;
er:=json.NewDecoder(r.Body).Decode(&userId);
if er!=nil{
	fmt.Print("error while decoding from getnameshistory.go");
	return ;
}
	query :=`SELECT * from user_search_history where current_user_id =?`;

    rows, err := db.Query(query, userId);
	if err!= nil {
        fmt.Println("Error while querying database", err)
        return
    }
	for rows.Next() {
	rows.Scan(&searchHistory.currentUserId,&searchHistory.OtherUserId);
listOfNamesHistory = append(listOfNamesHistory, searchHistory);
	}
	errors:=json.NewEncoder(w).Encode(&listOfNamesHistory);
	if errors!=nil{
		fmt.Print("error wie encoding from getNamesHistory.go");
	}
}