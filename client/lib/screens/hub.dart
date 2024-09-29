import 'package:flutter/material.dart';
import 'package:convex_bottom_bar/convex_bottom_bar.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:client/screens/home.dart';
import 'package:client/screens/profile.dart';
import 'package:client/screens/search.dart';
import 'package:client/screens/notification.dart';

const hubList = [
  HubUI(),
  NotificationUI(),
  SearchUI(),
  ProfileUI(),
];

const hubListName = ["home", "search", "notify", "profile"];

const hubListIcon = [
  Icons.home,
  Icons.search,
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
          for (int i = 0; i < hubList.length; i++) listTile(i),
        ]),
      ),
      body: hubList[index.value],
      bottomNavigationBar: ConvexAppBar(
        style: TabStyle.react,
        backgroundColor: Colors.orange,
        items: [
          TabItem(icon: Icons.home, title: hubListName[0]),
          TabItem(icon: Icons.search, title: hubListName[1]),
          TabItem(icon: Icons.notifications, title: hubListName[2]),
          TabItem(icon: Icons.person, title: hubListName[3])
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

Widget listTile(int i) {
  return GestureDetector(
      child: Padding(
          padding: const EdgeInsets.all(8),
          child: Row(children: [
            Icon(hubListIcon[i], size: 32),
            const SizedBox(width: 8),
            Text(hubListName[i], style: const TextStyle(fontSize: 16)),
          ])),
      onTap: () {
        print("hello");
      });
}
