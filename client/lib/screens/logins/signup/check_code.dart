import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../../commons/input_form.dart';
import '../../commons/button.dart';
import '../../../logic/db/provider.dart';

class CheckCode extends HookConsumerWidget {
  CheckCode({
    super.key,
    required this.changeIndex,
  });

  final controllers = [TextEditingController()];
  final validators = [emptyValidator];
  final labels = ["Code"];
  final void Function(int) changeIndex;

  @override
  Widget build(context, ref) {
    final code = ref.watch(codeProvider);

    return Center(
        child: Column(children: [
      InputForm(
        controllers: controllers,
        validators: validators,
        labels: labels,
      ),
      SegonButton(
          handler: () {
            ref.read(codeProvider.notifier).state =
                controllers[0].text;
            WidgetsBinding.instance.addPostFrameCallback((_) {
              changeIndex(3);
            });
          },
          label: "Enter"),
    ]));
  }
}

String? emptyValidator(String? value) {
  return null;
}
