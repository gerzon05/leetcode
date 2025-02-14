import { assert } from "@std/assert/assert";
import { isPalindrome } from "./palindrome_number.ts";

Deno.test("Should return true", () => {
  const result = isPalindrome(121);
  const TOTAL_REFERENCE = true;
  assert(result === TOTAL_REFERENCE);
});

Deno.test("Should return false", () => {
  const result = isPalindrome(214);
  const TOTAL_REFERENCE = false;
  assert(result === TOTAL_REFERENCE);
});
