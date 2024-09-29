class User {
  final String id;
  String name;
  String iconURL;
  String headerURL;
  String profileMessage;
  List<String> postImageURLs;
  List<String> likeImageURLs;
  List<String> followingList;

  User({
    //必須
    required this.id,
    required this.name,
    required this.iconURL,
    required this.headerURL,
    required this.profileMessage,
    //任意
    this.postImageURLs = const <String>[],
    this.likeImageURLs = const <String>[],
    this.followingList = const <String>[],
  });

}
