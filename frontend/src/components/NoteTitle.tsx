import ColorPicker from "./ColorPicker";

interface NoteTitleProps {
    title: string;
    color: string;
    setTitle: React.Dispatch<React.SetStateAction<string>>;
    setColor: React.Dispatch<React.SetStateAction<string>>;
}

const NoteTitle = ({ title, setTitle, color, setColor }: NoteTitleProps) => {
    return (
        <div className="text-text flex flex-col gap-[4px] text-[20px]">
            <label htmlFor="title">
                Add title
                <span className="text-red-500">*</span>
            </label>
            <div className="flex gap-[12px]">
                <input
                    type="text"
                    placeholder="Title"
                    value={title}
                    id="title"
                    onChange={(e) => setTitle(e.target.value)}
                    required
                    style={{ border: `1px solid ${color}` }}
                    className="w-full py-[4px] px-[12px] rounded-[8px] outline-none"
                />
                <ColorPicker color={color} setColor={setColor} />
            </div>
        </div>
    );
};

export default NoteTitle;
