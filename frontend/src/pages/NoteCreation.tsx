import { useState } from "react";
import NoteContent from "../components/NoteContent";
import NoteTitle from "../components/NoteTitle";
import SubmitButton from "../components/SubmitButton";
import { useNavigate } from "react-router-dom";
import { newNote } from "../utilities/note";

const NoteCreation = () => {
    const [title, setTitle] = useState("");
    const [color, setColor] = useState("#00FFFF");
    const [content, setContent] = useState("");
    const [visibility, setVisibility] = useState("private");

    const navigate = useNavigate();

    const handleDiscard = () => {
        if (!confirm("are you sure you want to discard")) return;
        const url = new URL(window.location.toString());
        const segments = url.pathname.split("/");
        const projectID = segments.at(-3);
        navigate(`/projects/${projectID}`);
    };

    const handleCreate = async () => {
        if (content === "" || title === "") return;
        const url = new URL(window.location.toString());
        const segments = url.pathname.split("/");
        const projectID = segments.at(-3);
        const success = await newNote(
            +projectID!,
            title,
            content,
            visibility,
            color,
        );

        if (!success) return;
        navigate(`/projects/${projectID}`);
    };

    return (
        <form action={handleCreate} className="flex flex-col gap-[12px]">
            <p className="text-accent text-[32px] max-[619px]:text-[24px] font-semibold text-center">
                CREATE NOTE
            </p>
            <NoteTitle
                title={title}
                setTitle={setTitle}
                color={color}
                setColor={setColor}
            />

            <div className="text-text">
                <label htmlFor="visibility" className="text-[20px]">
                    Visibility
                    <span className="text-[red]">*</span>
                </label>
                <div className="flex items-center gap-[10px] text-[20px]">
                    <div className="flex items-center gap-[8px]">
                        <label htmlFor={"private"}>Private</label>
                        <input
                            type="radio"
                            name="projectVisibility"
                            id="private"
                            value={"private"}
                            className="w-[30px] h-[30px] max-[619px]:w-[20px] accent-accent"
                            checked={visibility === "private"}
                            onChange={(e) => setVisibility(e.target.value)}
                        />
                    </div>
                    <div className="flex items-center gap-[8px]">
                        <label htmlFor={"public"}>Public</label>
                        <input
                            type="radio"
                            name="visibility"
                            id="public"
                            value={"public"}
                            className="w-[30px] h-[30px] max-[619px]:w-[20px] accent-accent"
                            checked={visibility === "public"}
                            onChange={(e) => setVisibility(e.target.value)}
                        />
                    </div>
                </div>
            </div>

            <NoteContent
                content={content}
                setContent={setContent}
                color={color}
            />
            <div className="flex gap-[4px] text-[24px] max-[619px]:text-[16px]">
                <SubmitButton
                    handleSubmit={handleDiscard}
                    type="button"
                    text={"Discard  Note"}
                    textColor={"white"}
                    aria_label={"discard note"}
                    bgColor={"red"}
                />
                <SubmitButton
                    handleSubmit={handleCreate}
                    text={"Create Note"}
                    textColor={"white"}
                    aria_label={"create note"}
                />
            </div>
        </form>
    );
};

export default NoteCreation;
