import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
import 'package:http_parser/http_parser.dart';
import 'package:mime/mime.dart';

Future<http.Response> createUser(
    {required String name,
    required String description,
    required int birthday}) {
  return http.post(Uri.parse("http://localhost:8080/segon_pix/add/user"),
      headers: <String, String>{
        "Content-Type": "application/json; charset=UTF-8"
      },
      body: jsonEncode(
          {"name": name, "description": description, "birthday": birthday}));
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

  var mimeTypeData = lookupMimeType(imageFile.path, headerBytes: [0xFF, 0xD8])?.split('/');
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
