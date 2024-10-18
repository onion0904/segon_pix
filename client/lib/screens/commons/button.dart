import 'package:flutter/material.dart';

const double elevation = 10;

class SegonButton extends StatelessWidget {
  SegonButton({
    super.key,
    required this.handler,
    required this.label,
    isFixed,
  });

  final void Function()? handler;
  final String label;
  bool isFixed = false;

  @override
  Widget build(context) {
    final double a = 200; //MediaQuery.sizeOf(context).width;
    return SizedBox(
        width: a,

        child: OutlinedButton(
            style: OutlinedButton.styleFrom(
                minimumSize: isFixed ? Size(a * 0.1, 10) : null,
                backgroundColor: Colors.grey[100],
                elevation: elevation,
                shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10))),
            onPressed: handler,
            child: Text(label)));
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
