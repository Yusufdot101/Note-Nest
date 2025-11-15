interface submitButtonProps {
    text: string;
    handleSubmit: () => void;
    type?: "button" | "submit";
    aria_label: string;
    bgColor?: string;
    textColor?: string;
}
const SubmitButton = ({
    text,
    type,
    handleSubmit,
    aria_label,
    bgColor,
    textColor,
}: submitButtonProps) => {
    return (
        <button
            aria-label={aria_label}
            onClick={handleSubmit}
            style={{
                backgroundColor: bgColor ? bgColor : "",
                color: textColor ? textColor : "",
            }}
            type={type!}
            className="w-full py-[12px] rounded-[8px] cursor-pointer bg-accent mx-auto"
        >
            {text}
        </button>
    );
};

export default SubmitButton;
