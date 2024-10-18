import 'package:flutter/material.dart';

const double elevation = 4;
const double padding = 8;
const double width = 1;
const double height = 46;

class SegonButton extends StatelessWidget {
  const SegonButton({
    super.key,
    required this.handler,
    required this.label,
    this.minSize = 200,
    this.maxSize,
  });

  final void Function()? handler;
  final String label;
  final double minSize;
  final double? maxSize;

  @override
  Widget build(context) {
    return Container(
        width: maxSize,
        padding: EdgeInsets.all(padding),
        child: ElevatedButton(
            style: ElevatedButton.styleFrom(
              minimumSize: Size(minSize, height),
              maximumSize:
                  (maxSize != null) ? Size(maxSize as double, height) : null,
              elevation: elevation,
              backgroundColor: Colors.orange,
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(4)),
            ),
            onPressed: handler,
            child: Text(label,
                style: TextStyle(
                    fontWeight: FontWeight.bold, color: Colors.black))));
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
