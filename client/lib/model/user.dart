class User {
  final String id;
  String name;
  int birthday;
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
    required this.birthday,
    //任意(デフォルト等で対応)
    this.iconURL        = "",
    this.headerURL      = "",
    this.profileMessage = "",
    this.postImageURLs  = const <String>[],
    this.likeImageURLs  = const <String>[],
    this.followingList  = const <String>[],
  });
}
