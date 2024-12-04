// =============================================================================
// Add Two Numbers (Suma dos números)
// =============================================================================
//
// Enunciado en Inglés:
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Enunciado en Español:
// Se le dan dos listas enlazadas no vacías que representan dos números enteros no negativos. Los dígitos se almacenan en orden inverso, y cada uno de sus nodos contiene un solo dígito. Suma los dos números y devuelve la suma como una lista enlazada.

// Puede suponer que los dos números no contienen ningún cero a la izquierda, excepto el propio número 0.
//
// Ejemplo:
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
//

// =============================================================================

class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  return l1 || l2
};

const l1 = new ListNode(342,[2,4,3]);

console.log(addTwoNumbers([],[]));
