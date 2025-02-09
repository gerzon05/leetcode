// Longest Common Prefix
// Easy
// Topics
// Companies
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Español

// Prefijo común más largo
// Fácil
// Temas
// Empresas
// Escriba una función para encontrar la cadena de prefijo común más larga entre una matriz de cadenas.

// Si no hay prefijo común, devuelve una cadena vacía «».

// function longestCommonPrefix(strs: string[]): string {
//   if (!strs.length) return "";
//   let letters: string[] = [];
//   let length: number = strs[1].length
//   for (let index = 0; index < strs.length; index++) {
//       if (strs[index].length < length) length = strs[index].length
//   }
//   for (let index = 0; index < length; index++) {
    
//   }
//   console.log(letters)
//   return "";
// }


// console.log(longestCommonPrefix(["gerzon","geronimo","luis"]))
