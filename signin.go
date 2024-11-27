package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)
type SignIn struct{ 
	//@This is used to receive the data from first sign in only in the form of :INSERT DIRECTLY IF SAME EMAIL THEN NO WRITING BECAUSE OF PRIMARY KEY ELSE WRITTEN . 
	//@ request:{
	//@ 	"Email":"razkkesh.yadav121212@gmail.com",
	//@    "Name":"Rakesh jhadab",
	//@    "Key":"NEET@@123"
	//@ }
	//@response: simply writing in the messenger_db;registration_id,email,display_name,password
	Email string `json:"Email"`
	Name string `json:"Name"`
	Key string `json:"Key"`
}
var insertLock =&sync.Mutex{};

func signIn(db *sql.DB,_ http.ResponseWriter, r *http.Request) {
	insertLock.Lock();
	defer insertLock.Unlock();
	var signInInfo SignIn
	response, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("Error while marshalling in the signIn handler function")
		return;
	}
	
	json.Unmarshal(response, &signInInfo)

	

	query:="INSERT INTO loginCredential_table(email,display_name,passwords) VALUES (?,?,?)";
	_,er:=db.Exec(query,signInInfo.Email,signInInfo.Name,signInInfo.Key);
	if er!=nil{
        fmt.Print("Error while insertion from signIn() function.");
		return;
    }


}
