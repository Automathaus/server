import 'package:pocketbase/pocketbase.dart';

void main() async {
  final pb = PocketBase('http://192.168.1.253:8090');
  await pb.send("/print-body", query: { "abc": 123 });
}
