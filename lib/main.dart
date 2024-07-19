import 'package:flutter/material.dart';
import 'package:messenger/blueAppBar.dart';
import 'package:messenger/signIn/SignIn.dart';

void main() {
  runApp(const Messenger());
}

//@ Messenger():SignIn() in home label lincha.
class Messenger extends StatelessWidget {
  const Messenger({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: SignIn(),
    );
  }
}

//@ SignIn() :blueAppBar() within which eeuta string lincha for appbar and makes the color of the appbar blue.
class SignIn extends StatelessWidget {
  const SignIn({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: MainBodySignIn(), //@ consists of email,password and submit .
    );
  }
}
