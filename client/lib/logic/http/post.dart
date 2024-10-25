import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
import '../../model/user.dart';

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

// Future<http.Response> createUser({
//   required String name,
//   required String description,
//   required String email,
//   required String password,
//   required int birthday,
//   required String token,
// }) async {
//   final response = await http.post(
//       Uri.parse("http://localhost:8080/segon_pix_auth/add/user"),
//       headers: {
//         "Content-Type": "application/json; charset=UTF-8",
//         "Authorization": "Bearer $token"
//       },
//       body: jsonEncode({
//         "Name": name,
//         "Description": description,
//         "Email": email,
//         "Password": password,
//         "birthday": birthday,
//       }));
//   if (response.statusCode == 200) {
//     return response;
//   } else {
//     throw Exception("Failed createUser method lib/logic/http/post.dart 31");
//   }
// }

Future<http.Response> addImage({
  required final int userID,
  required final XFile file,
  required final String hashTag,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/add/image", {"userID": userID});

  final request = http.MultipartRequest("POST", url)
    ..headers["Authorization"] = "Bearer $token"
    ..fields["HashTag"] = hashTag
    ..files.add(await http.MultipartFile.fromPath("image", file.path));

  final response = await http.Response.fromStream(await request.send());
  if (response.statusCode == 200) {
    // final data = await response.stream.bytesToString();
    // return PostedImage.fromJson(jsonDecode(data));
    return response;
  } else {
    throw Exception("Failed addImage lib/logic/http/post.dart 65");
  }
}

Future<http.Response> addLike({
  required final int userID,
  required final int imageID,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/add/like");

  final response =
      await http.post(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed addLike lib/logic/http/post.dart 85");
  }
}

Future<http.Response> addComment({
  required final int userID,
  required final int imageID,
  required final String comment,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/add/comment",
      {"userID": userID, "imageID": imageID, "comment": comment});

  final response =
      await http.post(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed addComment /lib/logic/http/post.dart 104");
  }
}

Future<http.Response> addFollow({
  required final int followingID,
  required final int followedID,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/add/follow");

  final response =
      await http.post(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed addFollow /lib/logic/http/post.dart 121");
  }
}
