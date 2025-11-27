import { useEffect, useState } from "react";
import SearchBar from "../components/SearchBar";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/useAuthStore";

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

    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        const setupNotes = async () => {
            const notes = await fetchNotes(options);
            if (!notes) return;
            setNotes(notes);
        };

        setupNotes();
    }, [options, accessToken]);
    return (
        <div className="flex flex-col relative text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white">
            <div className="flex flex-col gap-y-[12px]">
                <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                    NOTES
                </h1>
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
        </div>
    );
};

export default Notes;
