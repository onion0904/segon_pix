import 'package:shared_preferences/shared_preferences.dart';

class UserIdManager {
  static Future<void> saveUserId(int userId) async {
    final SharedPreferences p = await SharedPreferences.getInstance();
    await p.setInt('userId', userId);
  }

  static Future<int> initializeUserId() async {
    final value = await getUserId();
    if (value != null) {
      return 1;
    } else {
      return 0;
    }
  }

  static Future<int?> getUserId() async {
    final SharedPreferences p = await SharedPreferences.getInstance();
    return p.getInt('userId');
  }

  static Future<void> updateUserId(int newUserId) async {
    await saveUserId(newUserId);
  }

  static Future<void> removeUserId() async {
    final SharedPreferences p = await SharedPreferences.getInstance();
    await p.remove('userId');
  }
}
