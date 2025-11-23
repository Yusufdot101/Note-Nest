interface submitButtonProps {
    text: string;
    handleSubmit: () => void;
    type?: "button" | "submit";
    aria_label: string;
    bgColor?: string;
    textColor?: string;
    disabled?: boolean;
}
const SubmitButton = ({
    text,
    type,
    handleSubmit,
    aria_label,
    bgColor,
    textColor,
    disabled,
}: submitButtonProps) => {
    return (
        <button
            disabled={disabled}
            aria-label={aria_label}
            onClick={handleSubmit}
            style={{
                backgroundColor: bgColor ? bgColor : "",
                color: textColor ? textColor : "",
            }}
            type={type ?? "button"}
            className={`${disabled ? "opacity-80 cursor-not-allowed" : "opacity-100 cursor-pointer"} w-full py-[12px] rounded-[8px] bg-accent mx-auto`}
        >
            {text}
        </button>
    );
};

export default SubmitButton;
