import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import '../../logic/http/auth.dart';
import '../commons/button.dart';
import '../commons/input_form.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import '../../logic/db/user_manager.dart';
import 'dart:convert';
import '../../logic/http/get.dart';
import 'package:go_router/go_router.dart';

const double p = 2;

class CheckCode extends HookConsumerWidget {
  const CheckCode({
    super.key,
    required this.changeIndex,
  });

  final void Function(int) changeIndex;

  @override
  Widget build(context, ref) {
    final formKey = GlobalKey<FormState>();
    final controllers = useState([
      TextEditingController(),
    ]);

    return Column(children: [
      Form(
          key: formKey,
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            Text("${UserManager.email}"),
            InputForm(
                validators: validators,
                controllers: controllers.value,
                labels: labels),
            SegonButton(
                handler:
                  handler(
                      code: controllers.value[0].text,
                      context: context,
                      changeIndex: changeIndex),

                label: "Enter")
          ]))
    ]);
  }
}

/////////////////////////////////////////////////////////////////////////

void Function() handler({
  required final String code,
  required final BuildContext context,
  required final void Function(int) changeIndex,
}) {
  //userがいたら、ホームに飛ばす
  return () async {
    try {

      final response = await verify(
        email: UserManager.email,
        password: UserManager.password,
        code: code,
      );
      print("1");

      if (response.statusCode == 200) {
        final json = jsonDecode(response.body);
        print("json $json");
        UserManager.token = json["token"];
      }

      print("2");

      UserManager.user = await getUserWithAuth(
          token: UserManager.token,
          email: UserManager.email,
          password: UserManager.password);

      print("4");
      if (context.mounted) {
        context.go("/hub");
      }
    } catch (e) {
      //todo printをデバック用のやつに変える
      print(e);
    }

    //まだuser登録していない場合
    changeIndex(3);
  };
}

const validators = [codeValidator];

const labels = ["Authentication Code"];

String? codeValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "入力してください";
  }

  return null;
}
