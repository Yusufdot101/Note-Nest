import { useCallback, useEffect, useState } from "react";
import SearchBar from "../components/SearchBar";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/useAuthStore";
import type { Metadata } from "../utilities/projects";
import PageNumbers from "../components/PageNumbers";

const Notes = () => {
    const [notes, setNotes] = useState<Note[]>([]);
    const pageSize = 10;
    const [metadata, setMetadata] = useState<Metadata>();
    const [options, setOptions] = useState<Map<string, number | string>>(
        new Map<string, number | string>([
            ["title", ""],
            ["page", 1],
            ["page_size", pageSize],
        ]),
    );

    const setupNotes = useCallback(
        async (currentOptions: Map<string, string | number>) => {
            const result = await fetchNotes(currentOptions);
            if (!result) return;
            const { notes, metadata } = result;
            setNotes(notes);
            setMetadata(metadata);
        },
        [],
    );

    const handleSearch = async () => {
        setupNotes(options);
    };

    const navigate = useNavigate();

    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        setupNotes(options);
    }, [accessToken, setupNotes, options]);

    const updateOptions = (key: string, value: string | number) => {
        setOptions((prev) => {
            const newOptions = new Map<string, string | number>([
                ...prev,
                [key, value],
            ]);
            return newOptions;
        });
    };

    return (
        <div className="flex flex-col relative text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white">
            <div className="flex flex-col gap-y-[12px]">
                <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                    NOTES
                </h1>

                <SearchBar
                    handleOptionsChange={updateOptions}
                    options={options}
                    searchPlaceholder="Search notes"
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

                {metadata ? (
                    <PageNumbers
                        options={options}
                        handleOptionsChange={updateOptions}
                        metadata={metadata}
                    />
                ) : undefined}
            </div>
        </div>
    );
};

export default Notes;
