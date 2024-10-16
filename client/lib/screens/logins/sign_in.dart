import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import '../commons/input_form.dart';
import '../commons/button.dart';

class SignIn extends HookWidget {
  final void Function(int) changeIndex;
  const SignIn({super.key, required this.changeIndex});

  @override
  Widget build(BuildContext context) {
    return Column(children: [
      InputForm(
          controllers: controllers, validators: validators, labels: labels),
      SegonButton(
        handler: () {
          changeIndex(2);
        },
        label: "Go to Sign up"
      ),
    ]);
  }
}

//////////////////////////////////////////////////////

final validators = [emailValiator, passwordValiator];

final controllers = [
  TextEditingController(),
  TextEditingController(),
];

final labels = ["Email", "Password"];

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
