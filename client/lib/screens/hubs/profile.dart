import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

const double p = 32;

class ProfileUI extends StatelessWidget {
  const ProfileUI({super.key});

  @override
  Widget build(context) {
    return const Column(children: [
      Padding(
        padding: EdgeInsets.all(p),
        child: LogOut()
      )
    ]);
  }
}

class LogOut extends StatelessWidget {
  const LogOut({super.key});

  @override
  Widget build(context) {
    return ElevatedButton(
        onPressed: () {
          context.go("/");
        },
        child: const Text("Log Out"));
  }
}
