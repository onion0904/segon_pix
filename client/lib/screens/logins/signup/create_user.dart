import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'dart:convert';
import 'package:flutter_hooks/flutter_hooks.dart';

import '../../commons/input_form.dart';
import '../../commons/button.dart';
import '../../../logic/db/provider.dart';
import '../../../logic/http/auth.dart';
import '../../../logic/db/user_manager.dart';

class CreateUser extends HookConsumerWidget {
  CreateUser({super.key});

  final controllers = [TextEditingController(), TextEditingController()];
  final validators = [emptyValidator, emptyValidator];
  final labels = ["Name", "Description"];

  int convertDateToInt(DateTime date) {
    return int.parse("${date.year.toString().padLeft(4, '0')}"
        "${date.month.toString().padLeft(2, '0')}"
        "${date.day.toString().padLeft(2, '0')}");
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final date = useState(DateTime.now());
    final dateInt = useState<int>(convertDateToInt(date.value));

    return Center(
        child: Column(children: [
      InputForm(
        controllers: controllers,
        validators: validators,
        labels: labels,
      ),
      SegonButton(
          handler: () async {
            final selectedDate = await showDatePicker(
              context: context,
              initialDate: date.value,
              firstDate: DateTime(1900),
              lastDate: DateTime.now(),
            );
            if (selectedDate != null) {
              date.value = selectedDate;
              dateInt.value = convertDateToInt(selectedDate);
            }
          },
          label: "Birthday"),
      SegonButton(
          handler: () async {
            try {
              final response = await verifyAddUser(
                name: controllers[0].text,
                description: controllers[1].text,
                email: UserManager.email,
                password: UserManager.password,
                birthday: dateInt.value,
                code: ref.watch(codeProvider.notifier).state.toString(),
              );
              final token = jsonDecode(response.body);
              print(token);
              print("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            } catch (e) {
              print(e);
            }
          },
          label: "Enter"),
    ]));
  }
}

String? emptyValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "This field can't be empty";
  }
  return null;
}
