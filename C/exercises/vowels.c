// counting vowels
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#define MAX_LIMIT_WORD 100

int counting_vowels() {
  int number_of_vowels = 0;
  char word[MAX_LIMIT_WORD];

  printf("Write a word: \n");
  if (fgets(word, MAX_LIMIT_WORD, stdin) == NULL)
  {
    fprintf(stderr, "Error reading the entrance");
    return (1);
  }
  
  int lenght = strlen(word) - 1;

  for (int i = 0; i < lenght; i++)
  {
    word[i] = tolower(word[i]);
    
    if (word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' ||
        word[i] == 'u')
    {
      number_of_vowels++;
    }
  }

  return number_of_vowels;
}

int main(){
  int value_vowels = counting_vowels();
  if (value_vowels > 1 || value_vowels == 0)
  {
    printf("the word has %d vowels.", value_vowels);
  } else {
    printf("the word has %d vowel", value_vowels);
  }
  
  return (0);
}