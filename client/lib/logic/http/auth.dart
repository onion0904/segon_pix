import 'package:http/http.dart' as http;
import 'dart:convert';

Future<http.Response> signUp({required String email}) async {
  final url = Uri.http("localhost:8080", "/signup", {"email": email});
  return http.post(
    url,
    headers: <String, String>{
      "Content-Type": "application/json; charset=UTF-8"
    },
  );
}

Future<http.Response> verify({
  required String email,
  required String password,
  required String code, // 認証コード 数値
}) {
  final url = Uri.http("localhost:8080", "/verify",
      {"email": email, "password": password, "code": code});

  return http.post(url, headers: <String, String>{
    "Content-Type": "application/json; charset=UTF-8"
  });
}

Future<http.Response> getJWT({
  //login
  required String email,
  required String password,
}) {
  final url = Uri.http(
    "localhost:8080",
    "/login",
    {"email": email, "password": password}
  );
  return http.post(
      url,
      headers: <String, String>{
        "Content-Type": "application/json; charset=UTF-8"
      },);
}
