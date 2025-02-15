import { twoSum } from "./two_sum.ts";
import { assertArrayIncludes } from "@std/assert/array-includes";

Deno.test("Should return [0, 2]", () => {
  const arrayIndex = twoSum([2, 11, 7, 15], 9) as unknown as Array<number>;

  assertArrayIncludes(arrayIndex, [0, 2]);
});

Deno.test("Should return [0, 1]", () => {
  const arrayIndex = twoSum([2, 7, 11, 15], 9) as unknown as Array<number>;

  assertArrayIncludes(arrayIndex, [0, 1]);
});

Deno.test("Should return [1, 3, 4]", () => {
  const arrayIndex = twoSum([8, 2, 8, 2, 2], 6) as unknown as Array<number>;

  assertArrayIncludes(arrayIndex, [1, 3, 4]);
});
