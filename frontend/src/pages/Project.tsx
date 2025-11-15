import newNote from "../assets/newNoteButton.svg";
import React, { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import ProjectCard from "../components/ProjectCard";
import ProjectActionsDialoge from "../components/ProjectActionsDialoge";
import { useNavigate, useParams } from "react-router-dom";
import { fetchNotes } from "../utilities/note";
import type { Note } from "../components/NoteCard";
import NoteCard from "../components/NoteCard";

const ProjectPage = () => {
    const [project, setProject] = useState<Project>();
    const [notes, setNotes] = useState<Note[]>([]);
    const [showDialoge, setShowDialoge] = useState(false);
    const [color, setColor] = useState("#ffffff");

    const navigate = useNavigate();
    const { id } = useParams();

    useEffect(() => {
        const setupProject = async () => {
            if (id == "") return;
            const project = await fetchProject(+id!);
            if (!project) return;
            setProject(project);
            setColor(project.Color);
        };

        const setupNotes = async () => {
            if (id == "") return;
            const notes = await fetchNotes(+id!);
            if (!notes) return;
            setNotes(notes);
        };
        setupProject();
        setupNotes();
    }, [id]);

    return (
        <div className="flex flex-col gap-[12px]">
            <div>
                {project && (
                    <ProjectCard
                        Color={color}
                        SetColor={setColor}
                        project={project}
                        handleMenuClick={(
                            e: React.MouseEvent<SVGElement>,
                            project: Project,
                        ) => {
                            e.stopPropagation();
                            setShowDialoge((prev) => !prev);
                            setProject(project);
                        }}
                    />
                )}
                {showDialoge && (
                    <ProjectActionsDialoge
                        color={color}
                        handleClose={() => {
                            setShowDialoge(false);
                        }}
                        project={project!}
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
                    {/* TODO: list the notes below */}
                    {notes.map((note) => (
                        <NoteCard
                            key={note.ID}
                            note={note}
                            handleNoteClick={() => {}}
                        />
                    ))}

                    <img
                        onClick={() => {
                            navigate(`/projects/${project?.ID}/notes/new`);
                        }}
                        src={newNote}
                        alt="logo"
                        className="absolute right-0 bottom-[-40px] cursor-pointer w-[90px] h-[90px] max-[619px]:w-[75px] max-[619px]:h-[75px]"
                    />
                </div>
            </div>
        </div>
    );
};

export default ProjectPage;
