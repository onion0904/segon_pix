import 'package:flutter/material.dart';

class SignIn extends StatelessWidget {
  final void Function(int) changeIndex;
  const SignIn({super.key, required this.changeIndex});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: TextButton(
        onPressed: () => changeIndex(0),
        child: const Text("Go to Sign Up"),
      ),
    );
  }
}
