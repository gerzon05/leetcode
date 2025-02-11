// =============================================================================
// Palindrome Number (Número palíndromo)
// =============================================================================
//
// Enunciado en Inglés:
// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.
//
// Enunciado en Español:
// Dado un número entero x, devuelve true si x es un palíndromo
// y falso en caso contrario.
//
// Ejemplo:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
//

function isPalindrome(x: number): boolean {
  const number: number = x
  const Arraynumber = number.toString().split("")
  const ArraysRever: string[] = []
  for (let index = Arraynumber.length-1; index >= 0; index--) {
    ArraysRever.push(Arraynumber[index])
  }
  const palindromeNumber: number = Number(ArraysRever.join(""))

  if (x === palindromeNumber) {
    return true
  }
  return false
};

console.log(isPalindrome(121))