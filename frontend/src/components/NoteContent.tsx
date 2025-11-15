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
            // const { offsetHeight } = textareaRef.current;
            textareaRef.current.style.height = "auto";
            const newHeight = Math.max(500, textareaRef.current.scrollHeight);
            setContentHeight(newHeight);
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
                            handleResize();
                        }}
                    />
                ) : undefined}
                {mode === "preview" ? (
                    <div className="markdown min-h-[500px] px-[12px] border-none outline-none w-full rounded-[8px] overflow-auto">
                        {content.trim() === "" ? (
                            <p className="opacity-50">Nothing to preview</p>
                        ) : (
                            <ReactMarkdown>{content}</ReactMarkdown>
                        )}
                    </div>
                ) : undefined}
            </div>
        </div>
    );
};

export default NoteContent;
