import romanToInt from "./number_roman.ts";
import { assert } from "@std/assert";

Deno.test("Should return 4", () => {
  const result = romanToInt("IV");
  const TOTAL_REFERENCE = 4;

  assert(result === TOTAL_REFERENCE);
});

Deno.test("optime the characters that belong to the Roman numerals", () => {
  const result = romanToInt("XXVnhII");
  const TOTAL_REFERENCE = 27;

  assert(result === TOTAL_REFERENCE);
});

Deno.test("A Roman numeral cannot be repeated more than three times.", () => {
  const result = romanToInt("IIII");
  const TOTAL_REFERENCE = "El Numero Romano se repite mas de 3 veces";

  assert(result === TOTAL_REFERENCE);
});

Deno.test("Prohibited Roman numerals", () => {
  const result = romanToInt("VX");
  const TOTAL_REFERENCE = "El numero Romano no Existe";

  assert(result === TOTAL_REFERENCE);
});
