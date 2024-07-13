package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func getNamesFromLoginTable(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//@ gets names from logincredential_table should be used when flag=NO.
	//# requires registration id .
	//! such comparision or business logic should be done in backend

	var currentUserId UserId
	var idAndName IdAndName
	var historyOfNames []IdAndName

	err := json.NewDecoder(r.Body).Decode(&currentUserId)
	if err != nil {
		fmt.Print("error while decoding from getnamesfromlogintable.go")
	}

	query := `select registration_id,display_name from loginCredential_table where registration_id !=?` //@ !=kina bhanda it just helps in things like getting all names except self as search ma afnai name search garna milena.

	rows, err := db.Query(query, currentUserId.UserId)
	if err != nil {
		fmt.Print("error while selecting from getNames.go")
	}

	for rows.Next() {

		er := rows.Scan(&idAndName.OtherUserId, &idAndName.DisplayName)
		if er != nil {
			fmt.Print("error while scanningg from getnamesfromlogintable.go")
			return
		}
		historyOfNames = append(historyOfNames, idAndName)
	}
	er := json.NewEncoder(w).Encode(historyOfNames)
	if er != nil {
		fmt.Print("error while encoding from getNames.go")
	}

}
