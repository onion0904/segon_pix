import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'splash.dart';
import './signin/sign_in.dart';
import './signup/sign_up.dart';
import './signup/check_code.dart';
import './signup/create_user.dart';
import 'package:go_router/go_router.dart';
import '../commons/button.dart';

const double p = 2;

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
      CreateUser(),
      CheckCode(
        changeIndex: (int a) {
          index.value = a;
        }
      )
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
                    handler: handler(context: context), label: "go"))
          ],
        )));
  }
}

void Function()? handler({
  required BuildContext context,
}) {
  return () async {
    context.go("/hub");
  };
}
