import newResource from "../assets/newResource.svg";
import { useEffect, useState } from "react";
import type { Project } from "../components/ProjectCard";
import ProjectCard from "../components/ProjectCard";
import { fetchProjects } from "../utilities/projects";
import { useNavigate } from "react-router-dom";
import SearchBar from "../components/SearchBar";
import { useAuthStore } from "../store/useAuthStore";

const AllProjects = () => {
    const [projects, setProjects] = useState<Project[]>([]);
    const [searchValue, setSearchValue] = useState("");

    const params = new URLSearchParams(window.location.search);
    const user = params.get("user");
    const [options, setOptions] = useState<Map<string, number | string>>(
        new Map<string, number | string>([["user_id", user ?? -1]]),
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

    // accessToken is null at first so fetch resources doesnt work as expected, this makes it reload when accessToken changes from null
    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        const setupProjects = async () => {
            const projects = await fetchProjects(options);
            setProjects(projects);
        };
        setupProjects();
    }, [options, accessToken]);

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
            <SearchBar
                placeholder="Search projects"
                searchValue={searchValue}
                handleValueChange={(value) => setSearchValue(value)}
                handleSearch={handleSearch}
            />
            <div className="py-[12px] items-center text-text grid gap-[16px]  grid-cols-[repeat(auto-fit,minmax(400px,1fr))] max-[619px]:grid-cols-[repeat(auto-fit,minmax(284px,1fr))]">
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
                className="sticky ml-auto bottom-[32px] mt-[-20px] cursor-pointer w-[90px] h-[90px] max-[619px]:w-[75px] max-[619px]:h-[75px]"
            />
        </div>
    );
};

export default AllProjects;
