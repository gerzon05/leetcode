// =============================================================================
// Two Sum (Suma de Dos)
// =============================================================================
//
// Enunciado en Inglés:
// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target. You may assume that each input
// would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Enunciado en Español:
// Dado un arreglo de números enteros nums y un entero target, devuelve los
// índices de los dos números tales que sumen target. Puedes asumir que cada
// entrada tendrá exactamente una solución, y no puedes usar el mismo elemento
// dos veces. Puedes devolver la respuesta en cualquier orden.
//
// Ejemplo:
//
// const nums: number[] = [2, 7, 11, 15];
// const target: number = 9;
// const result: number[] = twoSum(nums, target); // [0, 1]
//

// =============================================================================
#include <stdio.h>

void twoSum(int nums[], int numsSize, int target) {
  int result[2];
  int suma = 0;
  for (int i = 0; i < numsSize; i++) {
    for (int j = i + 1; j < numsSize; j++) {
      suma = nums[i] + nums[j];
      if (suma == target) {
        result[0] = i;
        result[1] = j;
        printf("Los indices son: %d y %d\n", result[0], result[1]);
        return;
      }
      suma = 0;
    };
  };
};

int main() {
  int arr[] = {2, 7, 11, 15};
  twoSum(arr, 4, 9);
  return 0;
};