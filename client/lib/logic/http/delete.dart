import 'package:http/http.dart' as http;

Future<http.Response> deleteUser(int id) async {
   return http.delete(
    Uri.parse("http://localhost:8080/segon_pix/delete/user?ID=$id"),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    }
  );
}
