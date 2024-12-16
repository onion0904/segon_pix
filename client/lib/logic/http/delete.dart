import 'package:http/http.dart' as http;

//aaa/bbb -> aaaBbb

const String host = "localhost:8080";

Map<String, String> returnHeaders({
  required final String token,
}) {
  return {
    "Content-Type": "application/json; charset=UTF-8",
    "Authorization": "Bearer $token"
  };
}

Future<http.Response> deleteUser(
    {required final int userID, required final String token}) async {
  final url = Uri.http(host, "/segon_pix_auth/delete/user", {"userID": userID});

  final response = await http.delete(url, headers: returnHeaders(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed deleteUser lib/logic/http/delete/dart 15line");
  }
}

Future<http.Response> deleteImage({
  required final int imageID,
  required final int userID,
  required final String token,
}) async {
  final url = Uri.http(host, "/segon_pix_auth/delete/image",
      {"imageID": imageID, "userID": userID});

  final response = await http.delete(url, headers: returnHeaders(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed deleteImage lib/logic/http/delete.dart 39");
  }
}

Future<http.Response> deleteLike({
  required final int userID,
  required final int imageID,
  required final String token,
}) async {
  final url =
      Uri.http(host, "/delete/like", {"imageID": imageID, "userID": userID});

  final response = await http.delete(url, headers: returnHeaders(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed deleteLike lib/logic/http/delete 58");
  }
}

Future<http.Response> deleteComment({
  required final int userID,
  required final int commentID,
  required final String token,
}) async {
  final url = Uri.http(
      host, "/delete/comment", {"userID": userID, "commentID": commentID});

  final response = await http.delete(url, headers: returnHeaders(token: token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed deleteComment lib/logic/http/delete 76");
  }
}

Future<http.Response> deleteFollow({
  required final int followingID,
  required final int followedID,
  required final String token,
})async{
  final url = Uri.http(
    host,
    "/delete/follow",
    {
      "followingID": followingID,
      "followedID": followedID
    }
  );

  final response = await http.delete(url, headers: returnHeaders(token:token));

  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed deletefollow lib/logic/http/delete 99");
  }
}
