#include <stdio.h>
#include <stdlib.h>

#define SIZE 5

void imprimir_array(int arr[], int size){
  for (int i = 0; i < size; i++){
    printf("%d ", arr[i]);
  }
  
}

int* create_arrays(int size){
  int *array = (int *)malloc(size * sizeof(int));

  for (int i = 0; i < size; i++){
    array[i] = i * 2;
  }

  return array;
}

void duplicate_elements(int *array, int size){
  for (int i = 0; i < size; i++){
    array[i] += 2;
  }
}

int** create_matriz(int filas, int colunms){
  int **matriz = (int **)malloc(filas * sizeof(int *));
  for (int i = 0; i < filas; i++){
    matriz[i] = (int *)malloc(colunms * sizeof(int));
  }
  return matriz;
}

void liberar_matriz(int **matriz, int filas){
  for (int i = 0; i < filas; i++){
    free(matriz[i]);
  }
  free(matriz);
}

void par_impar_numbers(int len) {
  int **arr = create_matriz(2, len/2);

  int par = 0;
  int impar = 0;
  for (int j = 0; j < len; j++) {
    if (j % 2 == 0) {
      arr[0][par] = j;
      par++;
    } 
  }
  for (int j = 0; j < len; j++) {
    if (j % 2 != 0) {
      arr[1][impar] = j;
      impar++;
    } 
  }

  for (int i = 0; i < 2; i++) {
    for (int j = 0; j < len/2; j++) {
      printf("%d\t", arr[i][j]);
    }
    printf("\n");
  }

  free(arr);
}

int main() {
  // Arrays
  int nums[SIZE];
  // Asignar decalracion de valores
  int years[SIZE] = {2023, 2024, 2025, 2026};
  // arrays con tamaño automatico
  int days[] = {1, 2, 3, 4, 5};
  // Arrays bidiemnsional
  int matrizy[2][2] = {{1, 2}, {3, 4}};
  int matrizx[3][3] = {{1, 2, 3}, {4, 5, 6}, {7,8,9}};

  // Manipulacion de arrays
  // Acceder a un valor
  printf("%d\n", years[0]);
  
  // Modificar un valor
  years[3] = 2027;

  // Imprimir un array
  for (int i = 0; i < SIZE; i++){
    printf("%d\n", years[i]);
  }

  // Imprimir un array con una funcion
  imprimir_array(years, SIZE);

  // Arrays Dinamicos
  int *arrays = (int *)malloc(SIZE * sizeof(int));

  // llenar el array
  for (int i = 0; i < SIZE; i++){
    arrays[i] = i * 10;
  }

  for (int i = 0; i < SIZE; i++){
    printf("%d\n", arrays[i]);
  }
  
  // Liberar memoria
  free(arrays);

  // Mostrar un array de una funcion
  int *mi_array = create_arrays(SIZE);

  // Imprimir un array original
  for (int i = 0; i < SIZE; i++){
    printf("%d\n", mi_array[i]);
  }


  // Duplicar elementos de un array
  duplicate_elements(years, SIZE);
  for (int i = 0; i < SIZE; i++){
    printf("%d\n", mi_array[i]);
  }

  free(mi_array);

  // Arrays multidimensionales
  int **matriz = create_matriz(SIZE, SIZE);

  // rellear la matriz
  for (int i = 0; i < SIZE; i++){
    for (int j = 0; j < SIZE; j++){
      matriz[i][j] = i * SIZE + j;
    }
  }
  
  // Imprimir la matriz
  for (int i = 0; i < SIZE; i++){
    for (int j = 0; j < SIZE; j++){
      printf("%d\t", matriz[i][j]);
    }
    printf("\n");
  }
  
  liberar_matriz(matriz, SIZE);

  // ejercicio numeros pares e impares
  printf("Numeros pares e impares\n");
  par_impar_numbers(10);

  return (0);
}