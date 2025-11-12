import React, { useState } from "react";

export interface Project {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    UserID: number;
    Name: string;
    Description: string;
    Visibility: string;
    Color: string;
    EntriesCount: number;
    LikesCount: number;
    CommentsCount: number;
}

interface ProjectCardProps {
    project: Project;
    handleMenuClick?: (
        e: React.MouseEvent<SVGElement>,
        project: Project,
    ) => void;
    handleProjectClick?: (
        e: React.MouseEvent<HTMLDivElement>,
        projectID: number,
    ) => void;
}
const ProjectCard = ({
    project,
    handleMenuClick,
    handleProjectClick,
}: ProjectCardProps) => {
    const [color, setColor] = useState(project.Color ? project.Color : "white");
    return (
        <div
            style={{ boxShadow: `0px 0px 4px 1px ${color}` }}
            className="text-text cursor-pointer bg-primary p-[12px] flex flex-col gap-[12px] h-[200px]"
            onClick={(e) =>
                handleProjectClick
                    ? handleProjectClick(e, project.ID)
                    : () => {}
            }
        >
            <div className="flex items-center justify-between gap-[4px]">
                <div className="flex items-center gap-[8px]">
                    <div
                        className="relative min-w-[40px] h-[30px] rounded-lg"
                        style={{ backgroundColor: color }}
                        onClick={(e) => {
                            e.stopPropagation();
                        }}
                    >
                        <input
                            className="inline-block absolute cursor-pointer w-full h-full opacity-0"
                            type="color"
                            value={color}
                            onChange={(e) => {
                                e.stopPropagation();
                                setColor(e.target.value);
                            }}
                            onInput={() => {}}
                        />
                    </div>
                    <span
                        style={{ color: color }}
                        className="text-[28px] max-[629px]:text-[20px] font-bold w-full line-clamp-1 underline"
                    >
                        {project.Name}
                    </span>
                </div>
                <div className="flex gap-[12px] items-center">
                    <span className="text-right font-bold">
                        [{project.Visibility}]
                    </span>
                    <span>
                        <svg
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

            <div className="flex flex-col gap-[12px] font-bold">
                <p
                    className={`line-clamp-2 wrap-break-word ${!project.Description ? "opacity-50" : ""}`}
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
