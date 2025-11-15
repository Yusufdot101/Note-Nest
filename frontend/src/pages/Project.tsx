import newNote from "../assets/newNoteButton.svg";
import React, { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import ProjectCard from "../components/ProjectCard";
import ProjectActionsDialoge from "../components/ProjectActionsDialoge";
import { useNavigate } from "react-router-dom";

const ProjectPage = () => {
    const [project, setProject] = useState<Project>();
    const [showDialoge, setShowDialoge] = useState(false);

    const navigate = useNavigate();

    useEffect(() => {
        const setupProject = async () => {
            const url = new URL(window.location.toString());
            const segments = url.pathname.split("/").filter(Boolean);
            const projectID = segments.at(-1);
            if (projectID == "") return;
            const project = await fetchProject(+projectID!);
            if (!project) return;
            setProject(project);
        };
        setupProject();
    }, []);

    return (
        <div className="flex flex-col gap-[12px]">
            <div>
                {project && (
                    <ProjectCard
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
                        handleClose={() => {
                            setShowDialoge(false);
                        }}
                        project={project!}
                    />
                )}
            </div>
            <div>
                <div
                    style={{ border: `1px solid ${project?.Color}` }}
                    className="relative text-text cursor-pointer bg-primary p-[12px] flex flex-col gap-[12px] h-fit"
                >
                    <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                        NOTES
                    </h1>
                    {/* TODO: list the notes below */}
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
