import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'splash.dart';
import 'sign_in.dart';
import 'sign_up.dart';
import 'create_user.dart';
import 'package:go_router/go_router.dart';
import '../commons/button.dart';

class Logins extends HookWidget {
  const Logins({super.key});

  @override
  Widget build(context) {
    final index = useState(0);

    final loginUIList = [
      Splash(changeIndex: (int a) {
        index.value = a;
      }),
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
            child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            loginUIList[index.value],
            Padding(
                padding: const EdgeInsets.all(p),
                child: SegonButton(
                  handler: handler(context: context),
                  label: "go"
                )
            )
          ],
        )));
  }
}

void Function()? handler({
  required BuildContext context,
}) {
  return ()async{
    context.go("/hub");
  };
}
