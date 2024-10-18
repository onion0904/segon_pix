import 'package:flutter/material.dart';

const double elevation = 4;
const double padding = 8;
const double width = 1;

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
    return Container(
        width: a,
        padding: EdgeInsets.all(padding),
        child: ElevatedButton(
            style: ElevatedButton.styleFrom(
                minimumSize: isFixed ? Size(a * 0.1, 10) : null,
                elevation: elevation,
                backgroundColor: Colors.orange,
                shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(4)),
            ),
            onPressed: handler,
            child: Text(
              label,
              style: TextStyle(
                fontWeight: FontWeight.bold,
                color: Colors.black
              )
            )));
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
