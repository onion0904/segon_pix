import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import '../../logic/db/user_manager.dart';
import '../commons/button.dart';

const double p = 32;
const double width = 2000;

class ProfileUI extends StatelessWidget {
  const ProfileUI({super.key});

  @override
  Widget build(context) {
    return Column(children: [
      // ShowUserInfo(),
      Padding(
          padding: EdgeInsets.all(p),
          child: SegonButton(
              maxSize: width,
              handler: () async {
                await UserManager.resetUserManager();
                if (context.mounted) context.go("/");
              },
              label: "Sign out"))
    ]);
  }
}

class ShowUserInfo extends ConsumerWidget {
  const ShowUserInfo({super.key});

  @override
  Widget build(context, ref) {
    return Center(
        child: Container(
            padding: const EdgeInsets.all(32),
            //TODO decoration
            child: Row(mainAxisAlignment: MainAxisAlignment.center, children: [
              if (UserManager.user!.icon != "")
                Image.network(UserManager.user!.icon),
              Column(mainAxisAlignment: MainAxisAlignment.center, children: [
                Text("Name: ${UserManager.user!.name}"),
                Text("Email: ${UserManager.user!.email}")
              ])
            ])));
  }
}
