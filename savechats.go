package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"sync"
// )

// type Chat struct {
// 	RoomId int `json:"RoomId"`
// 	SenderName string `json:"SenderName"`
// 	Chats string `json:"Chats"`

// }

// var saveChatsLock =&sync.Mutex{};
// func saveChats(db *sql.DB,chatStructure Chat){
// /// THIS FUNCTION WORKS LIVE with GOROUTINE.
// 	//// this is used for storing live chats in database.
// 	//@ is is required for sending the current chats into the database and writing there . THIS CAN HAPPEN PARALLELY. USE GOROUTINES.
// 	//@ request: {RoomId: XXXXXX,SenderName:"a",Chat: "heyy!! whats up? "}
// 	//@response: returns the current chats in the room.

// 	query :="INSERT INTO private_chats_table(private_room_id,receiver_name,chat) VALUES (?,?,?)";
// 	saveChatsLock.Lock();
// 	defer saveChatsLock.Unlock();
// roomId:=chatStructure.RoomId;
// senderName:=chatStructure.SenderName;
// chat:=chatStructure.Chats;

//  _,err:=db.Exec(query  ,roomId ,senderName, chat );
//  if err!=nil{
// 	fmt.Print("error while inserting into private_chats_table from storeChats ")
// 	return ;
//  }

// }

