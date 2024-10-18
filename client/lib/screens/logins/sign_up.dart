import 'package:flutter/material.dart';
import '../commons/input_form.dart';
import '../commons/button.dart';
import '../../logic/http/auth.dart' as auth;
import '../../logic/db/user_manager.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

class SignUp extends HookWidget {
  const SignUp({super.key, required this.changeIndex});
  final void Function(int) changeIndex;

  @override
  Widget build(context) {
final controllers = useState([
  TextEditingController(),
  TextEditingController(),
]);

    return Column(children: [
      InputForm(
          controllers: controllers.value, validators: validators, labels: labels),
      SegonButton(
          handler: () async {
            auth.signUp(email: UserManager.email);
            changeIndex(4);
          },
          label: "create"),
      SegonButton(
          handler: () async {
            UserManager.email = controllers.value[0].text;
            UserManager.password = controllers.value[1].text;
            changeIndex(1);
          },
          label: "Go to Sign in")
    ]);
  }
}

//////////////////////////////////////////////////////

final validators = [emailValiator, passwordValiator];

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
