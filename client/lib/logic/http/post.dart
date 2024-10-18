import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
import 'package:http_parser/http_parser.dart';
import 'package:mime/mime.dart';
import '../../model/user.dart';

Future<http.Response> createUser({
  required String name,
  required String description,
  required String email,
  required String password,
  required int birthday,
  required String token,
}) async {
  final response =
      await http.post(Uri.parse("http://localhost:8080/segon_pix_auth/add/user"),
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
  required List<String> hashTags,
  required int userId,
  required XFile imageFile,
}) async {
  var uri = Uri.http(
    "localhost:8080",
    "/segon_pix/add/image",
    {"ID": "$userId"},
  );
  var request = http.MultipartRequest('POST', uri);
  var mimeTypeData =
      lookupMimeType(imageFile.path, headerBytes: [0xFF, 0xD8])?.split('/');
  request.files.add(
    await http.MultipartFile.fromPath(
      'File',
      imageFile.path,
      contentType: MediaType(mimeTypeData![0], mimeTypeData[1]),
    ),
  );
  for (var tag in hashTags) {
    request.fields['Hashtags'] = tag;
  }
  return await request.send();
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
