#include <stdio.h>

int main() {

  FILE *file = fopen("input.txt", "r");

  fseek(file, 0, SEEK_END);
  long file_size = ftell(file);
  fseek(file, 0, SEEK_SET);

  char content[file_size + 1];

  fread(content, 1, file_size, file);
  content[file_size] = '\0';
  fclose(file);

  int count = 0;
  for (long i = 0; i < file_size; i++) {
    char current = content[i];
    if (current == '(') {
      count++;
    } else if (current == ')') {
      count--;
    }
  }

  printf("%i\n", count);
  return 0;
}
