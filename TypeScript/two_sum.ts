// =============================================================================
// Two Sum (Suma de Dos)
// =============================================================================
//
// Enunciado en Inglés:
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Enunciado en Español:
// Dado un arreglo de números enteros nums y un entero target, devuelve los índices de los dos números tales que sumen target.
// Puedes asumir que cada entrada tendrá exactamente una solución, y no puedes usar el mismo elemento dos veces.
// Puedes devolver la respuesta en cualquier orden.
//
// Ejemplo:
//
// const nums: number[] = [2, 7, 11, 15];
// const target: number = 9;
// const result: number[] = twoSum(nums, target); // [0, 1]
//

// =============================================================================

export function twoSum(nums: number[], target: number): number[] {
  const sum: number[] = [];
  let suma: number = 0;
  for (let index = 0; index < nums.length; index++) {
    for (let index2 = index + 1; index2 < nums.length; index2++) {
      if (sum.length > 0) {
        suma += nums[index2];
      } else {
        suma += nums[index] + nums[index2];
      }
      if (suma > target) {
        suma -= nums[index] + nums[index2];
        console.log("res");
      } else if (suma < target) {
        if (sum.length > 0) {
          sum.push(index2);
          index++;
        } else {
          sum.push(index, index2);
        }
      } else {
        if (sum.length == 0) {
          sum.push(index, index2);
        } else {
          sum.push(index2);
        }

        index += nums.length;
        index2 += nums.length;
      }
    }
  }
  return sum;
}
