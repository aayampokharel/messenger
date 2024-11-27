import 'dart:async';
import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/insideChat/ChatText.dart'; // Assuming your ChatText widget is in this file

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

  Future<Map> getChatHistoryList() async {
    var response = await http.post(
        Uri.parse("http://localhost:8080/chathistory"),
        body: json.encode({"RoomId": widget.RoomId}));
    return json.decode(response.body);
  }

  @override
  void initState() {
    super.initState();
    getChatHistoryList().then((value) => print(value));
  }

  @override
  Widget build(BuildContext context) {
    if (widget.ReceiverId == globalCurrentUserId) {
      globalOtherUserId =
          widget.SenderId; // Handle the current/other user logic
    } else {
      globalOtherUserId = widget.ReceiverId;
    }

    return Scaffold(
      appBar: AppBar(title: Text(widget.DisplayName)), // Your custom AppBar
      body: Stack(
        children: [
          FutureBuilder(
            future: getChatHistoryList(),
            builder: (context, futureSnapshot) {
              if (futureSnapshot.hasData) {
                return StreamBuilder(
                  stream: widget.chatStream.asBroadcastStream(),
                  builder: (context, streamSnapshot) {
                    return Column(
                      children: [
                        Expanded(
                          child: ListView.builder(
                            itemCount:
                                futureSnapshot.data!['PrivateChats'].length,
                            itemBuilder: (context, index) {
                              // Extract the individual chat message
                              var chatMessage =
                                  futureSnapshot.data!['PrivateChats'][index];
                              bool isReceiver = chatMessage['ReceiverId'] ==
                                  widget.ReceiverId;

                              // Conditional rendering based on ReceiverId
                              return Align(
                                alignment: isReceiver
                                    ? Alignment
                                        .centerRight // If message is for the current user
                                    : Alignment
                                        .centerLeft, // If message is from the current user
                                child: Padding(
                                  padding: const EdgeInsets.all(8.0),
                                  child: ChatText(chatMessage['Chat']),
                                ),
                              );
                            },
                          ),
                        ),
                      ],
                    );
                  },
                );
              } else {
                return const Center(child: CircularProgressIndicator());
              }
            },
          ),

          // Send button and text field
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
                  onPressed: () {}, // You can implement send functionality here
                  icon: const Icon(Icons.send),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
