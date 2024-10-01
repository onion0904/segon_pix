import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:client/screens/hubs/hub.dart';
import 'screens/logins/logins.dart';

void main() {
  runApp(const MainApp());
}

final GoRouter _router = GoRouter(routes: [
  GoRoute(path: "/", builder: (context, stage) => const Logins()),
  GoRoute(path: "/hub", builder: (context, stage) => const Hub()),
]);

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(routerConfig: _router);
  }
}