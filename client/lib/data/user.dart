class User {
  final String id;
  String name;
  int age;
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
    required this.age,
    required this.iconURL,
    required this.headerURL,
    required this.profileMessage,
    //任意
    this.postImageURLs = const <String>[],
    this.likeImageURLs = const <String>[],
    this.followingList = const <String>[],
  });
}


//疑問
//生年月日ではなく年齢　年齢確認はしておいたほうが後々楽だと思う(課金要素や表示する画像によって制限をかけれる)
