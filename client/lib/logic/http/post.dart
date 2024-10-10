import 'package:http/http.dart' as http;
import 'dart:convert';

Future<http.Response> createUser({
  required String name, required String description, required int birthday
}) {
  return http.post(
    Uri.parse("http://localhost:8080/segon_pix/add/user"),
    headers: <String, String>{"Content-Type": "application/json; charset=UTF-8"},
    body:jsonEncode({
      "name": name,
      "description": description,
      "birthday": birthday
    })
  );
}
