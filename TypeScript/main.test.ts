import { assert } from "@std/assert";
import { add } from "./main.ts";

Deno.test(function addTest() {
  assert(add(2, 3), 5);
});
