import { useEffect, useRef } from "react";
import type { Project } from "./ProjectCard";

const ProjectActionsDialoge = ({
    handleClose,
    project,
}: {
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

    return (
        <div
            ref={ref}
            className="fixed text-[28px] max-[629px]:text-[20px] top-0 mt-[30vh] left-0 right-0 m-auto h-fit w-[90vw] max-w-[700px] bg-primary shadow-[0px_0px_4px_1px_white] text-text p-[12px]"
        >
            <p
                style={{ color: project.Color }}
                className="font-bold text-[32px] max-[629px]:text-[24px] p-[8px] text-center"
            >
                {project.Name}
            </p>
            <ul className="flex flex-col gap-[8px]">
                <li className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer">
                    Edit Prject
                </li>
                <li className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer">
                    Add New Note
                </li>
                <li className="bg-[#FF0000] p-[8px] hover:opacity-80 duration-300 cursor-pointer">
                    Delete Project
                </li>
            </ul>
        </div>
    );
};

export default ProjectActionsDialoge;
