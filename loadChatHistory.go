package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatConnection struct {
	RoomId     sql.NullInt64 `json:"RoomId"`
	SenderId   sql.NullInt64 `json:"SenderId"`
	ReceiverId sql.NullInt64 `json:"ReceiverId"`
}
type LoadChat struct {
	ReceiverId int    `json:"ReceiverId"`
	Chat       string `json:"Chat"`
}
type ForTime struct {
	SenderId   int            `json:"senderId"`
	ReceiverId int            `json:"receiverId"`
	LatestTime sql.NullString `json:"latestTime"`
}

func loadChatHistory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	//// This is used to load the chat history from the database after a user presses inside the other user.
	//@ no LOCK required .
	//@request ={ roomId:1234, OtherUserId:,currentUserId };
	//# response:

	var chatConnection ChatConnection

	err := json.NewDecoder(r.Body).Decode(&chatConnection)
	if err != nil {
		fmt.Print("error while decoding in loadchathistory.go")
		return
	}

	var forTime ForTime
	if chatConnection.RoomId.Valid {
		checkTimeQuery := `select sender_id,receiver_id,latest_time from chat_connections where room_id=?`
		db.QueryRow(checkTimeQuery, chatConnection.RoomId).Scan(&forTime.SenderId, &forTime.ReceiverId, &forTime.LatestTime)
		var mapJSON map[string]interface{} /// { PrivateChats, }
		var localPrivateChat LoadChat
		var listOfLocalPrivateChat []LoadChat

		if forTime.LatestTime.Valid {
			query := `select receiver_id,chat from private_chats_table where private_room_id=?`
			rows, er := db.Query(query, chatConnection.RoomId)

			if er != nil {
				fmt.Print("error while selecting rows inside the loadChatHistory.go")
				return
			}

			for rows.Next() {
				rows.Scan(&localPrivateChat.ReceiverId, &localPrivateChat.Chat)
				listOfLocalPrivateChat = append(listOfLocalPrivateChat, localPrivateChat)
			}
			//query = ` select sender_id,receiver_id from chat_connections where room_id=?`
			//db.QueryRow(query, chatConnection.RoomId).Scan(&senderId, &receiverId)
			mapJSON["PrivateChats"] = listOfLocalPrivateChat
			mapJSON["SenderId"] = forTime.SenderId
			mapJSON["ReceiverId"] = forTime.ReceiverId
			mapJSON["Status"] = "Accepted"

			er = json.NewEncoder(w).Encode(mapJSON)
			if er != nil {
				fmt.Print("error while encoding from loadchathistory.go")

			}
		} else {
			mapJSON["PrivateChats"] = listOfLocalPrivateChat
			mapJSON["SenderId"] = forTime.SenderId
			mapJSON["ReceiverId"] = forTime.ReceiverId
			mapJSON["Status"] = "NotAccepted"
			er := json.NewEncoder(w).Encode(mapJSON)
			if er != nil {
				fmt.Print("error while encoding from loadchathistory.go")
			}

		}
	} else {
		//@ yo chai from search bata aauda ho where we wont fetch roomid .so roomid=null yo send garda chai . so tesbela i have to check row empty huna ni sakcha cause there might not be connection established . as eeuta completely new manche bhayo bhane there wont be any friend request exchanged.
		var roomId int
		forEmptyRoomQuery := `select room_id,sender_id,receiver_id,latest_time from chat_connections where room_id=?`
		db.QueryRow(forEmptyRoomQuery, chatConnection.RoomId).Scan(&roomId, &forTime.SenderId, &forTime.ReceiverId, &forTime.LatestTime)

	}

	//# select * froom connection where sender_id=? and receiver=? OR sender
	//# null aaema we understand nobody has sent the request .
	//# answer aae ma somebody has already sent the request .

	rows, er := db.Query("select p1.receiver_id,chat from private_chats_table p1 inner join chat_connections c1 on p1.private_room_id=c1.room_id where c1.latest_time is not null and p1.private_room_id=? ;")
	if er != nil {
		fmt.Print("error while selecting rows inside the loadChatHistory.go")
		return
	}
	defer rows.Close()
	var loadChat LoadChat
	var listOfChats []LoadChat
	for rows.Next() {
		err := rows.Scan(&loadChat.ReceiverId, &loadChat.Chat)
		if err != nil {
			fmt.Print("error while scanning from loadChatHistory.go \n")
			return
		}
		listOfChats = append(listOfChats, loadChat)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(listOfChats); err != nil {
		fmt.Print("there is error while encoding inside loadchathistory.go")
	}
}
