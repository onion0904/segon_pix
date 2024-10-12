import 'package:flutter/material.dart';
import '../commons/input_form.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

class Post extends HookWidget {
  const Post({super.key});

  @override
  Widget build(context) {
      final image = useState<XFile?>(null);

    return Center(
      child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
        if (image.value != null) Container(constraints: const BoxConstraints(maxWidth: 1000, maxHeight:1000), child: Image.file(File(image.value!.path))),
        InputForm(
          controllers: controllers, validators: validators, labels: labels),
        OutlinedButton(
          child: const Text("image"),
          onPressed: () async {
            final ImagePicker picker = ImagePicker();
            image.value = await picker.pickImage(source: ImageSource.gallery);
          })
    ]));
  }
}

final controllers = [TextEditingController()];

const labels = ["Hash Tag"];

const validators = [hashTagValidator];

String? hashTagValidator(String? value) {
  return null;
}
