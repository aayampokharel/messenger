package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatConnection struct {
	RoomId     sql.NullInt64 `json:"RoomId"`
	SenderId   int           `json:"SenderId"`
	ReceiverId int           `json:"ReceiverId"`
}
type ChatConnectionPointerInt struct {
	RoomId     *int64 `json:"RoomId"`
	SenderId   int    `json:"SenderId"`
	ReceiverId int    `json:"ReceiverId"`
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

	//@request ={ roomId:1234, OtherUserId:,currentUserId };
	//# response:{
	// #mapJSON["PrivateChats"] = listOfLocalPrivateChat
	//# mapJSON["SenderId"] = forTime.SenderId
	//# mapJSON["ReceiverId"] = forTime.ReceiverId
	//# mapJSON["Status"] = "Accepted"
	//# other status : "RequestPending"
	//#NoConnection
	//#
	//
	//
	//}

	var chatConnection ChatConnection
	var chatConnectionPointer ChatConnectionPointerInt

	err := json.NewDecoder(r.Body).Decode(&chatConnectionPointer)

	if err != nil {
		fmt.Print("error while decoding in loadchathistory.go")
		return
	}
	//
	chatConnection.ReceiverId = chatConnectionPointer.ReceiverId
	chatConnection.SenderId = chatConnectionPointer.SenderId
	//
	if chatConnectionPointer.RoomId != nil {
		chatConnection.RoomId = sql.NullInt64{Int64: *chatConnectionPointer.RoomId, Valid: true}

	} else {
		chatConnection.RoomId = sql.NullInt64{
			Int64: 0, //ignored by db but still written for clarity.
			Valid: false}

	}
	fmt.Print(chatConnection, " \n\n\n\n\n\n\n\n")
	fmt.Print("yes inside loadchathistory")
	//!!!!!!!!!!!!!!!!!!!!!!!!!!

	mapJSON := make(map[string]interface{}) /// { PrivateChats, }
	var listOfLocalPrivateChat []LoadChat
	var localPrivateChat LoadChat
	//
	//
	//@ to determine whether the roomId is null or not . null when request sent from search ,
	if chatConnection.RoomId.Valid {
		var forTime ForTime
		//? checkTimeQuery is here because  i want to know the time,also senderand receiver  id is extracted for displaying who sent request to whom at first.
		//
		//
		checkTimeQuery := `select sender_id,receiver_id,latest_time from chat_connections where room_id=?`
		db.QueryRow(checkTimeQuery, chatConnection.RoomId.Int64).Scan(&forTime.SenderId, &forTime.ReceiverId, &forTime.LatestTime)
		//
		//#ambiguous as room_id ko satta same tala use gareko can be used.
		//? to check roomid=yes and latesttime if null: means request sent but not responded.
		//? if not null when there is roomid:request has been accepted and there might be chats available.
		fmt.Print("\n\n\n\n\n\n", forTime, chatConnection.RoomId.Int64, "\n\n\n ", "ues")

		if forTime.LatestTime.Valid {
			query := `select receiver_id,chat from private_chats_table where private_room_id=?`
			rows, er := db.Query(query, chatConnection.RoomId.Int64)

			if er != nil {
				fmt.Print("error while selecting rows inside the loadChatHistory.go")
				return
			}

			for rows.Next() {
				rows.Scan(&localPrivateChat.ReceiverId, &localPrivateChat.Chat)
				listOfLocalPrivateChat = append(listOfLocalPrivateChat, localPrivateChat)
			}

			mapJSON["PrivateChats"] = listOfLocalPrivateChat
			mapJSON["SenderId"] = forTime.SenderId
			mapJSON["ReceiverId"] = forTime.ReceiverId
			mapJSON["Status"] = "Accepted"
			fmt.Print("\n\n\n\n", mapJSON, "\n\n\n\n")
			er = json.NewEncoder(w).Encode(mapJSON)
			if er != nil {
				fmt.Print("error while encoding from loadchathistory.go")

			}
		} else {
			//? time =null
			//? here condition of no input in table cannot occur as here only request from the homelist can arrive here as they only have roomid.
			//
			mapJSON["PrivateChats"] = listOfLocalPrivateChat
			mapJSON["SenderId"] = forTime.SenderId
			mapJSON["ReceiverId"] = forTime.ReceiverId
			mapJSON["Status"] = "RequestPending"
			fmt.Print("\n\n\n\n", mapJSON, "\n\n\n\n")
			er := json.NewEncoder(w).Encode(mapJSON)
			if er != nil {
				fmt.Print("error while encoding from loadchathistory.go")
			}

		}
	} else {
		//@ yo chai from search bata aauda ho where we wont fetch roomid .so roomid=null yo send garda chai .
		//@ so tesbela i have to check row empty huna ni sakcha cause there might not be connection table ma related insertion. as eeuta completely new manche bhayo bhane there wont be any friend request exchanged.
		var nullTime sql.NullTime
		var extractedRoomId, extractedSenderId, extractedReceiverId int
		query := `select room_id,sender_id,receiver_id,latest_time from chat_connections where sender_id=? and receiver_id=? or receiver_id=? and sender_id=? `
		oneRow := db.QueryRow(query, chatConnection.SenderId, chatConnection.ReceiverId, chatConnection.SenderId, chatConnection.ReceiverId).Scan(&extractedRoomId, &extractedSenderId, &extractedReceiverId, &nullTime)
		//@ below condition checks if A profile is searching someone who is not connected. if i search hrithwik roshan, i dont have any chat message neither i have any friend request sent. so no history in chat_connecitons .

		if oneRow == sql.ErrNoRows {
			fmt.Print("\n\n\n\n\n hello brother\n\n\n\n\n\n")
			fmt.Print("\n\n\n\n\n\n", extractedReceiverId, "\n\n\n\n\n\n")
			fmt.Print("\n\n\n\n\n\n", extractedSenderId, "\n\n\n\n\n\n")
			//@ for unknown names clicked from searchbar.
			//@ for such thing , sender and receiverId will be 0 as no entry , so no registration.
			mapJSON["PrivateChats"] = listOfLocalPrivateChat
			mapJSON["SenderId"] = extractedSenderId
			mapJSON["ReceiverId"] = extractedReceiverId
			mapJSON["Status"] = "NotConnected"
			fmt.Print("\n\n\n\n", mapJSON, "\n\n\n\n")
			errors := json.NewEncoder(w).Encode(mapJSON)
			if errors != nil {
				fmt.Print("error while errnoRows fromloadchathistory.go")
			}
			return
		} else {
			//@ if single row displayed this is run as this means there is already a request, or chat history pre existing.
			//
			if nullTime.Valid {
				//@ valid means there is time value (!=null), so there can be messages.

				query := `select receiver_id,chat from private_chats_table where private_room_id=?`
				rows, er := db.Query(query, extractedRoomId)

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
				mapJSON["SenderId"] = extractedSenderId
				mapJSON["ReceiverId"] = extractedReceiverId
				mapJSON["Status"] = "Accepted"
				fmt.Print("\n\n\n\n", mapJSON, "\n\n\n\n")
				er = json.NewEncoder(w).Encode(mapJSON)
				if er != nil {
					fmt.Print("error while encoding from loadchathistory.go")

				}

			} else {
				//@ request sent but yet to be accepted.
				mapJSON["PrivateChats"] = listOfLocalPrivateChat
				mapJSON["SenderId"] = extractedSenderId
				mapJSON["ReceiverId"] = extractedReceiverId
				mapJSON["Status"] = "RequestPending"
				fmt.Print("\n\n\n\n", mapJSON, "\n\n\n\n")
				errors := json.NewEncoder(w).Encode(mapJSON)
				if errors != nil {
					fmt.Print("error while noHistory from loadchathistory.go")
				}
				return
			}

		}

		//# select * froom connection where sender_id=? and receiver=? OR sender
		//# null aaema we understand nobody has sent the request .
		//# answer aae ma somebody has already sent the request .

	}
}
