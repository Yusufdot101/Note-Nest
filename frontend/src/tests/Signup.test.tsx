import { render, screen } from "@testing-library/react"
import { test, expect, vi, describe } from "vitest"
import userEvent from "@testing-library/user-event"
import Signup from "../pages/Signup"
import { signup } from "../utilities/auth/signup"
import { MemoryRouter } from "react-router-dom"

vi.mock("../utilities/auth/signup", () => ({
    signup: vi.fn()
}))

describe("Signup", () => {
    test("shows errors when submitting empty form", async () => {
        render(
            <MemoryRouter>
                <Signup />
            </MemoryRouter>
        )

        const submitButton = screen.getByRole("button", { name: /sign up/i })
        const usernameError = screen.getByRole("paragraph", { name: /username error/i })
        const emailError = screen.getByRole("paragraph", { name: /email error/i })
        const passwordError = screen.getByRole("paragraph", { name: /password error/i })
        await userEvent.click(submitButton)

        // error elements should be visible
        expect(usernameError).toBeVisible()
        expect(emailError).toBeVisible()
        expect(passwordError).toBeVisible()

        expect(signup).not.toBeCalled()
    })

    test("should call signup with correct values when form is valid", async () => {
        render(
            <MemoryRouter>
                <Signup />
            </MemoryRouter>
        )

        const submitButton = screen.getByRole("button", { name: /sign up/i })

        const usernameInput = screen.getByLabelText(/Username/)
        const emailInput = screen.getByLabelText(/Email/)
        const passwordInput = screen.getByLabelText(/Password/)

        await userEvent.type(usernameInput, "yusuf")
        await userEvent.type(emailInput, "ym@gmail.com")
        await userEvent.type(passwordInput, "12345678")
        await userEvent.click(submitButton)

        expect(signup).toHaveBeenCalledWith(
            "yusuf",
            "ym@gmail.com",
            "12345678",
            expect.any(Function) // handleErrors callback
        )
    })
})
