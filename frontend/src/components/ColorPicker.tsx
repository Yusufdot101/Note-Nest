import { useRef } from "react";

const ColorPicker = ({
    color,
    handleChange,
}: {
    color: string;
    handleChange?: (value: string) => void;
}) => {
    const ref = useRef<HTMLInputElement>(null);
    return (
        <div
            className="relative w-[40px] max-[619px]:w-[35px] min-h-full rounded-lg"
            style={{ backgroundColor: color }}
            onClick={(e) => e.stopPropagation()}
            role="button"
            aria-label="open color picker to change card color"
            tabIndex={handleChange ? 0 : -1}
            onKeyDown={(e) => {
                e.stopPropagation();
                if (!handleChange) return;
                if (e.key === "Enter" || (e.key === " " && ref.current)) {
                    ref.current?.click();
                }
            }}
        >
            <input
                ref={ref}
                tabIndex={-1}
                className="inline-block absolute cursor-pointer w-full h-full opacity-0"
                type="color"
                required
                disabled={handleChange ? false : true}
                value={color}
                onChange={(e) => {
                    e.stopPropagation();
                    if (!handleChange) return;
                    handleChange(e.target.value);
                }}
            />
        </div>
    );
};

export default ColorPicker;
