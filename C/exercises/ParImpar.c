#include <stdio.h>
#include <stdlib.h>

#define ARRAYS_OF_NUMBER_LEN 10

int main() {
  int arrays_number[ARRAYS_OF_NUMBER_LEN] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  int *arrays_of_par = NULL;
  int *arrays_of_impar = NULL;

  int size_of_arrays_par = 0;
  int size_of_arrays_impar = 0;

  arrays_of_par = (int *)malloc(ARRAYS_OF_NUMBER_LEN * sizeof(int));
  arrays_of_impar = (int *)malloc(ARRAYS_OF_NUMBER_LEN * sizeof(int));

  if (arrays_of_par == NULL || arrays_of_impar == NULL) {
    printf("La memoria no fue asiganda");
    return (1);
  };

  for (int i = 0; i < ARRAYS_OF_NUMBER_LEN; i++) {
    if (arrays_number[i] % 2 == 0) {
      arrays_of_par[size_of_arrays_par++] = arrays_number[i];
    } else {
      arrays_of_impar[size_of_arrays_impar++] = arrays_number[i];
    }
  }

  arrays_of_par =
      (int *)realloc(arrays_of_par, size_of_arrays_par * sizeof(int));
  arrays_of_impar =
      (int *)realloc(arrays_of_impar, size_of_arrays_impar * sizeof(int));

  if (arrays_of_par == NULL && size_of_arrays_par > 0) {
    printf("Error al reasignar la memoria del arrays par\n");
    return (1);
  }
  if (arrays_of_impar == NULL && size_of_arrays_impar > 0) {
    printf("Error al reasignar la memoria del arrays impar\n");
    return (1);
  }

  printf("Numeros Pares: \n");
  for (int i = 0; i < size_of_arrays_par; i++) {
    printf(" %d ", arrays_of_par[i]);
  }
  printf("\n");

  printf("Numeros Impares: \n");
  for (int i = 0; i < size_of_arrays_impar; i++) {
    printf(" %d ", arrays_of_impar[i]);
  }

  printf("\n");
  free(arrays_of_par);
  free(arrays_of_impar);

  return (0);
}