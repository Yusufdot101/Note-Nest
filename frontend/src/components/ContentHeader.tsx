interface ContentHeaderProps {
    mode: string;
    color: string;
    setMode: React.Dispatch<React.SetStateAction<string>>;
}
const ContentHeader = ({ mode, color, setMode }: ContentHeaderProps) => {
    const ACTIVE_BG_COLOR = "#38a100";
    return (
        <div
            style={{ borderBottom: `1px solid ${color}` }}
            className="flex h-[50px] flex text-text w-full"
        >
            {/* TODO: Add Rich Text mode support */}
            {/* <button */}
            {/*     type="button" */}
            {/*     style={{ */}
            {/*         borderRight: `${mode === "richText" ? `1px solid ${color}` : ""}`, */}
            {/*         borderBottom: `${mode === "richText" ? "none" : ""}`, */}
            {/*         marginBottom: ` ${mode === "richText" ? "-1px" : "0"}`, */}
            {/*         background: `${mode === "richText" ? "#38a100" : ""}`, */}
            {/*     }} */}
            {/*     className="mb-[-10px] z-1 rounded-t-[4px] px-[8px] w-full cursor-pointer" */}
            {/*     onClick={() => setMode("richText")} */}
            {/* > */}
            {/*     Rich Text */}
            {/* </button> */}
            <button
                type="button"
                style={{
                    borderRight: `${mode === "markdown" ? `1px solid ${color}` : ""}`,
                    borderLeft: `${mode === "markdown" ? `1px solid ${color}` : ""}`,
                    borderBottom: `${mode === "markdown" ? "none" : ""}`,
                    marginBottom: ` ${mode === "markdown" ? "-1px" : "0"}`,
                    background: `${mode === "markdown" ? ACTIVE_BG_COLOR : ""}`,
                }}
                className="mb-[-10px] z-1 px-[8px] rounded-t-[4px] w-full cursor-pointer"
                onClick={() => setMode("markdown")}
            >
                Markdown
            </button>
            <button
                type="button"
                style={{
                    borderLeft: `${mode === "preview" ? `1px solid ${color}` : ""}`,
                    borderBottom: `${mode === "preview" ? "none" : ""}`,
                    marginBottom: ` ${mode === "preview" ? "-1px" : "0"}`,
                    background: `${mode === "preview" ? ACTIVE_BG_COLOR : ""}`,
                }}
                className="mb-[-10px] z-1 px-[8px] rounded-t-[4px] w-full cursor-pointer"
                onClick={() => setMode("preview")}
            >
                Preview
            </button>
        </div>
    );
};

export default ContentHeader;
