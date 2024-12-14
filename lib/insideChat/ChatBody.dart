import 'dart:async';
import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/insideChat/ChatText.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
// Assuming your ChatText widget is in this file

class ChatBody extends StatefulWidget {
  int RoomId;
  int SenderId;
  int ReceiverId;
  String DisplayName;
  Stream chatStream;
  WebSocketChannel chatChannel;
  ChatBody(this.RoomId, this.SenderId, this.ReceiverId, this.DisplayName,
      this.chatStream, this.chatChannel);

  @override
  State<ChatBody> createState() => _ChatBodyState();
}

class _ChatBodyState extends State<ChatBody> {
  ScrollController scrollController =
      ScrollController(); // Controller for ListView
  late StreamController _streamController;
  TextEditingController chatController = TextEditingController();
  var ListOfMessages;

  Future<Map> getChatHistoryList() async {
    var response = await http.post(
        Uri.parse("http://localhost:8080/chathistory"),
        body: json.encode({"RoomId": widget.RoomId}));
    return json.decode(response.body);
  }

  @override
  void initState() {
    super.initState();
    getChatHistoryList().then((val) => ListOfMessages = val[
        "PrivateChats"]); ////i can also check if list is null initially if yes tala inside futurebuilder i can also assign and next time null hunna . this can also be done .
  }

  void _scrollToBottom() {
    if (ListOfMessages.isNotEmpty) {
      Future.delayed(const Duration(milliseconds: 100), () {
        scrollController.jumpTo(scrollController.position.maxScrollExtent);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    var MapOfMessages;

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
                // print("💦💦\n ${futureSnapshot.data}\n");
                // ListOfMessages = futureSnapshot.data![
                //   "PrivateChats"]; //for setstate to work for new messages .

                return StreamBuilder(
                  stream: widget.chatStream.asBroadcastStream(),

                  ///backend bata message is sent to another user only , eeutai user ko ma locally list ma direct store huncha to save bandwidth. so maile "hi " pathae bhane direct add mero list ma but will send this msg to another through this Stream and will display in real time .
                  builder: (context, streamSnapshot) {
                    if (streamSnapshot.hasData) {
                      print("\n\n\n data aaeko cha !!!!\n\n");
                      print("\n\n\n😂😂😂");
                      print("Stream snapshot data: ${streamSnapshot.data}");

                      print("\n\n\n😂😂😂");
                      // if (int.parse((json
                      //         .decode(streamSnapshot.data["RoomId"])
                      //         .toString())) ==
                      //     widget.RoomId) {
                      //   print(streamSnapshot.data);
                      //   print("😂😂😂\n\n");
                      // }
                    }

                    return Column(
                      children: [
                        Expanded(
                          child: ListView.builder(
                            controller: scrollController,
                            itemCount: ListOfMessages.length,
                            itemBuilder: (context, index) {
                              // Extract the individual chat message
                              var chatMessage = ListOfMessages[index];
                              bool isReceiver = chatMessage['ReceiverId'] ==
                                  globalOtherUserId;

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
                  onPressed: () {
                    //    print(futureSnapshot.data);
                    Map<String, dynamic> map = {
                      "RoomId": widget.RoomId,
                      "ReceiverId": globalOtherUserId,
                      "Chat": chatController.text.trim().toString(),
                    };
                    print(globalOtherUserId);
                    print(globalCurrentUserId);
                    print(map);
                    widget.chatChannel.sink.add(json.encode(map));
                    setState(
                      () {
                        ListOfMessages.add({
                          "ReceiverId": map["ReceiverId"],
                          "Chat": map["Chat"],
                        });
                        _scrollToBottom();
                      },
                    );
//  chatStream..sink.add()
                  }, // You can implement send functionality here
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
