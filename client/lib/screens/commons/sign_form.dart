import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:go_router/go_router.dart';

const p = 64.0;
const double n = p;

class SignForm extends HookWidget {
  SignForm({super.key});

  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  Widget build(context) {
    //TODO パスワードのハッシュ化
    return Form(
        child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
      Padding(
          padding: const EdgeInsets.fromLTRB(n, p, n, p),
          child: TextFormField(
            controller: emailController,
            validator: emailValiator,
          )),
      Padding(
          padding: const EdgeInsets.fromLTRB(n, p, n, p),
          child: TextFormField(
            controller: passwordController,
            validator: passwordValiator,
          )),
      Padding(
          padding: const EdgeInsets.all(p),
          child: ElevatedButton(
              onPressed: () {
                //TODO サーバに送信

                context.go("/hub");
              },
              child: const Text("決定")))
    ]));
  }
}

//////////////////////////////////////////////////////

String? emailValiator(String? value) {
  //TODO 正規表現によるメアドの形式チェック
  if (value == null || value == "") {
    return "メールアドレスを入力してください";
  } else if (!value.contains("@")) {
    return "有効な形式ではありません";
  }
  return "";
}

String? passwordValiator(String? value) {
  //TODO 正規表現による有効な文字チェック
  if (value == null || value == "") {
    return "パスワードを入力してください";
  } else if (8 <= value.length && value.length <= 32) {
    return "8文字以上32文字以下にしてください";
  }
  return "";
}
