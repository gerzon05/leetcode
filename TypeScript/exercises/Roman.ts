// =============================================================================
// Roman to Integer (De romano a entero)
// =============================================================================
//
// Enunciado en Inglés:
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
// Enunciado en Español:
// Los números romanos se representan mediante siete símbolos diferentes: I, V, X, L, C, D y M.

// Símbolo Valor
// I 1
// V 5
// X 10
// L 50
// C 100
// D 500
// M 1000
// Por ejemplo, 2 se escribe como II en números romanos, simplemente dos unos sumados. El 12 se escribe XII, que es simplemente X + II. El número 27 se escribe XXVII, que es XX + V + II.

// Los números romanos suelen escribirse de mayor a menor, de izquierda a derecha. Sin embargo, el número cuatro no es IIII. En su lugar, el número cuatro se escribe IV. Como el uno está antes del cinco, lo restamos y obtenemos cuatro. El mismo principio se aplica al número nueve, que se escribe IX. Hay seis casos en los que se utiliza la resta:

// I puede colocarse antes de V (5) y X (10) para formar 4 y 9. 
// X puede anteponerse a L (50) y C (100) para formar 40 y 90. 
// C puede colocarse antes de D (500) y M (1000) para formar 400 y 900.
// Dado un número romano, conviértelo en un número entero.
//
// Ejemplo:
//
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
//
// =============================================================================