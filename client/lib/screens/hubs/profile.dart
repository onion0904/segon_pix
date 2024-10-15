import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../../logic/db/user_manager.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import '../../model/user.dart';

const double p = 32;

class ProfileUI extends StatelessWidget {
  const ProfileUI({super.key});

  @override
  Widget build(context) {
    return const Column(children: [
      // ShowUserInfo(),
      Padding(padding: EdgeInsets.all(p), child: LogOut())
    ]);
  }
}

class ShowUserInfo extends ConsumerWidget {
  const ShowUserInfo({super.key});

  @override
  Widget build(context, ref) {
    final user = ref.watch(userProvider);
    return Center(
        child: Container(
            padding: const EdgeInsets.all(32),
            //TODO decoration
            child: Row(mainAxisAlignment: MainAxisAlignment.center, children: [
              if (user != null)
                Image.network(user.icon),
              if(user != null)
                Column(mainAxisAlignment: MainAxisAlignment.center, children: [
                  Text("Name: ${user.name}"),
                  Text("Email: ${user.email}")
                ])
            ])));
  }
}

class LogOut extends HookConsumerWidget {
  const LogOut({super.key});

  @override
  Widget build(context, ref) {
    return ElevatedButton(
        onPressed: () async {
          await UserIdManager.removeUserId();
          ref.read(userProvider.notifier).state = null;
          WidgetsBinding.instance.addPostFrameCallback((_) {
            context.go("/");
          });
        },
        child: const Text("Log Out"));
  }
}
