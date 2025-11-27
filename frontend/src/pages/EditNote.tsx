import { useEffect, useState } from "react";
import NoteContent from "../components/NoteContent";
import NoteTitle from "../components/NoteTitle";
import SubmitButton from "../components/SubmitButton";
import { useNavigate, useParams } from "react-router-dom";
import {
    editNote,
    editNoteColor,
    fetchNote,
    toggleVisibility,
} from "../utilities/note";
import { useAuthStore } from "../store/useAuthStore";

const EditNote = () => {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const [color, setColor] = useState("#00FFFF");
    const [visibility, setVisibility] = useState("private");

    const [errors, setErrors] = useState<string[]>([]);
    const [showErrors, setShowErrors] = useState(false);

    const navigate = useNavigate();
    const { projectid, noteid } = useParams();

    const handleDiscard = () => {
        if (!confirm("are you sure you want to discard changes? ")) return;
        navigate(`/projects/${projectid}`);
    };

    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        if (!accessToken) return;
        const setupNote = async () => {
            if (!noteid) return;
            const note = await fetchNote(+noteid);
            if (!note) return;
            setTitle(note.Title);
            setContent(note.Content);
            setColor(note.Color);
            setVisibility(note.Visibility);
        };

        setupNote();
    }, [noteid, accessToken]);

    const handleSave = async () => {
        if (content === "" || title === "" || visibility === "") return;
        setShowErrors(false);
        const handleError = (errors: Record<string, string>) => {
            setShowErrors(true);
            const errorMessages = Object.entries(errors).map(
                ([key, value]) => `${key}: ${value}`,
            );
            setErrors(errorMessages);
        };

        if (!noteid) return;

        let success = await editNote(+noteid, title, content, handleError);
        if (!success) return;

        success = await editNoteColor(+noteid, color);
        if (!success) return;

        success = await toggleVisibility(+noteid, visibility);
        if (!success) return;

        navigate(`/projects/${projectid}`);
    };

    return (
        <form
            onSubmit={(e) => {
                e.preventDefault();
                e.stopPropagation();
                handleSave();
            }}
            className="flex flex-col gap-[12px]"
        >
            <p className="text-accent text-[32px] max-[619px]:text-[24px] font-semibold text-center">
                EDIT NOTE
            </p>
            <NoteTitle
                title={title}
                setTitle={setTitle}
                color={color ?? "#ffffff"}
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
                color={color ?? "#ffffff"}
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
                    text={"Discard Changes"}
                    textColor={"white"}
                    aria_label={"discard changes"}
                    bgColor={"red"}
                />
                <SubmitButton
                    handleSubmit={() => {}}
                    text={"Save Changes"}
                    textColor={"white"}
                    aria_label={"save changes"}
                    type="submit"
                />
            </div>
        </form>
    );
};

export default EditNote;
