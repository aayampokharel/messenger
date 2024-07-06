import 'dart:async';
import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/blueAppBar.dart';
import 'package:messenger/insideChat/ChatText.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:http/http.dart' as http;

class ChatBody extends StatefulWidget {
  int RoomId;
  int SenderId;
  int ReceiverId;
  String DisplayName;
  Stream chatStream;
  ChatBody(this.RoomId, this.SenderId, this.ReceiverId, this.DisplayName,
      this.chatStream);

  @override
  State<ChatBody> createState() => _ChatBodyState();
}

class _ChatBodyState extends State<ChatBody> {
  late StreamController _streamController;
  TextEditingController chatController = TextEditingController();
  Future<List> getChatHistoryList() async {
    var response = await http.post(
        Uri.parse("http://localhost:8080/chathistory"),
        body: json.encode({"RoomId": widget.RoomId}));
    return json.decode(response.body);
  }

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    getChatHistoryList().then((value) => print(value));
  }

  @override
  Widget build(BuildContext context) {
    if (widget.ReceiverId == globalCurrentUserId) {
      globalOtherUserId = widget
          .SenderId; // as globalotherUserid always should be opposite of globalCurrentuserid and this should be true whenever the currentuser opens any chat.
    } else {
      globalOtherUserId = widget.ReceiverId;
    }
    return Scaffold(

        ///FETCHING PROCESS AT THE BEGINNING FOR HISTORY IS MISSING.
        appBar: blueAppBar(widget.DisplayName),
        body: Stack(
          children: [
            FutureBuilder(
                future: getChatHistoryList(),
                builder: (context, futureSnapshot) {
                  if (futureSnapshot.hasData) {
                    return StreamBuilder(
                        stream: widget.chatStream.asBroadcastStream(),
                        builder: (context, streamSnapshot) {
                          //if (snapshot.data["roomId"] == widget.RoomId) {
                          //! add the value at the last of messagelistview  .tei ho nothing else to be done.
                          //! ONLY CHECK THAT IF THE USER'S MESSSAGE DISPLAY SIDE. TO CHECK .
                          //~ a gentle reminder is that : also handle if NO DATA AVAILABLE IN DATABASE
                          // }
                          return Column(
                            children: [
                              Expanded(
                                child: ListView.builder(
                                    itemCount: futureSnapshot.data!.length,

                                    ///required data here .
                                    itemBuilder: (context, index) {
                                      return ChatText(
                                          futureSnapshot.data![index]["Chat"]);
                                    }),
                              ),
                            ],
                          );
                        });
                  } else {
                    return const Center(child: CircularProgressIndicator());
                  }
                }),
            //! use stack to have chats at back and send button at front.

            Align(
              alignment: Alignment.bottomCenter,
              child: Row(
                children: [
                  Expanded(
                    child: Container(
                      width: 500,
                      height: 100,
                      child: TextField(
                        controller: chatController,
                      ),
                    ),
                  ),
                  IconButton(
                    onPressed: () {
                      widget.chatStream.add();
                    }, //! ESKO KAAM BHANEKO SIMPLY SEND MATRA GARNE HO .NOTHING ELSE.
                    icon: const Icon(Icons.send),
                  ),
                ],
              ),
            ),
          ],
        ));
  }
}
