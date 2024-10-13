import 'package:flutter_riverpod/flutter_riverpod.dart';

final userProvider = StateProvider<User?>((ref) => null);

class User {
  final int id;
  String name;
  String icon;
  String description;
  String headerImage;
  String email;
  String password;
  String birthday;
  List<PostedImage> postedImages;
  List<PostedImage> likedImages;
  List<User> follows;
  List<User> followers;

  User({
    required this.id,
    required this.name,
    required this.icon,
    required this.description,
    required this.headerImage,
    required this.email,
    this.password = "nothing",
    required this.birthday,
    this.postedImages = const [],
    this.likedImages = const [],
    this.follows = const [],
    this.followers = const [],
  });

  Map<String, dynamic> toJson() {
    return {
      "name": name,
      "description": description,
      "email": email,
      "birthday": birthday
    };
  }

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json["id"],
      name: json["name"],
      icon: json["icon"],
      description: json["description"],
      headerImage: json["headerImage"],
      email: json["email"],
      birthday: json["birthday"],
      postedImages: json["postedImages"],
      likedImages: json["likedImages"],
      follows: json["follows"],
      followers: json["followers"]
    );
  }
}

class PostedImage {
  final int id;
  final int userId;
  final String url;
  final User user;
  final List<User> likes;
  final List<Comment> comments;

  PostedImage({
    required this.id,
    required this.userId,
    required this.url,
    required this.user,
    required this.likes,
    required this.comments,
  });

  Map<String, dynamic> toJson() {
    return {};
  }
}

class Comment {
  final int id;
  final int userId;
  String message;

  Comment({
    required this.id,
    required this.userId,
    required this.message,
  });
}

class HashTag {
  final int id;
  final String tag;
  final List<PostedImage> postedImages;

  HashTag({required this.id, required this.tag, required this.postedImages});
}
