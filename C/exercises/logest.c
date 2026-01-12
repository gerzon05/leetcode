#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* longestCommonPrefix(char** strs, int strsSize) {

  if(strsSize == 0) {
    return "";
  }

  char *response = malloc(10);

  for (int i = 0; strs[0][i] != '\0' ; i++) {
      char value = strs[0][i];
      for (int j = 1; j < strsSize; j++) {
          if (strs[j][i] == '\0' || strs[j][i] != value) {
              return response;
          }
      }

        int len = strlen(response);
        response[len] = value;
        response[len + 1] = '\0';
  }

  return response;
}

int main() {

  char *values[3] = {"flower","flow","floght"};

  char *response = longestCommonPrefix(values, 3);

  printf("la respuesta es de: %s\n", response);

  return 0;
}