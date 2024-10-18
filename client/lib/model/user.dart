class User {
  final int id;
  String name;
  String icon;
  String description;
  String headerImage;
  String email;
  String password;// 多分これはいらない
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
        followers: json["followers"]);
  }
}

class PostedImage {
  final int id;
  final int userID;
  final String url;
  final User user;
  final List<User> likes;
  final List<Comment> comments;
  final String hashTag;

  PostedImage({
    required this.id,
    required this.userID,
    required this.url,
    required this.user,
    required this.likes,
    required this.comments,
    required this.hashTag,
  });

  Map<String, dynamic> toJson() {
    return {
      "ID": id,
      "UserID": userID,
      "URL": url,
      "PostUser": user,
      "Likes": likes,
      "Comments": comments,
      "Hashtags": hashTag,
    };
  }

  factory PostedImage.fromJson(Map<String, dynamic> json) {
    return PostedImage(
        id: json["ID"],
        userID: json["UserID"],
        url: json["URL"],
        user: json["PostUser"],
        likes: json["Likes"],
        comments: json["Comments"],
        hashTag: json["Hashtags"],
    );
  }
}

class Comment {
  final int id;
  final int userID;
  String message;

  Comment({
    required this.id,
    required this.userID,
    required this.message,
  });
}

class HashTag {
  final int id;
  final String tag;
  final List<PostedImage> postedImages;

  HashTag({required this.id, required this.tag, required this.postedImages});
}

class SimpleImages {
  final List<int> imageIDs;
  final List<String> imageURLs;

  SimpleImages({
    required this.imageIDs,
    required this.imageURLs,
  });

  factory SimpleImages.fromJson(Map<String, dynamic> json) {
    return SimpleImages(
      imageIDs: json["ID"],
      imageURLs: json["URL"],
    );
  }
}
