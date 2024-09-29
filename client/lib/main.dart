import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:client/screens/hub.dart';
import 'package:client/screens/splash.dart';

void main() {
  runApp(const MainApp());
}

final GoRouter _router = GoRouter(routes: [
  GoRoute(path: "/", builder: (context, stage) => const Splash()),
  GoRoute(path: "/hub", builder: (context, stage) => const Hub()),
]);

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(routerConfig: _router);
  }
}
