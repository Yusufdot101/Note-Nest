import { useState } from "react";

export interface ProjectCardProps {
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
const ProjectCard = (project: ProjectCardProps) => {
    const [color, setColor] = useState(project.Color ? project.Color : "white");
    return (
        <div
            style={{ boxShadow: `0px 0px 4px 1px ${color}` }}
            className="cursor-pointer bg-primary p-[12px] flex flex-col gap-[8px] h-[200px]"
        >
            <div className="flex items-center justify-between gap-[4px]">
                <div className="flex items-center gap-[8px]">
                    <div
                        className="relative min-w-[40px] h-[30px] rounded-lg"
                        style={{ backgroundColor: color }}
                    >
                        <input
                            className="inline-block absolute cursor-pointer w-full h-full opacity-0"
                            type="color"
                            value={color}
                            onChange={(e) => {
                                setColor(e.target.value);
                            }}
                            onInput={() => {}}
                        />
                    </div>
                    <span
                        style={{ color: color }}
                        className="text-[18px] font-bold w-full line-clamp-1 underline"
                    >
                        {project.Name}
                    </span>
                </div>
                <span className="text-right font-bold">
                    [{project.Visibility}]
                </span>
            </div>
            <div className="flex flex-col gap-[4px] font-bold">
                <p
                    className={`line-clamp-2 ${!project.Description ? "opacity-50" : ""}`}
                >
                    {project.Description
                        ? project.Description
                        : "No Description"}
                </p>
                <div className="flex flex-col gap-[4px]">
                    <div className="flex justify-between font-semibold">
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
