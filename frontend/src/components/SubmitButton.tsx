interface submitButtonProps {
  text: string;
  handleSubmit: () => void;
  aria_label: string;
}
const SubmitButton = ({
  text,
  handleSubmit,
  aria_label,
}: submitButtonProps) => {
  return (
    <button
      aria-label={aria_label}
      onClick={handleSubmit}
      className="w-full py-[12px] rounded-[8px] cursor-pointer bg-accent mx-auto"
    >
      {text}
    </button>
  );
};

export default SubmitButton;
