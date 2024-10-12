import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'splash.dart';
import 'sign_in.dart';
import 'sign_up.dart';
import 'create_user.dart';
import 'package:go_router/go_router.dart';

class Logins extends HookWidget {
  const Logins({super.key});

  @override
  Widget build(context) {
    final index = useState(3);

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
        child:  Column(
          mainAxisAlignment: MainAxisAlignment.center,
          // crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            loginUIList[index.value],
            Padding(
                padding: const EdgeInsets.all(p),
                child: ElevatedButton(
                    onPressed: () {
                      //TODO サーバに送信

                      context.go("/hub");
                    },
                    child: const Text("決定")))
          ],
        )
      )
    );
  }
}
