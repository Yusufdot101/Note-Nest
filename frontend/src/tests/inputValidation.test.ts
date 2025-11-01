import { test, expect, describe } from "vitest"
import { getEmailErrorMessages, getPasswordErrorMessages, getUsernameErrorMessages } from "../utilities/inputValidation"

describe("input validation", () => {
    test("valid inputs", () => {
        const username = "yusuf"
        const email = "ym@gmail.com"
        const password = "12345678"
        expect(getUsernameErrorMessages(username)).toEqual("")
        expect(getEmailErrorMessages(email)).toEqual("")
        expect(getPasswordErrorMessages(password)).toEqual("")
    })
    test("invalid inputs", () => {
        const username = ""
        const email = "ym@gmail.com"
        const password = "123"
        expect(getUsernameErrorMessages(username)).toEqual("Username must be at least 2 characters.")
        expect(getEmailErrorMessages(email)).toEqual("")
        expect(getPasswordErrorMessages(password)).toEqual("Password must be at least 8 characters.")
    })
})
