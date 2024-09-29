import 'package:flutter/material.dart';
import 'package:convex_bottom_bar/convex_bottom_bar.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:client/screens/hubs/home.dart';
import 'package:client/screens/hubs/profile.dart';
import 'package:client/screens/hubs/search.dart';
import 'package:client/screens/hubs/notification.dart';
import 'post.dart';

const hubUIList = [
  HubUI(),
  SearchUI(),
  Post(),
  NotificationUI(),
  ProfileUI(),
];

const hubNameList = ["home", "search", "post", "notify", "profile"];

const hubIconList = [
  Icons.home,
  Icons.search,
  Icons.image,
  Icons.notifications,
  Icons.person
];

/////////////////////////////////////////////////////////////////////////

class Hub extends HookWidget {
  const Hub({super.key});

  @override
  Widget build(context) {
    final index = useState(0);

    return Scaffold(
      appBar: AppBar(
        title: const Center(
          child: Text("segon_pix", style: TextStyle(color: Colors.white)),
        ),
        backgroundColor: Colors.orange,
      ),
      drawer: Drawer(
        child: ListView(children: [
          const DrawerHeader(child: Center(child: Text("Menu"))),
          for (int i = 0; i < hubUIList.length; i++)
            listTile(i, () {
              index.value = i;
            }),
        ]),
      ),
      body: hubUIList[index.value],
      bottomNavigationBar: ConvexAppBar(
        style: TabStyle.fixed,
        backgroundColor: Colors.orange,
        items: [
          for (int i = 0; i < hubUIList.length; i++)
            TabItem(icon: hubIconList[i], title: hubNameList[i])
        ],
        initialActiveIndex: index.value,
        onTap: (int i) {
          index.value = i;
        },
      ),
    );
  }
}

/////////////////////////////////////////////////////////////////////////

Widget listTile(int i, void Function() operateIndex) {
  return GestureDetector(
      child: Padding(
          padding: const EdgeInsets.all(8),
          child: Row(children: [
            Icon(hubIconList[i], size: 32),
            const SizedBox(width: 8),
            Text(hubNameList[i], style: const TextStyle(fontSize: 16)),
          ])),
      onTap: () => operateIndex());
}
