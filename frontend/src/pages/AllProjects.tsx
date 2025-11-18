import newResource from "../assets/newResource.svg";
import { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import ProjectCard from "../components/ProjectCard";
import { fetchProjects } from "../utilities/projects";
import { useNavigate } from "react-router-dom";

const AllProjects = () => {
    const [projects, setProjects] = useState<Project[]>([]);

    useEffect(() => {
        const setupProjects = async () => {
            const projects = await fetchProjects();
            setProjects(projects);
        };
        setupProjects();
    }, []);

    const navigate = useNavigate();

    const handleProjectClick = (
        e:
            | React.MouseEvent<HTMLDivElement>
            | React.KeyboardEvent<HTMLDivElement>,
        projectID: number,
    ) => {
        e.stopPropagation();
        navigate(`/projects/${projectID}`);
    };

    return (
        <div className="flex flex-col relative text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white">
            <h1 className="text-text font-bold text-[32px] max-[629px]:text-[24px] text-center">
                PROJECTS
            </h1>
            <div
                className={`py-[12px] items-center text-text grid gap-[16px] ${projects.length > 1 ? " grid-cols-[repeat(auto-fit,minmax(400px,1fr))] max-[619px]:grid-cols-[repeat(auto-fit,minmax(284px,1fr))]" : "grid-cols-[repeat(auto-fit,minmax(284px,700px))]"}`}
            >
                {projects.map((project) => {
                    return (
                        <ProjectCard
                            key={project.ID}
                            Color={project.Color}
                            project={project}
                            handleProjectClick={handleProjectClick}
                        />
                    );
                })}
            </div>
            <img
                tabIndex={0}
                aria-label="new project"
                onKeyDown={(e) => {
                    if (e.key === "Enter" || e.key === " ") {
                        navigate("/projects/new");
                    }
                }}
                onClick={() => {
                    navigate("/projects/new");
                }}
                src={newResource}
                alt="new project"
                className="sticky ml-auto bottom-[32px] mt-[-64px] cursor-pointer w-[90px] h-[90px] max-[619px]:w-[75px] max-[619px]:h-[75px]"
            />
        </div>
    );
};

export default AllProjects;
