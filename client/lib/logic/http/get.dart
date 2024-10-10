import 'package:http/http.dart' as http;
import "../../model/user.dart";
import 'dart:convert';

Future<User> getUser(int userId) async {
  final response = await http
      .get(Uri.parse("http://localhost:8080/segon_pix/get/user?ID=$userId"));

  if (response.statusCode == 200) {
    return User.fromJson(jsonDecode(response.body) as Map<String, dynamic>);
  } else {
    throw Exception("Failed getUser method");
  }
}
