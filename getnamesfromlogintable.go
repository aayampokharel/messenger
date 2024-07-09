package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)


func getNamesFromLoginTable(db *sql.DB,w http.ResponseWriter,_ *http.Request) {
	//@ gets names fromlogincredential_table should be used when flag=NO.
	var displayName string;
	var historyOfNames []string;

	 query:=`select display_name from loginCredential_table `;

	 rows, err:= db.Query(query);
	 if err!=nil{
		fmt.Print("error while selecting from getNames.go");
	 }

	 for rows.Next(){
	
		rows.Scan(&displayName);
historyOfNames=append(historyOfNames, displayName);
	 }
	er:= json.NewEncoder(w).Encode(historyOfNames);
if er!=nil{
	fmt.Print("error while encoding from getNames.go");
}

}