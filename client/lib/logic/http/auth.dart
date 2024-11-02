import 'package:http/http.dart' as http;
import 'dart:convert';

Future<http.Response> signUp({required final String email}) async {
  final url = Uri.http("localhost:8080", "/signup", {"email": email});
  return http.post(
    url,
    headers: <String, String>{
      "Content-Type": "application/json; charset=UTF-8"
    },
  );
}

Future<http.Response> verifyAddUser({
  required final String name,
  final String description = "",
  required final String email,
  required final String password,
  required final int birthday,
  required final String code,
}) async {
  final url = Uri.http("localhost:8080", "/verify", {"code": code});

  final headers = {
    "Content-Type": "application/json",
  };

  final body = jsonEncode({
    "name": name,
    "description": description,
    "email": email,
    "password": password,
    "birthday": birthday,
  });

  // POSTリクエストを送信
  final response = await http.post(
    url,
    headers: headers,
    body: body,
  );

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed verifyAddUser lib/logic/http/auth.dart");
  }
}


Future<http.Response> login({
  required final String email,
  required final String password,
}) async {
  final url = Uri.http("localhost:8080", "/login");

  final request = http.MultipartRequest("POST", url)
    ..fields["email"] = email
    ..fields["password"] = password;

  final response = await http.Response.fromStream(await request.send());

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed login lib/logic/http/auth/dart");
  }
}
