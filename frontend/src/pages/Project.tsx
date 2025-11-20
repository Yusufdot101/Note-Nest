import newResource from "../assets/newResource.svg";
import React, { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import ProjectCard from "../components/ProjectCard";
import ProjectActionsDialoge from "../components/ProjectActionsDialoge";
import { useNavigate, useParams } from "react-router-dom";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";
import { useAuthStore } from "../store/useAuthStore";
import SearchBar from "../components/SearchBar";

const ProjectPage = () => {
    const [project, setProject] = useState<Project>();
    const [notes, setNotes] = useState<Note[]>([]);
    const [showDialoge, setShowDialoge] = useState(false);
    const [color, setColor] = useState("#ffffff");

    const { id } = useParams();
    const userID = useAuthStore((state) => state.userID);
    const [searchValue, setSearchValue] = useState("");
    const [options, setOptions] = useState<Map<string, number | string>>(
        new Map<string, number | string>(id ? [["project_id", id]] : []),
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
        const setupProject = async () => {
            if (id == "") return;
            const project = await fetchProject(+id!);
            if (!project) return;
            setProject(project);
            setColor(project.Color);
        };

        const setupNotes = async () => {
            if (!id) return;
            const notes = await fetchNotes(options);
            if (!notes) return;
            setNotes(notes);
        };
        setupProject();
        setupNotes();
    }, [options, id]);

    return (
        <div className="flex flex-col gap-[12px]">
            <div>
                {project && (
                    <ProjectCard
                        Color={color}
                        SetColor={setColor}
                        project={project}
                        handleMenuClick={
                            project.UserID === userID
                                ? (
                                      e: React.MouseEvent<SVGElement>,
                                      project: Project,
                                  ) => {
                                      e.stopPropagation();
                                      setShowDialoge((prev) => !prev);
                                      setProject(project);
                                  }
                                : undefined
                        }
                    />
                )}
                {showDialoge && project && (
                    <ProjectActionsDialoge
                        color={color}
                        handleClose={() => {
                            setShowDialoge(false);
                        }}
                        project={project}
                    />
                )}
            </div>
            <div>
                <div
                    style={{ border: `1px solid ${color}` }}
                    className="relative text-text cursor-pointer bg-primary p-[12px] flex flex-col gap-[12px] h-fit"
                >
                    <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                        NOTES
                    </h1>

                    <SearchBar
                        placeholder="Search notes"
                        searchValue={searchValue}
                        handleValueChange={(value) => setSearchValue(value)}
                        handleSearch={handleSearch}
                    />

                    {notes.map((note) => (
                        <NoteCard
                            colorEditable={true}
                            key={note.ID}
                            note={note}
                            handleNoteClick={() => navigate(`notes/${note.ID}`)}
                        />
                    ))}

                    {project?.UserID === userID && (
                        <img
                            tabIndex={0}
                            aria-label="new note"
                            onKeyDown={(e) => {
                                if (e.key === "Enter" || e.key === " ") {
                                    navigate(
                                        `/projects/${project?.ID}/notes/new`,
                                    );
                                }
                            }}
                            onClick={() => {
                                navigate(`/projects/${project?.ID}/notes/new`);
                            }}
                            src={newResource}
                            alt="new note"
                            className="sticky bottom-[32px] mt-[-8px] ml-auto cursor-pointer w-[90px] h-[90px] max-[619px]:w-[75px] max-[619px]:h-[75px]"
                        />
                    )}
                </div>
            </div>
        </div>
    );
};

export default ProjectPage;
