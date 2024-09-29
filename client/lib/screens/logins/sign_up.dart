import 'package:flutter/material.dart';

class SignUp extends StatelessWidget {
  final void Function(int) changeIndex;
  const SignUp({super.key, required this.changeIndex});

  @override
  Widget build(context) {
    return Center(
      child: TextButton(
        child: const Text("go to splash"),
        onPressed: () => changeIndex(0)
      )
    );
  }
}
