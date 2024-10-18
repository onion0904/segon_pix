import 'package:http/http.dart' as http;

Future<http.Response> deleteUser({
  required final int id,
}) async {
  return http.delete(
      Uri.parse("http://localhost:8080/segon_pix/delete/user?ID=$id"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      });
}

Future<http.Response> deleteImage({
  required final int imageID,
  required final String token,
}) async {
  final url = Uri.http("localhost:8080", "/segon_pix_auth/delete/image",
      {"ID", imageID} as Map<String, dynamic>);
  final response = await http.delete(url, headers: {
    'Content-Type': 'application/json; charset=UTF-8',
    "Authorization": "Bearer $token"
  });
  if (response.statusCode != 200) {
    throw Exception("Failed deleteImage lib/logic/http/delete.dart");
  }
  return response;
}
