// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Spanish:

// Dada una cadena s que sólo contiene los caracteres '(', ')', '{', '}', '[' y ']', determine si la cadena de entrada es válida.

// Una cadena de entrada es válida si

// Los corchetes abiertos deben estar cerrados por el mismo tipo de corchetes.
// Los corchetes abiertos deben cerrarse en el orden correcto.
// Cada corchete cerrado tiene su correspondiente corchete abierto del mismo tipo.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

export function isValidParentheses(s: string): boolean {
  const bracketsList = [
    { value: "{", type: "open" },
    { value: "[", type: "open" },
    { value: "(", type: "open" },
    { value: ")", type: "close" },
    { value: "]", type: "close" },
    { value: "}", type: "close" },
  ];

  const brackets = Map.groupBy(bracketsList, ({ type }) =>
    type === "open" ? "open" : "close",
  );

  const parents = "() {} []";

  const result: string[] = [];

  for (const character of s) {
    if (brackets.get("open")?.find((b) => b.value === character)) {
      result.push(character);
    } else if (brackets.get("close")?.find((b) => b.value === character)) {
      if (result.length === 0) {
        return false;
      }

      result.push(character);

      const open = result[Number(result.length - 2)];
      const close = result[Number(result.length - 1)];

      if (parents.includes(open + close)) {
        if (result.length === 2) {
          result.pop();
          result.pop();
          break;
        }
        result.pop();
        result.pop();
      }
    } else {
      continue;
    }
  }

  return result.length === 0;
}
