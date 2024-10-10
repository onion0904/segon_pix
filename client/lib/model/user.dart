class User {
  final int id;
  String name;
  String icon;
  String headerImage;
  String email;
  String birthday;

  User({
    required this.id,
    required this.name,
    required this.icon,
    required this.headerImage,
    required this.email,
    required this.birthday,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return switch (json) {
      {
        "id": int id,
        "name": String name,
        "icon": String icon,
        "headerImage": String headerImage,
        "email": String email,
        "birthday": String birthday
      } =>
        User(
            id: id,
            name: name,
            icon: icon,
            headerImage: headerImage,
            email: email,
            birthday: birthday),
      _ => throw const FormatException("Failed get uesr")
    };
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
