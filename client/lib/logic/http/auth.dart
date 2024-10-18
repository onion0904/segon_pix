import 'package:http/http.dart' as http;
import 'dart:convert';

Future<http.Response> signUp({
  required String email
})async{
  return http.post(
    Uri.parse("http://localhost:8080/signup"),
    headers: <String, String>{"Content-Type": "application/json; charset=UTF-8"},
    body: jsonEncode({
      "Email": email
    })
  );
}

Future<http.Response> verify({
  required String email,
  required String password,
  required String code, // 認証コード 数値
}){
  return http.post(
    Uri.parse("http://localhost:8080/verify"),
    headers: <String, String>{"Content-Type": "application/json; charset=UTF-8"},
    body: jsonEncode({
      "Email": email,
      "password": password,
      "code": code
    })
  );
}

Future<http.Response> getJWT({
  required String email,
  required String password,
}){
  return http.post(
    Uri.parse("http://localhost:8080/verify"),
    headers: <String, String>{"Content-Type": "application/json; charset=UTF-8"},
    body: jsonEncode({
      "Email": email,
      "password": password
    })
  );
}
