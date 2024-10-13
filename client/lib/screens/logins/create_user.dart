import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import "../../logic/http/post.dart";
import 'package:go_router/go_router.dart';
import '../commons/input_form.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import '../../model/user.dart';
import 'dart:convert';

const double p = 2;

class CreateUser extends HookConsumerWidget {
  const CreateUser({super.key});

  @override
  Widget build(context, ref) {
    final birthday = useState(0);
    final initialDate = useState(DateTime.now());
    final formKey = GlobalKey<FormState>();
    final controllers = useState([
      TextEditingController(),
      TextEditingController(),
    ]);

    return Column(children: [
      Form(
          key: formKey,
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            InputForm(
                validators: validators,
                controllers: controllers.value,
                labels: labels),
            Text("${initialDate.value}"),
            Padding(
                padding: const EdgeInsets.all(p),
                child: OutlinedButton(
                    onPressed: () async {
                      final date = await showDatePicker(
                          context: context,
                          initialDate: initialDate.value,
                          lastDate: DateTime.now(),
                          firstDate: DateTime(1900, 1, 1));
                      if (date != null) {
                        initialDate.value = date;
                        birthday.value =
                            date.year * 10000 + date.month * 100 + date.day;
                      }
                    },
                    child: const Text("Select Birthday"))),
            Padding(
                padding: const EdgeInsets.all(p),
                child: OutlinedButton(
                    onPressed: () async {
                      if (formKey.currentState!.validate()) {
                        final response = await createUser(
                            name: controllers.value[0].value.text,
                            description: controllers.value[1].value.text,
                            birthday: birthday.value);
                        if (response.statusCode == 200 && context.mounted) {
                          //ここのjsonがどんなかんじになっているか確認
                          ref.read(userProvider.notifier).state = User.fromJson(jsonDecode(response.body));
                          context.go("/hub");
                        }
                      }
                    },
                    child: const Text("Decide"))),
          ]))
    ]);
  }
}

/////////////////////////////////////////////////////////////////////////

const validators = [nameValidator, descriptionValidator];

const labels = ["Name", "Description"];

String? nameValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "名前を入力してください";
  }

  return null;
}

String? descriptionValidator(String? value) {
  return null;
}
