import { describe, it, expect } from "vitest";
import { addTwoNumbers } from "./addTwoNumbers";

describe("addTwoNumbers function", () => {
    it("adds two numbers correctly", () => {
        expect(addTwoNumbers(1, 2)).toBe(3)
        expect(addTwoNumbers(225, 75)).toBe(300)
    })
    it("incorrect example", () => {
        expect(addTwoNumbers(77, 33)).toBe(110)
    })
})
