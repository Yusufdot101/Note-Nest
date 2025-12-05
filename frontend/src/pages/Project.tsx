import newResource from "../assets/newResource.svg";
import React, { useCallback, useEffect, useState } from "react";
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
import type { Metadata } from "../utilities/projects";
import PageNumbers from "../components/PageNumbers";

const ProjectPage = () => {
    const [project, setProject] = useState<Project>();
    const [notes, setNotes] = useState<Note[]>([]);
    const pageSize = 10;
    const [metadata, setMetadata] = useState<Metadata>();

    const [showDialoge, setShowDialoge] = useState(false);
    const [color, setColor] = useState("#ffffff");

    const { id } = useParams();
    const userID = useAuthStore((state) => state.userID);
    const [options, setOptions] = useState<Map<string, number | string>>(
        new Map<string, number | string>([
            ["project_id", id ?? -1],
            ["title", ""],
            ["page", 1],
            ["page_size", pageSize],
        ]),
    );

    const setupNotes = useCallback(
        async (currentOptions: Map<string, string | number>) => {
            if (!id) return;
            const result = await fetchNotes(currentOptions);
            if (!result) return;
            const { notes, metadata } = result;
            setNotes(notes);
            setMetadata(metadata);
        },
        [id],
    );

    const handleSearch = async () => {
        setupNotes(options);
    };

    const navigate = useNavigate();

    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        if (!accessToken) return;
        const setupProject = async () => {
            if (id == "") return;
            const project = await fetchProject(+id!);
            if (!project) return;
            setProject(project);
            setColor(project.Color);
        };
        setupProject();
        setupNotes(options);
    }, [id, accessToken, setupNotes, options]);

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
        <div className="flex flex-col gap-[12px]">
            <div>
                {project && (
                    <ProjectCard
                        project={project}
                        handleMenuClick={
                            project.UserID === userID
                                ? (
                                      e:
                                          | React.MouseEvent<SVGElement>
                                          | React.KeyboardEvent<SVGElement>,
                                      project: Project,
                                  ) => {
                                      e.stopPropagation();
                                      setShowDialoge((prev) => !prev);
                                      setProject(project);
                                  }
                                : undefined
                        }
                        colorEditable={(userID ?? -1) === project.UserID}
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
                    className="relative text-text cursor-pointer bg-primary p-[12px] flex flex-col gap-[12px] h-fit rounded-[8px]"
                >
                    <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                        NOTES
                    </h1>

                    <SearchBar
                        handleOptionsChange={updateOptions}
                        options={options}
                        searchPlaceholder="Search notes"
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

                    {metadata ? (
                        <PageNumbers
                            options={options}
                            handleOptionsChange={updateOptions}
                            metadata={metadata}
                        />
                    ) : undefined}
                </div>
            </div>
        </div>
    );
};

export default ProjectPage;
