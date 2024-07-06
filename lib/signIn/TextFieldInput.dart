import 'package:flutter/material.dart';

Widget TextFieldInput(TextEditingController _controller,
    {required String labelText, bool forPassword = false}) {
  return Padding(
    padding: const EdgeInsets.all(8.0),
    child: Container(
        width: 500,
        child: TextField(
          controller: _controller,
          decoration: InputDecoration(
              border: const OutlineInputBorder(
                  borderRadius:
                      BorderRadius.all(Radius.circular(20))), //! to set border.
              labelText: labelText),
          obscureText: forPassword,
        )),
  );
}
