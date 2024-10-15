import 'package:flutter/material.dart';

const double elevation = 10;

class SegonButton extends StatelessWidget {
  const SegonButton({
    super.key,
    required this.handler,
    required this.label,
  });

  final void Function()? handler;
  final String label;

  @override
  Widget build(context) {
    return OutlinedButton(
        style: OutlinedButton.styleFrom(
            backgroundColor: Colors.grey[100],
            elevation: elevation,
            shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(10))),
        onPressed: handler,
        child: Text(label));
  }
}

class ImageButton extends StatelessWidget {
  const ImageButton({
    super.key,
    required this.imageUri,
  });

  final String imageUri;

  @override
  Widget build(context) {
    return GestureDetector(child: Image.network(imageUri));
  }
}
