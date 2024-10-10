import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import "../../logic/http/post.dart";
import 'package:go_router/go_router.dart';

const double p = 2;

class CreateUser extends HookWidget {
  const CreateUser({super.key});

  @override
  Widget build(BuildContext context) {
    final birthday = useState(0);
    final formKey = GlobalKey<FormState>();
    final nameController = TextEditingController();
    final descriptionController = TextEditingController();

    return Column(children: [
      Form(
          key: formKey,
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            Padding(
                padding: const EdgeInsets.all(p),
                child: TextFormField(
                    //name
                    controller: nameController,
                    validator: nameValidator)),
            Padding(
                padding: const EdgeInsets.all(p),
                child: TextFormField(
                    //description
                    controller: descriptionController,
                    validator: descriptionValidator)),
            Text("${birthday.value}"),
            Padding(
                padding: const EdgeInsets.all(p),
                child: OutlinedButton(
                    onPressed: () async {
                      final date = await showDatePicker(
                          context: context,
                          initialDate: DateTime.now(),
                          lastDate: DateTime.now(),
                          firstDate: DateTime(1900, 1, 1));
                      if(date != null){
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
                            name: nameController.text,
                            description: descriptionController.text,
                            birthday: birthday.value);
                        if (response.statusCode == 200 && context.mounted) {
                          context.go("/hub");
                        }
                      }
                    },
                    child: const Text("Deside"))),
          ]))
    ]);
  }
}

/////////////////////////////////////////////////////////////////////////

String? nameValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "名前を入力してください";
  }

  return null;
}

String? descriptionValidator(String? value) {
  return null;
}

String? birthdayValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "誕生日を入力してください";
  }

  return null;
}
