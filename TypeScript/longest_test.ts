import { assert } from "@std/assert";
import { longestCommonPrefix } from "./Longest.ts";

Deno.test("Should return fl", () => {
  const RESPONSE = longestCommonPrefix(["flower", "flow", "flight"]);
  const VALUE = "fl";
  assert(RESPONSE === VALUE);
});

Deno.test('Should return ""', () => {
  const RESPONSE = longestCommonPrefix(["dog", "racecar", "car"]);
  const VALUE = "";
  assert(RESPONSE === VALUE);
});

Deno.test("Should return program", () => {
  const RESPONSE = longestCommonPrefix([
    "programacion",
    "programador",
    "programming",
  ]);
  const VALUE = "program";
  assert(RESPONSE === VALUE);
});
