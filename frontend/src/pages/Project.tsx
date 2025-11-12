import React, { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import ProjectCard from "../components/ProjectCard";
import ProjectActionsDialoge from "../components/ProjectActionsDialoge";

const ProjectPage = () => {
    const [project, setProject] = useState<Project>();
    const [showDialoge, setShowDialoge] = useState(false);

    useEffect(() => {
        const setupProject = async () => {
            const project = await fetchProject();
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
                    style={{ boxShadow: `0px 0px 4px 1px white` }}
                    className="text-text cursor-pointer bg-primary p-[12px] flex flex-col gap-[12px] h-fit"
                >
                    <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                        NOTES
                    </h1>
                    {/* TODO: list the notes below */}
                </div>
            </div>
        </div>
    );
};

export default ProjectPage;
