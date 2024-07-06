import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/blueAppBar.dart';
import 'package:messenger/signIn/Login.dart';
import 'package:messenger/signIn/TextFieldInput.dart';
import 'package:http/http.dart' as http;

class MainBodySignIn extends StatelessWidget {
  TextEditingController emailController = TextEditingController();
  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  void signInToDatabase() async {
    //@ check for email,name validity, keep checking email after submit.
    Map<String, String> map = {
      "Email": emailController.text,
      "Name": nameController.text,
      "Key": passwordController.text
    };
    print(emailController.text);
    await http.post(Uri.parse("http://localhost:8080/signin"),
        body: json.encode(map));

    //# this has to be defined in homebody when submit clicked this can be done right inside onpressed thing .
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: blueAppBar("Create New account"),
      body: Center(
        child: Column(
          children: [
            TextFieldInput(
              emailController,
              labelText: "E-mail",
            ),
            //@ simply specifying whether its password is true or default=false where false=no obscure text, default=true obscure text.
            TextFieldInput(passwordController,
                labelText: "Password", forPassword: true),
            TextFieldInput(nameController, labelText: "Your Name"),
            ElevatedButton(
                onPressed: () async {
                  signInToDatabase(); //@ no dont wait for display of initial blank page. but when he tries to do anything just display thing like loading icon , as we cant do anything without his registration first. else wait for 3 seconds first. always.
                  Navigator.of(context).push(MaterialPageRoute(
                      builder: (context) => HomeBody(emailController
                          .text))); //@ represents the messager's body after signin.
                },
                child: const Text("Submit")),
            const Divider(
              thickness: 2,
              height: 100,
              indent: 100,
              endIndent: 100,
              color: Colors.grey,
            ),
            Center(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text("Already have an account?"),
                  ElevatedButton(
                      onPressed: () {
                        Navigator.of(context).pop();
                        Navigator.of(context).push(
                            MaterialPageRoute(builder: (context) => Login()));
                      },
                      child: const Text("login")),
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
