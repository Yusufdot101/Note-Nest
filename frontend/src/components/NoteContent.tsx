import { useRef, useState } from "react";
import ContentHeader from "./ContentHeader";
import ReactMarkdown from "react-markdown";

interface NoteContentProps {
    color: string;
    content: string;
    setContent: React.Dispatch<React.SetStateAction<string>>;
}

const NoteContent = ({ color, content, setContent }: NoteContentProps) => {
    const [mode, setMode] = useState("markdown");

    const [contentHeight, setContentHeight] = useState(500);
    const textareaRef = useRef<HTMLTextAreaElement>(null);

    const handleResize = () => {
        if (textareaRef.current) {
            const { offsetHeight } = textareaRef.current;
            setContentHeight(offsetHeight);
        }
    };
    return (
        <div className="text-text flex flex-col gap-[4px]">
            <label htmlFor="content" className="text-[20px]">
                Add content
                <span className="text-[red]">*</span>
            </label>
            <div
                style={{ border: `1px solid ${color}` }}
                className="h-fit text-[20px] max-[619px]:text-[16px] border-none outline-none w-full rounded-[8px] overflow-hidden"
            >
                <ContentHeader color={color} mode={mode} setMode={setMode} />
                {mode === "markdown" ? (
                    <textarea
                        ref={textareaRef}
                        placeholder="Type your content here..."
                        required
                        minLength={1}
                        id="content"
                        style={{
                            height: `${contentHeight}px`,
                        }}
                        className="px-[12px] py-[8px] min-h-[500px] border-none outline-none w-full rounded-[8px] overflow-auto"
                        value={content}
                        onChange={(e) => {
                            setContent(e.target.value);
                        }}
                        onMouseUp={handleResize} // for mouse
                        onTouchEnd={handleResize} // for touch devices
                    />
                ) : undefined}
                {mode === "preview" ? (
                    <div
                        style={{
                            height: `${contentHeight}px`,
                        }}
                        className="markdown px-[12px] py-[8px] border-none outline-none w-full rounded-[8px] overflow-auto"
                    >
                        <div>
                            {content.trim() === "" ? (
                                <p className="opacity-50">Nothing to preview</p>
                            ) : (
                                <ReactMarkdown>{content}</ReactMarkdown>
                            )}
                        </div>
                    </div>
                ) : undefined}
            </div>
        </div>
    );
};

export default NoteContent;
