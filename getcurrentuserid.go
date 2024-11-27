package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var websocketLock = &sync.Mutex{}
var id int
var ch = make(chan int)

func getCurrentUserIdAndFlag(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//# response  :
	//@currentUserId["CurrentUserId"]=id;
	//@currentUserId["Flag"]=flag;
	var email EmailReceived

	var flag bool

	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		fmt.Print("error while decoding in getcurrentu serid.go")
		return
	}
	//websocketLock.Lock();
	query := "select registration_id,has_search_history from loginCredential_table where email=? "
	//@ for login arkai page ma run garna parcha else error aaucha as that email doesnot exist in the database.

	db.QueryRow(query, email.Email).Scan(&id, &flag) //! have to check for error else signin ko lagi arkai thau ma lagna parcha brother .
	var currentUserId = make(map[string]interface{})
	currentUserId["CurrentUserId"] = id
	currentUserId["Flag"] = flag
	go func() {
		ch <- id
	}()
	json.NewEncoder(w).Encode(currentUserId)
	//websocketLock.Unlock();

}
