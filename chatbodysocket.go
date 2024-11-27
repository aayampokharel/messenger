package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
)

type ChatMessage struct{
	RoomId int `json:"RoomId"`
	ReceiverId int `json:"ReceiverId"` 
	Chat string `json:"Chat"` //@ not chat message . 
	 
}
type IdRoom struct{
	Id int `json:"Id"`
	Room int `json:"Room"`
}

var insertDbLock =&sync.Mutex{};


var lockBeforeInsertion=&sync.Mutex{};
var mapForRoomId=make(map[*websocket.Conn]int);
var  mapForSenderReceiverId= make(map[ *websocket.Conn ] int);

func deleteConnectionFromMap( address *websocket.Conn ) {

	delete(mapForSenderReceiverId,address);
    

}

func chatBodySocket(db *sql.DB, w http.ResponseWriter, r *http.Request){
	///initially , connect to server after fetching from database the registration_id . , etc should be done while connecting . or do the global thing in another while fetching the database in another file . 
	 
	//@ this actually should get the integer room id from the client.
	//@first SEND : {Id:1 , Room: 123456***}

	//@ request:to connect to the ws , NOTHING ! ;
	//#requst:{RoomID,Receiverid, ChatMessage};
	//#response : same as above ....

	
chatHandshake,err:=	websocket.Accept(w,r ,&websocket.AcceptOptions{
		OriginPatterns: []string{"*"},	})
	
	if err!=nil{
		fmt.Print("print couldnot connect to chatbodysocket:problem");
		return;
		
		} 
		defer chatHandshake.Close(websocket.StatusNormalClosure,"websocket closed normally");
		
		mapForSenderReceiverId[chatHandshake]=<-ch;//! search for better approach later in the future. 
		fmt.Print(mapForSenderReceiverId,"❌❌❌❌\n\n");
//websocketLock.Unlock();
//mapForRoomId[chatHandshake]=chatRoom.Room; ////this has to be at the LAST.why here ?? i mean after reading from websocket.  

for{//! there was also a solution to store connections and reuse it later on . when connected . 
	_,msgByte,err:=chatHandshake.Read(r.Context());
	//@ here msgByte should be of form ChatMessage;
	fmt.Print(" \n\n called \n\n\n\n");
	if err!=nil{
		fmt.Print("error while reading from chathandshake.read()");
		deleteConnectionFromMap(chatHandshake);
		return; 
	}

	var chatMessageSent ChatMessage;
	if errors:=json.Unmarshal(msgByte,&chatMessageSent);errors !=nil{
		fmt.Print("problem while unmarshalling ");
		return ;
	}
	fmt.Print("\n\n 😂😂😂\n\n")
	fmt.Print(chatMessageSent);
	fmt.Print(mapForSenderReceiverId);
	fmt.Print("\n\n 😂😂😂\n\n")
	go updateDatabase(db, chatMessageSent);//! lock ? 
 messageSendJSON,JSONerr:=json.Marshal(chatMessageSent);
 
 if JSONerr!=nil{
	fmt.Print("problem while marshalling inside chatbody.go ");
 }

 

 for mapKey,mapValue:=range mapForSenderReceiverId{//@ this is a required if in amy case both sender and receiver are there . either opening chat or not.in any case it is valid as same message ma messengerbody ma and chatbody ma dubai ma ma insert garirako chu . 
	if mapValue==chatMessageSent.ReceiverId{
			//? logic slightly change garna parcha as in database i have to replace the sender with receiver in private_chats_table . 
		mapKey.Write(r.Context(),websocket.MessageText,[]byte(messageSendJSON));
		break;
		
		
		}
	}
	// return;
 
//  for _,mapValue := range mapForRoomId{
	
	
// 	///when receiver is not opening messenger this is used and this is okay to get delayed.
// 	///instead of sending message this should be used to store the message in the  db above as well as well as here 
// 	if mapValue==chatMessageSent.RoomId{
	
// 		//mapKey.Write(context.Background(),websocket.MessageText,[]byte(messageSendJSON));
// //@ ahile dont send push notification just do other stuffs .
// 		//! push notification 
// 		//! save in db . 
	
// 	//// maintain a queue to store latest live 50 messages in frontend . 
// 	}
//  }

 }


	
 
}

func updateDatabase(db *sql.DB,chatMessageSent ChatMessage ){
	insertDbLock.Lock();
	defer insertDbLock.Unlock();
query:=`insert into private_chats_table(private_room_id, receiver_id, chat) values(?,?,?)`;
_,err:=db.Exec(query,chatMessageSent.RoomId,chatMessageSent.ReceiverId,chatMessageSent.Chat);
if err!=nil{
	fmt.Print("error while inserting into db from chatbodysocket updatadatabase function");
}


}
	
