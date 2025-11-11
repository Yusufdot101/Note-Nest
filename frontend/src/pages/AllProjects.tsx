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
        e: React.MouseEvent<HTMLDivElement>,
        projectID: number,
    ) => {
        e.stopPropagation();
        navigate(`/projects/${projectID}`);
    };

    return (
        <div
            className={`py-[12px] items-center text-text grid gap-[16px] ${projects.length > 1 ? "grid-cols-[repeat(auto-fit,minmax(300px,1fr))]" : "grid-cols-[repeat(auto-fit,minmax(300px,400px))]"}`}
        >
            {projects.map((project) => (
                <ProjectCard
                    key={project.ID}
                    project={project}
                    handleProjectClick={handleProjectClick}
                />
            ))}
        </div>
    );
};

export default AllProjects;
