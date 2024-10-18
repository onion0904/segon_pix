import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
// import 'package:http_parser/http_parser.dart';
// import 'package:mime/mime.dart';

Future<http.Response> createUser({
  required String name,
  required String description,
  required String email,
  required String password,
  required int birthday,
  required String token,
}) async {
  final response = await http.post(
      Uri.parse("http://localhost:8080/segon_pix_auth/add/user"),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Authorization": "Bearer $token"
      },
      body: jsonEncode({
        "Name": name,
        "Description": description,
        "Email": email,
        "Password": password,
        "birthday": birthday,
      }));
  if (response.statusCode == 200) {
    return response;
  } else {
    throw Exception("Failed createUser method lib/logic/http/post.dart 31");
  }
}

Future<http.StreamedResponse> postImage({
  required String hashTags,
  required int userId,
  required XFile imageFile,
  required String token,
}) async {
  final uri = Uri.http(
    "localhost:8080",
    "/segon_pix/add/image",
    {"ID": "$userId"},
  );

  var request = http.MultipartRequest(
    'POST',
    uri,
  )
    ..headers["Authorization"] = "Bearer $token"
    ..files.add(await http.MultipartFile.fromPath("File", imageFile.path))
    ..fields["Hashtags"] = hashTags;


    final response =  await request.send();
    if(response.statusCode == 200){
      return response;
    }else {
    throw Exception("Failed postImage lib/logic/http/post.dart");
  }
}

Future<http.Response> addLike({
  required int userId,
  required int imageId,
  required String token,
}) {
  return http.post(
    Uri.parse(
        "http://localhost:8080/segon_pix_auth/add/like?userID=$userId&imageID=$imageId"),
    headers: {
      "Content-Type": "application/json; charset=UTF-8",
      "Authorization": "Bearer $token"
    },
  );
}
