// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// )

// func selectChatsFromDB(db *sql.DB, chatConnection ChatConnection,localPrivateChats) {
// 	query := `select receiver_id,chat from private_chats_table where private_room_id=?`
// 	rows, er := db.Query(query, chatConnection.RoomId)

// 	if er != nil {
// 		fmt.Print("error while selecting rows inside the loadChatHistory.go")
// 		return
// 	}

// 	for rows.Next() {
// 		rows.Scan(&localPrivateChat.ReceiverId, &localPrivateChat.Chat)
// 		listOfLocalPrivateChat = append(listOfLocalPrivateChat, localPrivateChat)
// 	}
// 	//query = ` select sender_id,receiver_id from chat_connections where room_id=?`
// 	//db.QueryRow(query, chatConnection.RoomId).Scan(&senderId, &receiverId)
// 	mapJSON["PrivateChats"] = listOfLocalPrivateChat
// 	mapJSON["SenderId"] = forTime.SenderId
// 	mapJSON["ReceiverId"] = forTime.ReceiverId
// 	mapJSON["Status"] = "Accepted"

// 	er = json.NewEncoder(w).Encode(mapJSON)
// 	if er != nil {
// 		fmt.Print("error while encoding from loadchathistory.go")

//		}
//	}
package main
