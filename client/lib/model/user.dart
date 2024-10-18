import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

class SimpleImage {
  final int imageID;
  final String imageURL;

  SimpleImage({
    required this.imageID,
    required this.imageURL,
  });

  factory SimpleImage.fromJson(Map<String, dynamic> json) {
    return SimpleImage(
      imageID: json["ID"],
      imageURL: json["URL"],
    );
  }
}

Future<List<SimpleImage>> getRecentImages() async {
  final url = Uri.http("localhost:8080", "/get/list/recent");
  final response = await http.get(url, headers: {"Content-Type": "application/json"});

  if (response.statusCode == 200) {
    final List<dynamic> jsonList = jsonDecode(response.body) as List<dynamic>;
    final images = jsonList.map((item) => SimpleImage.fromJson(item)).toList();
    return images;
  } else {
    throw Exception('Failed to load images');
  }
}

class HubUI extends HookWidget {
  const HubUI({super.key});

  @override
  Widget build(BuildContext context) {
    final images = useState<List<SimpleImage>>([]);
    final isLoading = useState(true);

    useEffect(() {
      getRecentImages().then((fetchedImages) {
        images.value = fetchedImages;
        isLoading.value = false;
      }).catchError((error) {
        print("Error loading images: $error");
      });

      return null;
    }, []);

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
          child: isLoading.value
              ? const Center(child: CircularProgressIndicator())
              : GridView.count(
                  crossAxisCount: 2,
                  children: images.value.map((image) {
                    return Container(
                      padding: const EdgeInsets.all(2),
                      child: Image.network(
                        image.imageURL,
                        fit: BoxFit.cover,
                      ),
                    );
                  }).toList(),
                ),
        ),
      ],
    );
  }
}
