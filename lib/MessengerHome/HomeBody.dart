import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:messenger/MessengerHome/CardView.dart';
import 'package:messenger/MessengerHome/emptyCase.dart';
import 'package:messenger/blueAppBar.dart';
import 'package:messenger/insideChat/ChatBody.dart';
import 'package:http/http.dart' as http;
import 'package:web_socket_channel/web_socket_channel.dart';

class HomeBody extends StatefulWidget {
  ////
  ////
  ////
  ///ONLY FOR TESYING PURPOSE >
  String email = "srk_offficial@gmail.com";
  HomeBody(this.email);

  @override
  State<HomeBody> createState() => _HomeBodyState();
}

int? globalCurrentUserId = null;
int? globalOtherUserId = null;

class _HomeBodyState extends State<HomeBody> {
  final WebSocketChannel chatConnectionChannel =
      WebSocketChannel.connect(Uri.parse("ws://localhost:8080/wschatbody"));

  late final Stream chatStream;

  Future? _dataForHome;
  Future? homeHistory(String email) async {
    var currentIdResponse = await http.post(
        Uri.parse(
            "http://localhost:8080/getcurrentuserid"), //@ thsi also inserts the currentuserid for websocket ,
////
////
////
        //// FOR testing purpose only ....
        //  body: json.encode({"Email": widget.email}));
        body: json.encode({"Email": "srk_offficial@gmail.com"}));
    globalCurrentUserId = json.decode(currentIdResponse.body)["CurrentUserId"];
    var response =
        await http.post(Uri.parse("http://localhost:8080/homehistory"),
            //body: json.encode({"Email": widget.email}));
////
////
////
            //// FOR testing purpose only ....
            body: json.encode({"Email": "srk_offficial@gmail.com"}));

    var decodedList = json.decode(response.body);
    return decodedList;
  }

  @override
  void initState() {
    super.initState();
    chatStream = chatConnectionChannel.stream.asBroadcastStream();
    _dataForHome = homeHistory(widget.email);
  }

  @override
  void dispose() {
    // TODO: implement dispose
    super.dispose();

    chatConnectionChannel.sink.close();
  }

  forNewMessage(dynamic dataFromServer, Map<dynamic, dynamic> newMessage) {
    //!code changed where datafromserver ko type previously was list<dynamic> to dynamic
    int index = dataFromServer
        .indexWhere((element) => element["RoomId"] == newMessage["RoomID"]);

    if (index != -1) {
      dataFromServer.removeAt(index);
      dataFromServer.insert(0, newMessage);
    } else {
      dataFromServer.insert(0, newMessage);
    }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: blueAppBar("Messages"),
        body: FutureBuilder(
            future: _dataForHome,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                var dataFromServer = snapshot.data;
                if (dataFromServer == "NODATA") {
                  return const EmptyCase();
                }
//dataFromServer is always list of map here.
                return Column(
                  children: [
                    Expanded(
                      child: StreamBuilder(
                          stream: chatStream,
                          builder: (context, snapshot) {
                            dataFromServer as List;
                            if (snapshot.hasData) {
                              //  forNewMessage(dataFromServer, snapshot.data);
                            }
                            return ListView.builder(
                                itemCount: dataFromServer.length,
                                itemBuilder: (context, index) => CardView(
                                    dataFromServer[index]["RoomId"],
                                    dataFromServer[index]["SenderId"],
                                    dataFromServer[index]["ReceiverId"],
                                    dataFromServer[index]["DisplayName"],
                                    chatStream,
                                    chatConnectionChannel));
                          }),
                    )
                  ],
                );
              } else {
                return CircularProgressIndicator();
              }
            }),
      ),
    );
  }
}
