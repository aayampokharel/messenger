package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type DecodingJSON struct {
	ReceieverId int     `json:"ReceiverId"`
	SenderId    int     `json:"SenderId"`
	Method      string  `json:"Method"`
	LatestTime  *string `json:"LatestTime"`
	RoomId      int     `json:"RoomId"`
}
type DecodedJSON struct {
	ReceieverId int       `json:"ReceiverId"`
	SenderId    int       `json:"SenderId"`
	Method      string    `json:"Method"`
	LatestTime  time.Time `json:"LatestTime"`
	RoomId      int       `json:"RoomId"`
}

// type InsertMethod struct {
// 	LatestTime time.Time `json:"LatestTime"`
// }

var chat_connectionsTableLock = &sync.Mutex{}

func sendRequest(db *sql.DB, _ http.ResponseWriter, r *http.Request) {

	//@ request sent after "send  request" button is pressed in frontend,  sends the request:{
	//@ Sender_id:1
	//@ Receiver_id:2
	//@time from flutter as null(insert)/time.now()(update)

	var decodingJSON DecodingJSON
	var decodedJSON DecodedJSON
	if err := json.NewDecoder(r.Body).Decode(&decodingJSON); err != nil {
		fmt.Print("Error while decoding from sendRequest.go")
		return
	}
	fmt.Print("\n")
	fmt.Print(decodedJSON)
	fmt.Print("\n")
	fmt.Print(decodingJSON)

	decodedJSON.SenderId = decodingJSON.SenderId
	decodedJSON.ReceieverId = decodingJSON.ReceieverId
	decodedJSON.Method = decodingJSON.Method
	decodedJSON.RoomId = decodingJSON.RoomId
	if decodingJSON.LatestTime == nil {

		decodedJSON.LatestTime = time.Time{}

	} else { //@ if there is time.
		layout := time.RFC3339Nano
		timeTimeFormat, err := time.Parse(layout, *decodingJSON.LatestTime)
		if err != nil {
			fmt.Print("error while converting to time.Time from sendRequest.go")
			return
		}
		decodedJSON.LatestTime = timeTimeFormat
	}
	fmt.Print("\n helllo brother \n")
	chat_connectionsTableLock.Lock()
	defer chat_connectionsTableLock.Unlock()
	if decodedJSON.Method == "INSERT" {
		roomId := generateRandomNumber()
		fmt.Print("\n\n")
		fmt.Print(roomId, decodedJSON)
		fmt.Print("\n\n")
		query := `insert into chat_connections(room_id,sender_id,receiver_id) values (?,?,?)`
		_, err := db.Exec(query, roomId, decodedJSON.SenderId, decodedJSON.ReceieverId)
		if err != nil {
			fmt.Print("error while executing in sendRequest.go")
			return
		}
		//@ no time required as default is null .
	}

	if decodedJSON.Method == "UPDATE" {
		query := `update chat_connections set latest_time=? where room_id=?`
		_, er := db.Exec(query, decodedJSON.LatestTime, decodedJSON.RoomId)
		if er != nil {
			fmt.Print("error while executing query in sendrequest.go in updating")
			return
		}

	}

}
