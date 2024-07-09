package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)


func getNamesFromLoginTable(db *sql.DB,w http.ResponseWriter,r *http.Request) {
	//@ gets names from logincredential_table should be used when flag=NO.
	var displayName string;
	var currentUserId UserId;
	var historyOfNames []string;

	err:=json.NewDecoder(r.Body).Decode(&currentUserId);
	if err!=nil{
		fmt.Print("error while decoding from getnamesfromlogintable.go")
	}

	 query:=`select display_name from loginCredential_table where registration_id !=?`;//@ !=kina bhanda it just helps in things like getting all names except self as search ma afnai name search garna milena.

	 rows, err:= db.Query(query,currentUserId.UserId);
	 if err!=nil{
		fmt.Print("error while selecting from getNames.go");
	 }

	 for rows.Next(){
	
		er:=rows.Scan(&displayName);
		if er!=nil{
			fmt.Print("error while scanningg from getnamesfromlogintable.go")
			return;
		}
historyOfNames=append(historyOfNames, displayName);
	 }
	er:= json.NewEncoder(w).Encode(historyOfNames);
if er!=nil{
	fmt.Print("error while encoding from getNames.go");
}

}