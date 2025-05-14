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
  let sum: number[] = [];
  let suma: number = 0;
  for (let index = 0; index < nums.length; index++) {
    suma = nums[index];
    if (suma === target) {
      sum.push(index);
      return sum;
    } else if (suma > target) {
      continue;
    } else {
      sum.push(index);
    }
    for (let index2 = index + 1; index2 < nums.length; index2++) {
      suma += nums[index2];
      if (suma === target) {
        sum.push(index2);
        return sum;
      } else if (suma > target) {
        if ((nums[index] + nums[index2]) === target) {
          sum = [index, index2];
          return sum;
        } else if ((nums[index] + nums[index2]) < target) {
          sum.push(index2);
        } else {
          suma -= nums[index2];
          continue;
        }
      } else {
        sum.push(index2);
      }
    }
  }
  return sum;
}

console.log(twoSum([8, 2, 8, 2, 2], 6));
