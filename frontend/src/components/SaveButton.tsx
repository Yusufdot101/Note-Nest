type SaveButtonProps = {
    onToggle: () => Promise<void>;
    count: number;
    saved: boolean;
};

const SaveButton = ({ onToggle, count, saved }: SaveButtonProps) => {
    return (
        <div className="flex items-center gap-x-[4px]">
            <svg
                onClick={() => {
                    onToggle();
                }}
                aria-label="save note"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => {
                    e.preventDefault();
                    if (e.key === "Enter" || e.key === " ") {
                        onToggle();
                    }
                }}
                viewBox="0 0 24 24"
                fill={`${saved ? "currentColor" : "none"}`}
                xmlns="http://www.w3.org/2000/svg"
                className="cursor-pointer w-[30px] h-[30px]"
                stroke={`${saved ? "none" : "currentColor"}`}
            >
                <path d="M6.75 6L7.5 5.25H16.5L17.25 6V19.3162L12 16.2051L6.75 19.3162V6Z" />
            </svg>
            <span>{count}</span>
        </div>
    );
};

export default SaveButton;
