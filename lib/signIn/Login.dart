import 'package:flutter/material.dart';
import 'package:messenger/MessengerHome/HomeBody.dart';
import 'package:messenger/blueAppBar.dart';
import 'package:messenger/main.dart';
import 'package:messenger/signIn/SignIn.dart';
import 'package:messenger/signIn/TextFieldInput.dart';

class Login extends StatelessWidget {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: blueAppBar("Login to your existing account"),
      body: Center(
        child: Column(
          children: [
            TextFieldInput(
              _emailController,
              labelText: "E-mail",
            ),
            //@ simply specifying whether its password is true or default=false where false=no obscure text, default=true obscure text.
            TextFieldInput(_passwordController,
                labelText: "Password", forPassword: true),

            ElevatedButton(
                onPressed: () {
                  Navigator.of(context).push(MaterialPageRoute(
                      builder: (context) => HomeBody(_emailController
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
                  Text("Want to create New Account?"),
                  ElevatedButton(
                      onPressed: () {
                        Navigator.of(context).pop();
                        Navigator.of(context).push(MaterialPageRoute(
                            builder: (context) => const SignIn()));
                      },
                      child: const Text("LOG IN ")),
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
