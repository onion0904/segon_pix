import 'package:shared_preferences/shared_preferences.dart';
import '../http/auth.dart';
import "dart:convert";
import '../../model/user.dart';

class UserManager {
  //保存すべきはemail password idの3つ
  static User? user;
  static int? userID;
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

    userID = p.getInt("userID");
    email = p.getString("email") as String;
    password = p.getString("password") as String;

    //TODO syuusei
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
    userID = null;
    email = "";
    password = "";
    token = "";
    user = null;
    p.remove("userID");
    p.remove("email");
    p.remove("password");
  }

  static Future<void> setMainInstance() async {
    if (userID == null || email == "" || password == "") {
      throw Exception("Failed setMainIncetance() lib/logic/db/user_manager.dart");
    }else{
      final p = await SharedPreferences.getInstance();
      await p.setString("email", email);
      await p.setString("password", password);
      await p.setInt("userID", userID as int);
    }
  }
}
