class User {
  final int id;
  String name;
  String icon;
  String profile;
  String headerImage;
  String email;
  String password;
  String birthday;

  User({
    required this.id,
    required this.name,
    required this.icon,
    required this.profile,
    required this.headerImage,
    required this.email,
    this.password = "nothing",
    required this.birthday,
  });

  Map<String, dynamic> toJson() {
    return {
      "name": name,
      "profile": profile,
      "email": email,
      "birthday": birthday
    };
  }

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json["id"],
      name: json["name"],
      icon: json["icon"],
      profile: json["profile"],
      headerImage: json["headerImage"],
      email: json["email"],
      birthday: json["birthday"],
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
