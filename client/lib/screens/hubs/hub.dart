import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'home.dart';
import 'profile.dart';
import 'search.dart';
import 'notification.dart';
import 'post.dart';
import 'package:go_router/go_router.dart';

/////////////////////////////////////////////////////////////////////////

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
        body: HubList.ui[index.value],
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

// Widget navigationBar(void Function(int) changeIndex) {
//TODO navigationBarを変える
// return ConvexAppBar(
//     style: TabStyle.fixed,
//     height: 64,
//     curveSize: 128,
//     elevation: 4,
//     backgroundColor: Colors.orange,
//     items: [
//       for (int i = 0; i < HubList.ui.length; i++)
//         TabItem(icon: HubList.icon[i], title: HubList.label[i])
//     ],
//     onTap: changeIndex);
// }

/////////////////////////////////////////////////////////////////////////

Widget listTile(BuildContext context, i, void Function(int) changeIndex) {
  return GestureDetector(
      child: Padding(
          padding: const EdgeInsets.all(16),
          child: Row(children: [
            const SizedBox(width: 8),
            Icon(HubList.icon[i], size: 32),
            const SizedBox(width: 8),
            Text(HubList.label[i], style: const TextStyle(fontSize: 16)),
          ])),
      onTap: () {
        changeIndex(i);
      });
}

class HubList {
  static const ui = [
    HubUI(),
    SearchUI(),
    Post(),
    NotificationUI(),
    ProfileUI(),
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
