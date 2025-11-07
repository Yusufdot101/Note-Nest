import { render, screen } from "@testing-library/react"
import { test, expect, vi, describe } from "vitest"
import userEvent from "@testing-library/user-event"
import Login from "../pages/Login"
import { login } from "../utilities/auth/login"
import { MemoryRouter } from "react-router-dom"

vi.mock("../utilities/auth/login", () => ({
    login: vi.fn()
}))

describe("Login", () => {
    test("should show errors when submitting empty form", async () => {
        render(
            <MemoryRouter>
                <Login />
            </MemoryRouter>
        )

        const submitButton = screen.getByRole("button", { name: /login/i })
        const emailError = screen.getByRole("paragraph", { name: /email error/i })
        const passwordError = screen.getByRole("paragraph", { name: /password error/i })
        await userEvent.click(submitButton)

        // error elements should be visible
        expect(emailError).toBeVisible()
        expect(passwordError).toBeVisible()

        expect(login).not.toBeCalled()
    })

    test("should call login with correct values when form is valid", async () => {
        render(
            <MemoryRouter>
                <Login />
            </MemoryRouter>
        )

        const submitButton = screen.getByRole("button", { name: /login/i })

        const emailInput = screen.getByLabelText(/Email/)
        const passwordInput = screen.getByLabelText(/Password/)

        await userEvent.type(emailInput, "ym@gmail.com")
        await userEvent.type(passwordInput, "12345678")
        await userEvent.click(submitButton)

        expect(login).toHaveBeenCalledWith(
            "ym@gmail.com",
            "12345678",
            expect.any(Function) // handleErrors callback
        )
    })
})
