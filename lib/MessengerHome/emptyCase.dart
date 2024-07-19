import 'package:flutter/material.dart';

class EmptyCase extends StatelessWidget {
  const EmptyCase({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Icon(
          Icons.person_add_alt_1_sharp,
          size: 100,
        ),
        Text(
            "You are yet to connect to someone.Use our search bbar above to connect to new friends."),
      ],
    );
  }
}
