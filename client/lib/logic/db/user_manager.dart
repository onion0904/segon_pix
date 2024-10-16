import 'package:shared_preferences/shared_preferences.dart';
import '../../logic/http/auth.dart';
import "dart:convert";

class UserManager {
  static int userID = 0;
  static String email = "";
  static String password = "";
  static String token = "";

  static Future<int?> initializeUserManager() async {
    final p = await SharedPreferences.getInstance();
    if (p.getString("email") == null ||
        p.getString("password") == null ||
        p.getInt("userID") == null) {
      return -1;
    }

    userID = p.getInt("userID") as int;
    email = p.getString("email") as String;
    password = p.getString("password") as String;

    final response = await getJWT(email: email, password: password);
    if (response.statusCode != 200) {
      return -1;
    }

    final json = jsonDecode(response.body);
    token = json["token"];
    return 1;
  }

  static Future<void> resetUserManager() async {
    final p = await SharedPreferences.getInstance();
    userID = 0;
    email = "";
    password = "";
    token = "";
    p.remove("userID");
    p.remove("email");
    p.remove("password");
  }

  static Future<void> saveEmailAndPassword(
      {required int userID,
      required String email,
      required String password,
      required String token}) async {
    final p = await SharedPreferences.getInstance();
    await p.setString("email", email);
    await p.setString("password", password);
  }
}
