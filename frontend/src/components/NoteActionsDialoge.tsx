import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import type { Note } from "./NoteCard";
import { deleteNote, toggleVisibility } from "../utilities/note";

const NoteActionsDialoge = ({
    color,
    handleClose,
    note,
}: {
    color: string;
    handleClose: () => void;
    note: Note;
}) => {
    const ref = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (!ref.current?.contains(e.target as Node)) {
                handleClose();
            }
        };
        document.addEventListener("click", handleClick);
        return () => document.removeEventListener("click", handleClick);
    }, [handleClose]);

    const navigate = useNavigate();

    return (
        <div
            ref={ref}
            className="fixed text-[28px] z-2 max-[629px]:text-[20px] top-0 mt-[30vh] left-0 right-0 m-auto h-fit w-[90vw] max-w-[700px] bg-primary border-[1px] border-solid border-[#ffffff] rounded-[8px] text-text p-[12px]"
        >
            <p
                style={{ color: color }}
                className="font-bold text-[32px] max-[629px]:text-[24px] p-[8px] text-center"
            >
                {note.Title}
            </p>
            <ul className="flex flex-col gap-[8px]">
                <li
                    onClick={() => {
                        navigate(
                            `/projects/${note.ProjectID}/notes/${note.ID}/edit`,
                        );
                    }}
                    className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Edit Note
                </li>
                <li
                    onClick={async () => {
                        const newVisibility =
                            note.Visibility === "private"
                                ? "public"
                                : "private";
                        if (
                            !confirm(
                                `Are you sure you want to make this note ${newVisibility}? `,
                            )
                        ) {
                            return;
                        }
                        const success = await toggleVisibility(
                            note.ID,
                            newVisibility,
                        );
                        if (!success) return;
                        navigate(0);
                    }}
                    className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Make Note{" "}
                    {note.Visibility === "private" ? "public" : "private"}
                </li>
                <li
                    onClick={async () => {
                        if (
                            !confirm(
                                "Are you sure you want to delete this note? ",
                            )
                        ) {
                            return;
                        }

                        const success = await deleteNote(note.ID);
                        if (!success) return;
                        navigate(`/projects/${note.ProjectID}`);
                    }}
                    className="bg-[#FF0000] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Delete Note
                </li>
            </ul>
        </div>
    );
};

export default NoteActionsDialoge;
