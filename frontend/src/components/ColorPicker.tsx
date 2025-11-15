import type React from "react";

const ColorPicker = ({
    color,
    setColor,
}: {
    color: string;
    setColor: React.Dispatch<React.SetStateAction<string>>;
}) => {
    return (
        <div
            className="relative w-[40px] max-[619px]:w-[35px] min-h-full rounded-lg"
            style={{ backgroundColor: color }}
        >
            {" "}
            <input
                className="inline-block absolute cursor-pointer w-full h-full opacity-0"
                type="color"
                required
                value={color}
                onChange={(e) => {
                    setColor(e.target.value);
                }}
                onInput={() => {}}
            />
        </div>
    );
};

export default ColorPicker;
