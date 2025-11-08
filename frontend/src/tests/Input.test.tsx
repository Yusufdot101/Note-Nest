import { render, screen, fireEvent } from "@testing-library/react";
import Input from "../components/Input";
import { test, expect, vi, describe } from "vitest";

describe("Input", () => {
  test("Label text should be dynamic", () => {
    const mockHandleChange = vi.fn();
    render(
      <Input
        labelString="name"
        inputType="text"
        inputValue="10"
        inputId="name"
        inputName="name"
        handleChange={mockHandleChange}
      />,
    );

    const label = screen.getByText("name");
    const input = screen.getByLabelText("name");

    fireEvent.change(input, { target: { value: "example" } });

    expect(label).toHaveTextContent("name");
    expect(label).toBeInTheDocument();
    expect(mockHandleChange).toHaveBeenCalledWith("example");
    expect(input).toBeInTheDocument();
  });
  test("handleChange should be called with the text", () => {
    const mockHandleChange = vi.fn();
    render(
      <Input
        labelString="name"
        inputType="text"
        inputValue="10"
        inputId="name"
        inputName="name"
        handleChange={mockHandleChange}
      />,
    );
    const input = screen.getByLabelText("name");

    fireEvent.change(input, { target: { value: "example" } });
    expect(mockHandleChange).toHaveBeenCalledWith("example");
  });
});
