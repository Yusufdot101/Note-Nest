import { useEffect, useState } from "react";
import SearchBar from "../components/SearchBar";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useNavigate } from "react-router-dom";

const Notes = () => {
    const [searchValue, setSearchValue] = useState("");
    const [notes, setNotes] = useState<Note[]>([]);
    const handleSearch = async () => {};

    const navigate = useNavigate();

    useEffect(() => {
        const setupNotes = async () => {
            const notes = await fetchNotes();
            if (!notes) return;
            setNotes(notes);
        };

        setupNotes();
    }, []);
    return (
        <div className="flex flex-col gap-y-[12px]">
            <SearchBar
                placeholder="Search notes"
                searchValue={searchValue}
                setSearchValue={setSearchValue}
                handleSearch={handleSearch}
            />
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
    );
};

export default Notes;
