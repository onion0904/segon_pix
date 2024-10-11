import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:go_router/go_router.dart';

const p = 32.0;
const double n = p;

class InputForm extends HookWidget {
  const InputForm({
    super.key,
    required this.controllers,
    required this.validators,
    required this.labels,
  });

  final List<TextEditingController> controllers;
  final List<String? Function(String?)?> validators;
  final List<String> labels;

  @override
  Widget build(context) {
    return Form(
        child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
      for (int i = 0; i < controllers.length; i++)
        Container(
          constraints: const BoxConstraints(maxWidth: 1000),
          padding: const EdgeInsets.fromLTRB(n, p, n, p),
          child: TextFormField(
            controller: controllers[i],
            validator: validators[i],
            decoration: InputDecoration(
                border: const OutlineInputBorder(), labelText: labels[i]),
          ),
        ),
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
