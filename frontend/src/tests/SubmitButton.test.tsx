import { render, screen, fireEvent } from "@testing-library/react"
import { test, expect, vi, describe } from "vitest"
import SubmitButton from "../components/SubmitButton"

describe("Input", () => {
    test("Button calls handleSubmit", () => {
        const mockHandleSubmit = vi.fn()
        const text = "example"
        render(<SubmitButton aria_label={"submit button"} text={text} handleSubmit={mockHandleSubmit} />)

        const button = screen.getByRole("button", { name: "submit button" })

        fireEvent.click(button)

        expect(button).toBeInTheDocument()
        expect(button).toHaveTextContent(text)
        expect(mockHandleSubmit).toHaveBeenCalled()
    })
})
