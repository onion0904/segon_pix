import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'home.dart';
import 'profile.dart';
import 'search.dart';
import 'notification.dart';
import 'post.dart';
import 'package:go_router/go_router.dart';

/////////////////////////////////////////////////////////////////////////
const double n = 1;
const double p = 8*1.5;
const double sWidth = 32;
const drawerImage =
    "https://onion0904.dev/ocGvg5tH5gfqsDS1715839141_1715839204.png";
const double imageSize = 100;

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
            DrawerHeader(
                decoration: const BoxDecoration(color: Colors.orange),
                child: Center(
                    child: SizedBox(
                        width: imageSize,
                        height: imageSize,
                        child: Image.network(drawerImage)))),
            for (int i = 0; i < HubList.ui.length; i++)
              listTile(context, i, (int a) {
                index.value = a;
                context.pop(context);
              }),
          ]),
        ),
        body: Center(child: HubList.ui[index.value]),
        bottomNavigationBar: SegonNavigationBar(
            index: index.value, changeIndex: (int i) => index.value = i),
        floatingActionButton: floatingActionButton(index.value));
  }
}

////////////////////////////////////////////////////////

Widget floatingActionButton(int index) {
  return Container(
      margin: const EdgeInsets.all(16),
      child: HubList.floatingActionButton[index]);
}

/////////////////////////////////////////////////////////

PreferredSizeWidget? appBar() {
  return AppBar(
    title: const Center(
      child: Text("\$eg0n", style: TextStyle(color: Colors.white)),
    ),
    elevation: 4,
    backgroundColor: Colors.orange,
  );
}

/////////////////////////////////////////////////////////

class SegonNavigationBar extends StatelessWidget {
  const SegonNavigationBar({
    super.key,
    required this.index,
    required this.changeIndex,
  });

  final int index;
  final void Function(int) changeIndex;

  @override
  Widget build(context) {
    return NavigationBar(
        backgroundColor: Colors.orange,
        onDestinationSelected: changeIndex,
        selectedIndex: index,
        destinations: [
          for (int i = 0; i < HubList.ui.length; i++)
            NavigationDestination(
              icon: Icon(HubList.icon[i]),
              label: HubList.label[i],
            )
        ]);
  }
}

/////////////////////////////////////////////////////////////////////////

Widget listTile(BuildContext context, i, void Function(int) changeIndex) {
  return OutlinedButton(
      style: OutlinedButton.styleFrom(
        side: const BorderSide(color: Colors.transparent),
        shape: const RoundedRectangleBorder(borderRadius: BorderRadius.zero),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: [
        Padding(
          padding: const EdgeInsets.fromLTRB(0,p,p,p),
          child: Icon(HubList.icon[i], size: 32 * n, color: Colors.black),
        ),
        Padding(
          padding: const EdgeInsets.all(p),
          child: Text(HubList.label[i],
              style: const TextStyle(fontSize: 16 * n, color: Colors.black)),
        )
      ]),
      onPressed: () {
        changeIndex(i);
      });
}

class HubList {
  static final ui = [
    const HubUI(),
    const SearchUI(),
    const Post(),
    const NotificationUI(),
    const ProfileUI(),
  ];

  static final floatingActionButton = [
    null,
    searchFloatingActionButton(),
    null,
    null,
    null,
  ];

  static const label = ["home", "search", "post", "notify", "profile"];

  static const icon = [
    Icons.home,
    Icons.search,
    Icons.image,
    Icons.notifications,
    Icons.person
  ];
}
