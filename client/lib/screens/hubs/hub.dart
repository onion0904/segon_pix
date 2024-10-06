import 'package:flutter/material.dart';
import 'package:convex_bottom_bar/convex_bottom_bar.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'home.dart';
import 'profile.dart';
import 'search.dart';
import 'notification.dart';
import 'post.dart';
import 'package:go_router/go_router.dart';

const hubUIList = [
  HubUI(),
  SearchUI(),
  Post(),
  NotificationUI(),
  ProfileUI(),
];

final floatingActionButtonList = [
  null,
  searchFloatingActionButton(),
  null,
  null,
  null,
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
      resizeToAvoidBottomInset: false,
      appBar: appBar(),
      drawer: Drawer(
        child: ListView(children: [
          const DrawerHeader(child: Center(child: Text("Menu"))),
          for (int i = 0; i < hubUIList.length; i++)
            listTile(context, i, (int a) {
              index.value = a;
              context.pop(context);
            }),
        ]),
      ),
      body: hubUIList[index.value],
      bottomNavigationBar: navigationBar((int i) => index.value = i),
      floatingActionButton: floatingActionButton(index.value)
    );
  }
}

////////////////////////////////////////////////////////

Widget floatingActionButton(
  int index
){
  return Container(
    margin: const EdgeInsets.all(16),
    child: floatingActionButtonList[index]
  );
}

/////////////////////////////////////////////////////////

PreferredSizeWidget? appBar(){
  return AppBar(
    title: const Center(
      child: Text("\$eg0n", style: TextStyle(color: Colors.white)),
    ),
    elevation: 4,
    backgroundColor: Colors.orange,
  );
}

/////////////////////////////////////////////////////////

Widget navigationBar(void Function(int) changeIndex) {
  //TODO navigationBarを変える
  return ConvexAppBar(
    style: TabStyle.fixed,
    height: 64,
    curveSize: 128,
    elevation: 4,
    backgroundColor: Colors.orange,
    items: [
      for (int i = 0; i < hubUIList.length; i++)
        TabItem(icon: hubIconList[i], title: hubNameList[i])
    ],
    onTap: changeIndex
    );
}

/////////////////////////////////////////////////////////////////////////

Widget listTile(
  BuildContext context, i, void Function(int) changeIndex
) {
  return GestureDetector(
    child: Padding(
        padding: const EdgeInsets.all(16),
        child: Row(children: [
          const SizedBox(width: 8),
          Icon(hubIconList[i], size: 32),
          const SizedBox(width: 8),
          Text(hubNameList[i], style: const TextStyle(fontSize: 16)),
        ])),
    onTap: () {
      changeIndex(i);
    });
}
