import 'package:flutter/material.dart';
import '../commons/sign_form.dart';

class SignUp extends StatelessWidget {
  final void Function(int) changeIndex;
  const SignUp({super.key, required this.changeIndex});

  @override
  Widget build(context) {
    return Column(
      children: [
        SignForm(),
        TextButton(
          child: const Text("go to splash"),
          onPressed: () => changeIndex(0)
          )
      ]
    );
  }
}
