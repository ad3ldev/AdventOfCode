import 'dart:io';

void main() async {
  final file = File('input.txt');
  var count = 0;
  
  try {
    final content = await file.readAsString();
    for (final char in content.runes) {
      final ch = String.fromCharCode(char);
      if (ch == '(') {
        count++;
      } else if (ch == ')') {
        count--;
      }
    }
    print(count);
  } catch (e) {
    print('Error: $e');
  }
}
