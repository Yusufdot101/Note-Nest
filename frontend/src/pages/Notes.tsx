import { useEffect, useState } from "react";
import SearchBar from "../components/SearchBar";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useNavigate } from "react-router-dom";

const Notes = () => {
    const [searchValue, setSearchValue] = useState("");
    const [notes, setNotes] = useState<Note[]>([]);
    const [options, setOptions] = useState<Map<string, number | string>>(
        new Map<string, number | string>(),
    );

    const handleSearch = async () => {
        if (!searchValue.trim()) {
            const newOptions = new Map(options);
            newOptions.delete("title");
            setOptions(newOptions);
            return;
        }

        setOptions(
            (prev) =>
                new Map<string, string | number>([
                    ...prev,
                    ["title", searchValue],
                ]),
        );
    };

    const navigate = useNavigate();

    useEffect(() => {
        const setupNotes = async () => {
            const notes = await fetchNotes(options);
            if (!notes) return;
            setNotes(notes);
        };

        setupNotes();
    }, [options]);
    return (
        <div className="flex flex-col gap-y-[12px]">
            <SearchBar
                placeholder="Search notes"
                searchValue={searchValue}
                handleValueChange={(value) => setSearchValue(value)}
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
