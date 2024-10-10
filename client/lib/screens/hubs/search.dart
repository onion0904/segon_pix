import 'package:flutter/material.dart';

const screen = [

];

class SearchUI extends StatelessWidget {
  const SearchUI({super.key});

  @override
  Widget build(context) {
    return const Center(child: Text("search"));
  }
}

////////////////////////////////////////////////////////

Widget searchFloatingActionButton() {
  return FloatingActionButton(
      onPressed: () {
        //TODO 画面移動
      },
      child: const Icon(Icons.search));
}

class SearchFieldUI extends StatelessWidget {
  SearchFieldUI({super.key});

  final textController = TextEditingController();

  @override
  Widget build(context) {
    return Row(children: [
      TextFormField(
        decoration: const InputDecoration(
          label: Text("検索"),
        ),
        controller: textController,
        validator: searchValidator,
      )
    ]);
  }
}

String? searchValidator(String? value) {
  if (value == null || value.isEmpty) {
    return "入力してください";
  }
  return null;
}
