import React, { useState } from "react";
import ColorPicker from "./ColorPicker";
import { editProjectColor } from "../utilities/project";

export interface Project {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    UserID: number;
    Title: string;
    Description: string;
    Visibility: string;
    Color: string;
    EntriesCount: number;
    LikesCount: number;
    CommentsCount: number;
}

interface ProjectCardProps {
    color: string;
    setColor?: React.Dispatch<React.SetStateAction<string>>;
    project: Project;
    handleMenuClick?: (
        e: React.MouseEvent<SVGElement> | React.KeyboardEvent<SVGElement>,
        project: Project,
    ) => void;
    handleProjectClick?: (
        e:
            | React.MouseEvent<HTMLDivElement>
            | React.KeyboardEvent<HTMLDivElement>,
        projectID: number,
    ) => void;
    colorEditable?: boolean;
}
const ProjectCard = ({
    project,
    handleMenuClick,
    handleProjectClick,
    colorEditable = false,
}: ProjectCardProps) => {
    const [color, setColor] = useState(project.Color ?? "#ffffff");

    return (
        <div
            tabIndex={0}
            aria-label="project card"
            role="group"
            onKeyDown={(e) => {
                if (
                    (e.key === "Enter" || e.key === " ") &&
                    handleProjectClick
                ) {
                    handleProjectClick(e, project.ID);
                }
            }}
            style={{ border: `1px solid ${color ?? "#ffffff"}` }}
            className="text-text cursor-pointer bg-primary p-[12px] rounded-[8px] flex flex-col justify-between gap-[12px] h-full"
            onClick={(e) =>
                handleProjectClick
                    ? handleProjectClick(e, project.ID)
                    : () => {}
            }
        >
            <div className="flex items-center justify-between gap-[4px]">
                <div className="flex items-center gap-[8px]">
                    <div className="h-[35px]">
                        <ColorPicker
                            color={color ?? "#ffffff"}
                            handleChange={
                                colorEditable
                                    ? (value) => {
                                          editProjectColor(project.ID, value);
                                          setColor(value);
                                      }
                                    : undefined
                            }
                        />
                    </div>
                    <span
                        style={{ color: color }}
                        className="text-[28px] max-[629px]:text-[20px] font-bold w-full line-clamp-1 underline"
                    >
                        {project.Title}
                    </span>
                </div>
                <div className="flex gap-[12px] items-center">
                    <span className="text-right font-bold">
                        {project.Visibility === "public" ? (
                            <svg
                                className="w-[32px] h-[32px]"
                                viewBox="0 0 24 24"
                                fill="none"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                                <g
                                    id="SVGRepo_tracerCarrier"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                ></g>
                                <g id="SVGRepo_iconCarrier">
                                    {" "}
                                    <path
                                        d="M16.584 6C15.8124 4.2341 14.0503 3 12 3C9.23858 3 7 5.23858 7 8V10.0288M12 14.5V16.5M7 10.0288C7.47142 10 8.05259 10 8.8 10H15.2C16.8802 10 17.7202 10 18.362 10.327C18.9265 10.6146 19.3854 11.0735 19.673 11.638C20 12.2798 20 13.1198 20 14.8V16.2C20 17.8802 20 18.7202 19.673 19.362C19.3854 19.9265 18.9265 20.3854 18.362 20.673C17.7202 21 16.8802 21 15.2 21H8.8C7.11984 21 6.27976 21 5.63803 20.673C5.07354 20.3854 4.6146 19.9265 4.32698 19.362C4 18.7202 4 17.8802 4 16.2V14.8C4 13.1198 4 12.2798 4.32698 11.638C4.6146 11.0735 5.07354 10.6146 5.63803 10.327C5.99429 10.1455 6.41168 10.0647 7 10.0288Z"
                                        stroke="currentColor"
                                        strokeWidth="2"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    ></path>{" "}
                                </g>
                            </svg>
                        ) : (
                            <svg
                                className="w-[32px] h-[32px]"
                                viewBox="0 0 24 24"
                                fill="none"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                                <g
                                    id="SVGRepo_tracerCarrier"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                ></g>
                                <g id="SVGRepo_iconCarrier">
                                    {" "}
                                    <path
                                        d="M12 14.5V16.5M7 10.0288C7.47142 10 8.05259 10 8.8 10H15.2C15.9474 10 16.5286 10 17 10.0288M7 10.0288C6.41168 10.0647 5.99429 10.1455 5.63803 10.327C5.07354 10.6146 4.6146 11.0735 4.32698 11.638C4 12.2798 4 13.1198 4 14.8V16.2C4 17.8802 4 18.7202 4.32698 19.362C4.6146 19.9265 5.07354 20.3854 5.63803 20.673C6.27976 21 7.11984 21 8.8 21H15.2C16.8802 21 17.7202 21 18.362 20.673C18.9265 20.3854 19.3854 19.9265 19.673 19.362C20 18.7202 20 17.8802 20 16.2V14.8C20 13.1198 20 12.2798 19.673 11.638C19.3854 11.0735 18.9265 10.6146 18.362 10.327C18.0057 10.1455 17.5883 10.0647 17 10.0288M7 10.0288V8C7 5.23858 9.23858 3 12 3C14.7614 3 17 5.23858 17 8V10.0288"
                                        stroke="currentColor"
                                        strokeWidth="2"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    ></path>{" "}
                                </g>
                            </svg>
                        )}
                    </span>
                    <span>
                        <svg
                            role="button"
                            aria-label="open project actions menu"
                            tabIndex={0}
                            onKeyDown={(e) => {
                                if (
                                    (e.key === "Enter" || e.key === " ") &&
                                    handleMenuClick
                                ) {
                                    handleMenuClick(e, project);
                                }
                            }}
                            fill="currentColor"
                            version="1.1"
                            id="Icons"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 32 32"
                            className={`${handleMenuClick ? "" : "hidden"} w-[30px] h-[30px] hover:text-accent active:text-text duration-300`}
                            onClick={(e) => {
                                handleMenuClick!(e, project);
                            }}
                        >
                            <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                            <g
                                id="SVGRepo_tracerCarrier"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            ></g>
                            <g id="SVGRepo_iconCarrier">
                                {" "}
                                <g>
                                    {" "}
                                    <path d="M16,10c1.7,0,3-1.3,3-3s-1.3-3-3-3s-3,1.3-3,3S14.3,10,16,10z"></path>{" "}
                                    <path d="M16,13c-1.7,0-3,1.3-3,3s1.3,3,3,3s3-1.3,3-3S17.7,13,16,13z"></path>{" "}
                                    <path d="M16,22c-1.7,0-3,1.3-3,3s1.3,3,3,3s3-1.3,3-3S17.7,22,16,22z"></path>{" "}
                                </g>{" "}
                            </g>
                        </svg>
                    </span>
                </div>
            </div>

            <div className="h-full flex flex-col gap-[12px] font-bold">
                <p
                    className={`line-clamp-3 wrap-break-word h-full ${!project.Description ? "opacity-50" : ""}`}
                >
                    {project.Description
                        ? project.Description
                        : "No Description"}
                </p>
                <div className="flex flex-col gap-[4px]">
                    <div className="flex gap-[12px] font-semibold">
                        <span>Entries: {project.EntriesCount}</span>
                        <span>Likes: {project.LikesCount}</span>
                        <span>Comments: {project.CommentsCount}</span>
                    </div>
                    <div className="font-semibold">
                        <p>
                            Created:{" "}
                            {new Date(project.CreatedAt).toDateString()}
                        </p>
                        <p
                            className={`${project.UpdatedAt == undefined ? "opacity-50" : ""}`}
                        >
                            Updated:{" "}
                            {project.UpdatedAt
                                ? new Date(project.UpdatedAt).toDateString()
                                : "Not Updated"}
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ProjectCard;
