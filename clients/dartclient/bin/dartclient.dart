import 'package:pocketbase/pocketbase.dart';

final pb = PocketBase('http://192.168.1.253:8090');

Future<void> controlLight(String roomName, String lightName, String duration) async {
  await pb.send("/rooms/$roomName/lights/$lightName", query: {"duration": duration});
}

void main() async {
  await controlLight("MakerFaire", "LuceTavolo", "1000");
}

