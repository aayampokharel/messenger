import 'package:flutter/material.dart';

class ChatText extends StatelessWidget {
  String Chat;
  ChatText(this.Chat);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: ConstrainedBox(
          constraints:
              const BoxConstraints(maxWidth: 200, minWidth: 50, minHeight: 50),
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Container(
                color: const Color.fromARGB(255, 54, 225, 255),
                child: Text(Chat)),
          )),
    );
  }
}
