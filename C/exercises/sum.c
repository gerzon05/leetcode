// Suma de dos números usando punteros:

// Declara dos punteros a enteros.
// Reserva memoria para ambos.
// Asigna valores a las ubicaciones de memoria.
// Crea una función que reciba como parámetros los dos punteros y devuelva la
// suma de los valores a los que apuntan. Imprime el resultado. Libera la
// memoria. Ejemplo de salida: Si los valores son 10 y 20, la salida sería: "La
// suma es: 30"

// Crea un programa que asigne memoria para dos caracteres, los inicialice con
// valores diferentes y los intercambie utilizando punteros.
#include <stdio.h>
#include <stdlib.h>

int sum(int *value1, int *value2) { return (*value1 + *value2); }

int example(int input1, int input2) {
  int *valor1 = malloc(sizeof(int));
  int *valor2 = malloc(sizeof(int));

  if (valor1 == NULL || valor2 == NULL) {
    fprintf(stderr, "Error: No se pudo asignar memoria.\n");
    return (1);
  }

  *valor1 = input1;
  *valor2 = input2;

  int result = sum(valor1, valor2);

  printf("La suma es: %d\n", result);

  free(valor1);
  free(valor2);

  return (0);
}

int example2() {
  char *letter1 = malloc(sizeof(char));
  char *letter2 = malloc(sizeof(char));

  if (letter1 == NULL || letter2 == NULL) {
    fprintf(stderr, "Error: No se pudo asignar memoria.\n");
    return 1;
  }

  *letter1 = 'R';
  *letter2 = 'G';

  char temp = *letter1;
  *letter1 = *letter2;
  *letter2 = temp;

  printf("Primer valor: %c\n", *letter1);
  printf("Segundo valor: %c\n", *letter2);

  free(letter1);
  free(letter2);

  return 0;
}

int main() {
  example(23, 20);
  example2();

  return (0);
}