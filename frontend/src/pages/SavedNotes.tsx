import { useEffect, useState } from "react";
import { fetchSavedNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useNavigate } from "react-router-dom";

const SavedNotes = () => {
    const [notes, setNotes] = useState<Note[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        const setupNotes = async () => {
            const notes = await fetchSavedNotes();
            if (!notes) return;
            setNotes(notes);
        };

        setupNotes();
    }, []);
    return (
        <div className="flex flex-col relative text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white">
            <div className="flex flex-col gap-y-[12px]">
                <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                    SAVED NOTES
                </h1>
                <div className="flex flex-col gap-[8px]">
                    {notes.map((note) => (
                        <NoteCard
                            key={note.ID}
                            colorEditable={false}
                            note={note}
                            handleNoteClick={() =>
                                navigate(
                                    `/projects/${note.ProjectID}/notes/${note.ID}`,
                                )
                            }
                        />
                    ))}
                </div>
            </div>
        </div>
    );
};

export default SavedNotes;
