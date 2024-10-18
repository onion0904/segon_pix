import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import '../commons/button.dart';
import '../commons/input_form.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import "../../logic/http/post.dart";
import 'package:go_router/go_router.dart';
import '../../logic/db/user_manager.dart';
import '../../logic/http/get.dart';

const double p = 2;

class CreateUser extends HookConsumerWidget {
  //TODO addUserでUserを作成する

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

    //気が向いたらきれいにする

    return Column(children: [
      Form(
          key: formKey,
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            InputForm(
                validators: validators,
                controllers: controllers.value,
                labels: labels),
            Padding(
                padding: const EdgeInsets.all(p),
                child: SegonButton(
                    handler: () async {
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
                    label: "Select birthday")),
            Padding(
                padding: const EdgeInsets.all(p),
                child: SegonButton(
                    handler: () async {
                      if (formKey.currentState!.validate()) {
                        try {
                          await createUser(//addUser
                            name: controllers.value[0].text,
                            description: controllers.value[1].text,
                            email: UserManager.email,
                            password: UserManager.password,
                            birthday: birthday.value,
                            token: UserManager.token,
                          );
                          UserManager.user = await getUserWithAuth(
                              token: UserManager.token,
                              email: UserManager.email,
                              password: UserManager.password
                          );

                          //db(偽)に保存

                          await UserManager.setMainInstance();
                          if (context.mounted) {
                            context.go("/hub");
                          }
                        } catch (e) {
                          print(e);
                        }
                      }
                    },
                    label: "create")),
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
