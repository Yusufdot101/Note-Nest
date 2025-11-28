import { useEffect, useRef } from "react";
import type { Project } from "./ProjectCard";
import { useNavigate } from "react-router-dom";
import { deleteProject, toggleProjectVisibility } from "../utilities/project";

const ProjectActionsDialoge = ({
    color,
    handleClose,
    project,
}: {
    color: string;
    handleClose: () => void;
    project: Project;
}) => {
    const ref = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (!ref.current?.contains(e.target as Node)) {
                handleClose();
            }
        };
        document.addEventListener("click", handleClick);
        return () => document.removeEventListener("click", handleClick);
    }, [handleClose]);

    const navigate = useNavigate();

    return (
        <div
            ref={ref}
            className="fixed text-[28px] z-2 max-[629px]:text-[20px] top-0 mt-[30vh] left-0 right-0 m-auto h-fit w-[90vw] max-w-[700px] bg-primary border-[1px] border-solid border-[#ffffff] rounded-[8px] text-text p-[12px]"
        >
            <p
                style={{ color: color }}
                className="font-bold text-[32px] max-[629px]:text-[24px] p-[8px] text-center"
            >
                {project.Title}
            </p>
            <ul className="flex flex-col gap-[8px]">
                <li className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer">
                    Add New Note
                </li>

                <li
                    onClick={async () => {
                        const newVisibility =
                            project.Visibility === "private"
                                ? "public"
                                : "private";
                        if (
                            !confirm(
                                `Are you sure you want to make this project ${newVisibility}? ${newVisibility === "private" ? "All notes in this project will become private" : ""} `,
                            )
                        ) {
                            return;
                        }
                        const success = await toggleProjectVisibility(
                            project.ID,
                            newVisibility,
                        );
                        if (!success) return;
                        navigate(0);
                    }}
                    className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Make Project{" "}
                    {project.Visibility === "private" ? "public" : "private"}
                </li>

                <li
                    onClick={() => {
                        navigate(`/projects/${project.ID}/edit`);
                    }}
                    className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Edit Project
                </li>

                <li
                    onClick={async () => {
                        const success = await deleteProject(project.ID);
                        if (!success) return;
                        navigate("/projects");
                    }}
                    className="bg-[#FF0000] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Delete Project
                </li>
            </ul>
        </div>
    );
};

export default ProjectActionsDialoge;
