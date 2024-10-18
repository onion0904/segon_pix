import 'package:http/http.dart' as http;
import "../../model/user.dart";
import 'dart:convert';

Future<User> getUser({
  required final String userID
}) async {
  final response = await http
      .get(Uri.parse("http://localhost:8080/segon_pix/get/user?ID=$userID"));

  if (response.statusCode == 200) {
    return User.fromJson(jsonDecode(response.body) as Map<String, dynamic>);
  } else {
    throw Exception("Failed getUser method");
  }
}

Future<User> getUserWithAuth({
  required final String token,
  required final String email,
  required final String password,
}) async {
  final url = Uri.http(
    "localhost:8080",
    "/segon_pix_auth/get/user",
      {"email": email, "password": password});
  final response = await http.get(
    url,
    headers:{    "Authorization": "Bearer $token", // Bearerトークンをヘッダーに追加
    "Content-Type": "application/json", // 必要に応じて他のヘッダーも追加
    }
  );
  if (response.statusCode == 200) {
    print(jsonDecode(response.body));
    return User.fromJson(jsonDecode(response.body));
  } else {
    throw Exception("Failed getUserWithAuth method");
  }
}

Future<SimpleImages> getSimpleImages({
  required final String hashTag,
}) async {
  final url = Uri.http(
      "localhost:8080", "/segon_pix/get/list/image", {"Hashtag": hashTag});
  final response = await http.get(url);
  if (response.statusCode == 200) {
    return SimpleImages.fromJson(jsonDecode(response.body));
  } else {
    throw Exception("Failed getSimpleImages method");
  }
}

Future<PostedImage> getPostedImage({
  required final int imageID,
}) async {
  final url = Uri.http("localhost:8080", "/segon_pix/get/image",
      {"imageID", imageID} as Map<String, dynamic>);
  final response = await http.get(url);
  if (response.statusCode == 200) {
    return PostedImage.fromJson(jsonDecode(response.body));
  } else {
    throw Exception("Failed getPostedImage method");
  }
}
