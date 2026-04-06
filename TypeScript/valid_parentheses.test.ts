import { assert } from "@std/assert";
import { isValidParentheses } from "./valid_parentheses.ts";

Deno.test("Should return true", () => {
  const result = isValidParentheses("()");
  const TOTAL_REFERENCE = true;
  assert(result === TOTAL_REFERENCE);
});

Deno.test("Should return true", () => {
  const result = isValidParentheses("()[]{}");
  const TOTAL_REFERENCE = true;
  assert(result === TOTAL_REFERENCE);
});

Deno.test("Should return false", () => {
  const result = isValidParentheses("(]");
  const TOTAL_REFERENCE = false;
  assert(result === TOTAL_REFERENCE);
});

Deno.test("Should return false", () => {
  const result = isValidParentheses("([)]");
  const TOTAL_REFERENCE = false;
  assert(result === TOTAL_REFERENCE);
});

Deno.test("Should return true", () => {
  const result = isValidParentheses("{[]}");
  const TOTAL_REFERENCE = true;
  assert(result === TOTAL_REFERENCE);
});
