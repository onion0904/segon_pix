import 'package:freezed_annotation/freezed_annotation.dart';

part 'user.freezed.dart';
part 'user.g.dart'; // JSONシリアライズ用

@freezed
class User with _$User {
  const factory User({
    required int id,
    required String name,
    String? profile,
    String? email,
  }) = _User;

  // JSONのシリアライズ/デシリアライズ用
  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}
