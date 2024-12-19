#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <string.h>

#define MAX_LIMIT 100
#define MIN 1

bool word_is_palindrome() {
  printf("Validate if the word is Palindrome \n");

  char word[MAX_LIMIT];
  bool word_palindrome = true;

  printf("Write a word: \n");

  fgets(word, MAX_LIMIT, stdin);

  int lenght = strlen(word) - 1;

  if (lenght <= 1) {
    return word_palindrome;
  }

  for (int i = 0; i < lenght / 2; i++) {
    char current_char = tolower(word[i]);
    char opposite_char = tolower(word[lenght - i - MIN]);

    if (current_char != opposite_char) {
      word_palindrome = false;
      break;
    }
  }

  return word_palindrome;
}

int main() {
  bool palindrome = word_is_palindrome();
  if (palindrome) {
    printf("The word is a palindrome\n");
  } else {
    printf("The word is not a palindrome\n");
  }
  return (0);
}
