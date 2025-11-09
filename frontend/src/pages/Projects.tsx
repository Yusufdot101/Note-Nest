import { useEffect, useState } from "react";
import type { ProjectCardProps } from "../components/ProjectCard";
import ProjectCard from "../components/ProjectCard";
import { fetchProjects } from "../utilities/projects";

const Projects = () => {
    const [projects, setProjects] = useState<ProjectCardProps[]>([]);

    useEffect(() => {
        const setupProjects = async () => {
            const projects = await fetchProjects();
            setProjects(projects);
        };
        setupProjects();
    }, []);

    return (
        <div
            className={`py-[12px] items-center text-text grid gap-[16px] ${projects.length > 1 ? "grid-cols-[repeat(auto-fit,minmax(300px,1fr))]" : "grid-cols-[repeat(auto-fit,minmax(300px,400px))]"}`}
        >
            {projects.map((project) => (
                <ProjectCard key={project.ID} {...project} />
            ))}
        </div>
    );
};

export default Projects;
