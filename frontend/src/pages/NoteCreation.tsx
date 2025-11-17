import { useState } from "react";
import NoteContent from "../components/NoteContent";
import NoteTitle from "../components/NoteTitle";
import SubmitButton from "../components/SubmitButton";
import { useNavigate, useParams } from "react-router-dom";
import { newNote } from "../utilities/note";

const NoteCreation = () => {
    const [title, setTitle] = useState("");
    const [color, setColor] = useState("#00FFFF");
    const [content, setContent] = useState("");
    const [visibility, setVisibility] = useState("private");

    const [errors, setErrors] = useState<string[]>([]);
    const [showErrors, setShowErrors] = useState(false);

    const navigate = useNavigate();
    const { projectid } = useParams();

    const handleDiscard = () => {
        if (!confirm("are you sure you want to discard")) return;
        navigate(`/projects/${projectid}`);
    };

    const handleCreate = async () => {
        if (content === "" || title === "" || !projectid) return;
        setShowErrors(false);
        const handleError = (errors: Record<string, string>) => {
            setShowErrors(true);
            const errorMessages = Object.entries(errors).map(
                ([key, value]) => `${key}: ${value}`,
            );
            setErrors(errorMessages);
        };

        const success = await newNote(
            +projectid,
            title,
            content,
            visibility,
            color,
            handleError,
        );

        if (!success) return;
        navigate(`/projects/${projectid}`);
    };

    return (
        <form
            onSubmit={(e) => {
                e.preventDefault();
                e.stopPropagation();
                handleCreate();
            }}
            className="flex flex-col gap-[12px]"
        >
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
                            name="visibility"
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

            <div
                className={`w-full text-text text-center text-[24px] max-[619px]:text-[16px] p-[12px] rounded-[8px] bg-red-500 mx-auto ${!showErrors ? "hidden" : ""}`}
            >
                {errors.map((error) => (
                    <p key={error}>{error}</p>
                ))}
            </div>

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
                    handleSubmit={() => {}}
                    text={"Create Note"}
                    textColor={"white"}
                    aria_label={"create note"}
                    type="submit"
                />
            </div>
        </form>
    );
};

export default NoteCreation;
