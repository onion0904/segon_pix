import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../commons/sign_form.dart';

class Splash extends StatelessWidget {
  const Splash({super.key});

  @override
  Widget build(context) {
    return Center(child: SignForm());
  }
}
