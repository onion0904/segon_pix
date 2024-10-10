import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import '../commons/sign_form.dart';

class SignIn extends HookWidget {
  final void Function(int) changeIndex;
  const SignIn({super.key, required this.changeIndex});

  @override
  Widget build(BuildContext context) {

    return Column(
      children: [
        SignForm(),
          TextButton(
          onPressed: () => changeIndex(0),
          child: const Text("Go to Sign Up"),
        ),
      ]
    );
  }
}
