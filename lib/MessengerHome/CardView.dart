import 'package:flutter/material.dart';
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/insideChat/ChatBody.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class CardView extends StatelessWidget {
  int RoomId;
  int SenderId;
  int ReceiverId;
  String DisplayName;
  Stream chatStream;
  WebSocketChannel chatChannel;
  CardView(this.RoomId, this.SenderId, this.ReceiverId, this.DisplayName,
      this.chatStream, this.chatChannel);

  @override
  Widget build(BuildContext context) {
    return Card(
      child: ListTile(
          title: Text(
            DisplayName,
            style: const TextStyle(fontSize: 20),
          ),
          leading: const Icon(Icons.person),
          subtitle: const Text("sample last message "),
          onTap: () {
            Navigator.of(context).push(MaterialPageRoute(
                builder: (context) => ChatBody(RoomId, SenderId, ReceiverId,
                    DisplayName, chatStream, chatChannel)));
          }),
    );
  }
}
