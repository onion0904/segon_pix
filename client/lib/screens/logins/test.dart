import 'package:flutter/material.dart';
import '../commons/sign_form.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

class Test extends StatelessWidget {
  final void Function(int) changeIndex;
  const Test({super.key, required this.changeIndex});

  @override
  Widget build(context) {
    return Column(children: [
      SignForm(),
      TextButton(
          child: const Text("TEST post"),
          onPressed: () async {
            const user = TestUser();
            await createAlbum(user);
            changeIndex(0);
          })
    ]);
  }
}

class TestUser {
  final String name = "hoge";
  final String profile = "fuga";
  final String email = "unko";
  final int age = 1;

  const TestUser();

  Map<String, dynamic> toJson() =>
      {"name": name, "profile": profile, "email": email, "age": age};
}

Future<http.Response> createAlbum(TestUser user) {
  return http.post(
    Uri.parse('http://localhost:8080/segon_pix/add/user'),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(user.toJson()),
  );
}
