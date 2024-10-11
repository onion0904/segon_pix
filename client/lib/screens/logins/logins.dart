import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'splash.dart';
import 'sign_in.dart';
import 'sign_up.dart';
import 'create_user.dart';

class Logins extends HookWidget {
  const Logins({super.key});

  @override
  Widget build(context) {
    final index = useState(0);

    final loginUIList = [
      const Splash(),
      SignIn(changeIndex: (int a) {
        index.value = a;
      }),
      SignUp(changeIndex: (int a) {
        index.value = a;
      }),
      const CreateUser()
    ];

    return Scaffold(
      resizeToAvoidBottomInset: false,
      appBar: AppBar(
        title: const Text("Login"),
        backgroundColor: Colors.orange,
      ),
      body: Center(
        child: loginUIList[index.value],
      )
    );
  }
}
