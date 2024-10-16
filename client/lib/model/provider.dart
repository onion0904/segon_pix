import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../model/user.dart';

final tokenProvider = StateProvider((ref) => "");
final userProvider = StateProvider<User?>((ref) => null);
