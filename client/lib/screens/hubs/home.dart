import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

class HubUI extends HookWidget {
  const HubUI({super.key});

  @override
  Widget build(BuildContext context) {
    useEffect((){
      final
    });

    return Column(
      children: [
        const Padding(
          padding: EdgeInsets.all(8),
          child: Row(children: [
            SizedBox(width: 16),
            Icon(Icons.favorite, size: 32),
            SizedBox(width: 8),
            Text("Recent Image", style: TextStyle(fontSize: 20))
          ]),
        ),
        Expanded(
          child: GridView.count(
            crossAxisCount: 2,
            children: List.generate(64, (index) {
              return Container(
                padding: const EdgeInsets.all(2),
                child: Image.network(
                  "https://dthezntil550i.cloudfront.net/i9/latest/i92307171113441820023299212/1280_960/326b00ed-3038-459e-96aa-a1692a925864.png",
                  fit: BoxFit.cover,
                ),
              );
            }),
          ),
        ),
      ],
    );
  }
}
