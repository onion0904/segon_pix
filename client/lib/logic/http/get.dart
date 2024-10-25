import 'package:http/http.dart' as http;
import "../../model/user.dart";
import '../db/user_manager.dart';
import 'dart:convert';

const String host = "localhost:8080";

Map<String, String> returnHeaderWithToken({
  required final String token,
}) {
  return {
    "Content-Type": "application/json; charset=UTF-8",
    "Authorization": "Bearer $token"
  };
}

Map<String, String> returnHeader() {
  return {"Content-Type": "application/json; charset=UTF-8"};
}

///

Future<User> getUser({required final int userID}) async {
  final url = Uri.http(host, "/segon_pix/get/user", {"userID": userID});

  final response = await http.get(url, headers: returnHeader());

  if (response.statusCode == 200) {
    final json = await jsonDecode(response.body);
    return User(
        id: json["id"],
        name: json["name"],
        icon: json["icon"],
        description: json["description"],
        headerImage: json["headerImage"],
        birthday: json["birthday"],
        postedImages: json["postedImage"],
        likedImages: json["likedImages"]);
  } else {
    throw Exception("Failed return");
  }
}

Future<http.Response> getUserWithToken({
  //TODO クエリパラで良いのか確認
  required final int userID,
  required final String email,
  required final String password,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/get/user",
      {"userID": userID, "email": email, "password": password});

  final response =
      await http.get(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    final json = await jsonDecode(response.body);
    UserManager.user = User(
        id: json["id"],
        name: json["name"],
        icon: json["icon"],
        description: json["description"],
        headerImage: json["headerImage"],
        email: UserManager.email,
        password: UserManager.password,
        birthday: json["birthday"],
        postedImages: json["postedImage"],
        likedImages: json["likedImages"]);
    return response;
  } else {
    throw Exception("Failed getUserWithToken lib/logic/http/get.dart 72");
  }
}

Future<List<SimpleImage>> getListSearch({required final String hashTag}) async {
  final url =
      Uri.http(host, "/segon_pix/get/list/search", {"Hashtag": hashTag});

  final response = await http.get(
    url,
    headers: returnHeader(),
  );

  if (response.statusCode == 200) {
    final json = await jsonDecode(response.body);
    final simpleImageList = json.map((item) {
      return SimpleImage.fromJson(item);
    });
    return simpleImageList;
  } else {
    throw Exception("Failed getListSearch lib/logic/http/get 92");
  }
}

Future<List<SimpleImage>> getListLike() async {
  final url = Uri.http(host, "/segon_pix/get/list/like");

  final response = await http.get(url, headers: returnHeader());

  if (response.statusCode == 200) {
    final json = await jsonDecode(response.body);
    return json.map((item) {
      return SimpleImage.fromJson(item);
    });
  } else {
    throw Exception("Failed getListLike lib/logic/http/get.dart 107");
  }
}

Future<List<SimpleImage>> gerImageRecent() async {
  final url = Uri.http(host, "/segon_pix/get/list/recent");

  final response = await http.get(url, headers: returnHeader());

  if (response.statusCode == 200) {
    final json = await jsonDecode(response.body);
    return json.map((item) {
      return SimpleImage.fromJson(item);
    });
  } else {
    throw Exception("Failed getListRecent lib/logic/http/get.dart 122");
  }
}

Future<PostedImage> gerImageDetail({
  required final int imageID,
}) async {
  final url =
      Uri.http(host, "/segon_pix/get/image_detail", {"imageID": imageID});

  final response = await http.get(url, headers: returnHeader());

  if (response.statusCode == 200) {
    return PostedImage.fromJson(jsonDecode(response.body));
  } else {
    throw Exception("Failed getImageDetail lib/logic/http/get.dart 137");
  }
}
