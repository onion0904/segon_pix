import 'package:http/http.dart' as http;
import 'package:image_picker/image_picker.dart';

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

Future<http.Response> updateComment({
  required final int userID,
  required final int commentID,
  required final int imageID,
  required final String newComment,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/update/comment", {
    "userID": userID,
    "commentID": commentID,
    "imageID": imageID,
    "newComment": newComment
  });

  final response =
      await http.put(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed updateContent /lib/logic/http/put.dart 38");
  }
}

Future<http.Response> updateUserHeader({
  //header de error deru kamo
  required final int userID,
  required final XFile file,
  required final String token,
}) async {
  final url =
      Uri.http(host, "/segon_pix_auth/update/user/header", {"userID": userID});

  final request = http.MultipartRequest("PUT", url)
    ..headers["Authorization"] = "Bearer $token"
    ..files.add(await http.MultipartFile.fromPath("image", file.path));

  final response = await http.Response.fromStream(await request.send());

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed updateUserHeader /lib/logic/http/put.dart f61");
  }
}

Future<http.Response> updateUserIcon({
  //header de error deru kamo
  required final int userID,
  required final XFile file,
  required final String token,
}) async {
  final url =
      Uri.http(host, "/segon_pix_auth/update/user/icon", {"userID": userID});

  final request = http.MultipartRequest("PUT", url)
    ..headers["Authorization"] = "Bearer $token"
    ..files.add(await http.MultipartFile.fromPath("image", file.path));

  final response = await http.Response.fromStream(await request.send());

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed updateUserIcon /lib/logic/http/put.dart 8f3");
  }
}

Future<http.Response> updateUser({
  required final int userID,
  required final String name,
  required final String description,
  required final int birthday,
  required final String email,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix/update/user", {
    "userID": userID,
    "name": name,
    "description": description,
    "birthday": birthday,
    "email": email
  }); // TODO maybe

  final response =
      await http.put(url, headers: returnHeaderWithToken(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed updateUser lib/logic/http/put.dart 109");
  }
}
