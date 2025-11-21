import React, { useEffect, useState } from "react";
import type { Note } from "../components/NoteCard";
import { useNavigate, useParams } from "react-router-dom";
import { fetchNote, likeUnlinkeNote, noteIsLiked } from "../utilities/note";
import ReactMarkdown from "react-markdown";
import Input from "../components/Input";
import { useAuthStore } from "../store/useAuthStore";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import NoteActionsDialoge from "../components/NoteActionsDialoge";

const NotePage = () => {
    const [note, setNote] = useState<Note>();
    const [project, setProject] = useState<Project>();
    const [comment, setComment] = useState("");
    const [showDialoge, setShowDialoge] = useState(false);

    const { projectid, noteid } = useParams();
    const navigate = useNavigate();

    const [liked, setLike] = useState(false);
    const handleLike = async () => {
        if (!noteid) return;
        const success = await likeUnlinkeNote(
            +noteid,
            liked ? "unlike" : "like",
        );
        if (!success) return;
        setLike((prev) => !prev);
    };

    const userid = useAuthStore((state) => state.userID);

    useEffect(() => {
        const setupLiked = async () => {
            if (!noteid) return;
            const liked = await noteIsLiked(+noteid);
            setLike(liked);
        };

        setupLiked();
    }, [noteid]);

    useEffect(() => {
        const setupNote = async () => {
            if (!noteid) return;
            const note = await fetchNote(+noteid);
            if (!note) return;
            setNote(note);
        };

        const setupProject = async () => {
            if (!projectid) return;
            const project = await fetchProject(+projectid);
            if (!project) return;
            setProject(project);
        };

        setupNote();
        setupProject();
    }, [noteid, projectid, liked]);

    const handleMenuClick = (
        e: React.MouseEvent<SVGElement> | React.KeyboardEvent<SVGElement>,
    ) => {
        e.stopPropagation();
        setShowDialoge((prev) => !prev);
    };

    return (
        <div className="text-[20px] max-[619px]:text-[16px] text-text flex flex-col gap-[8px]">
            <div
                aria-label="back to project page"
                className="flex text-[24px] max-[619px]:text-[16px] bg-accent p-[12px] rounded-[8px] justify-center gap-[12px] cursor-pointer hover:gap-[20px] duration-300 items-center"
                role="button"
                onClick={() => navigate(`/projects/${projectid}`)}
            >
                <svg
                    className="w-[32px] h-[32px] max-[619px]:w-[24px] max-[619px]:h-[32px]"
                    fill="currentColor"
                    viewBox="0 0 52 52"
                    id="Layer_1"
                    xmlns="http://www.w3.org/2000/svg"
                >
                    <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                    <g
                        id="SVGRepo_tracerCarrier"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                    ></g>
                    <g id="SVGRepo_iconCarrier">
                        <path d="M50,24H6.83L27.41,3.41a2,2,0,0,0,0-2.82,2,2,0,0,0-2.82,0l-24,24a1.79,1.79,0,0,0-.25.31A1.19,1.19,0,0,0,.25,25c0,.07-.07.13-.1.2l-.06.2a.84.84,0,0,0,0,.17,2,2,0,0,0,0,.78.84.84,0,0,0,0,.17l.06.2c0,.07.07.13.1.2a1.19,1.19,0,0,0,.09.15,1.79,1.79,0,0,0,.25.31l24,24a2,2,0,1,0,2.82-2.82L6.83,28H50a2,2,0,0,0,0-4Z"></path>
                    </g>
                </svg>

                <span>Back to project</span>
            </div>

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px] flex items-center justify-between"
            >
                <span className="font-bold text-[28px] max-[619px]:text-[20px]">
                    {note?.Title}
                </span>
                <div className="flex items-center gap-[8px]">
                    <span className="text-[20px] font-light">
                        {note?.Visibility === "public" ? (
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
                            aria-label="open note actions menu"
                            tabIndex={0}
                            onKeyDown={(e) => {
                                if (e.key === "Enter" || e.key === " ") {
                                    handleMenuClick(e);
                                }
                            }}
                            fill="currentColor"
                            version="1.1"
                            id="Icons"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 32 32"
                            className={`${userid === project?.UserID ? "" : "hidden"} w-[30px] h-[30px] hover:text-accent active:text-text duration-300`}
                            onClick={(e) => {
                                handleMenuClick(e);
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

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px]"
            >
                {/*TODO: show username*/}
                <p>By: </p>
                <p>Created: {new Date(note?.CreatedAt || "").toDateString()}</p>
                <div className="flex items-center justify-between">
                    <div className="flex items-center gap-x-[20px]">
                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                onClick={handleLike}
                                onKeyDown={(e) => {
                                    if (e.key === "Enter" || e.key === " ") {
                                        handleLike();
                                    }
                                }}
                                role="button"
                                aria-label="like note"
                                tabIndex={0}
                                width="28"
                                height="28"
                                className="cursor-pointer"
                                viewBox="0 0 24 24"
                                fill={`${liked ? "white" : "none"}`}
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
                                        d="M8 10V20M8 10L4 9.99998V20L8 20M8 10L13.1956 3.93847C13.6886 3.3633 14.4642 3.11604 15.1992 3.29977L15.2467 3.31166C16.5885 3.64711 17.1929 5.21057 16.4258 6.36135L14 9.99998H18.5604C19.8225 9.99998 20.7691 11.1546 20.5216 12.3922L19.3216 18.3922C19.1346 19.3271 18.3138 20 17.3604 20L8 20"
                                        stroke="#ffffff"
                                        strokeWidth="1.5"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    ></path>{" "}
                                </g>
                            </svg>
                            <span>{note?.LikesCount}</span>
                        </div>

                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="white"
                                className="cursor-pointer"
                            >
                                <path d="M18.006 16.803c1.533-1.456 2.234-3.325 2.234-5.321C20.24 7.357 16.709 4 12.191 4S4 7.357 4 11.482c0 4.126 3.674 7.482 8.191 7.482.817 0 1.622-.111 2.393-.327.231.2.48.391.744.559 1.06.693 2.203 1.044 3.399 1.044.224-.008.4-.112.486-.287a.49.49 0 0 0-.042-.518c-.495-.67-.845-1.364-1.04-2.057a4 4 0 0 1-.125-.598zm-3.122 1.055-.067-.223-.315.096a8 8 0 0 1-2.311.338c-4.023 0-7.292-2.955-7.292-6.587 0-3.633 3.269-6.588 7.292-6.588 4.014 0 7.112 2.958 7.112 6.593 0 1.794-.608 3.469-2.027 4.72l-.195.168v.255c0 .056 0 .151.016.295.025.231.081.478.154.733.154.558.398 1.117.722 1.659a5.3 5.3 0 0 1-2.165-.845c-.276-.176-.714-.383-.941-.59z"></path>
                            </svg>
                            <span>{note?.CommentsCount}</span>
                        </div>
                    </div>

                    <div className="flex gap-x-[20px]">
                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="25"
                                height="25"
                                fill="white"
                                className="cursor-pointer"
                                viewBox="0 0 25 25"
                                aria-label="Add to list bookmark button"
                            >
                                <path
                                    fill="currentColor"
                                    d="M18 2.5a.5.5 0 0 1 1 0V5h2.5a.5.5 0 0 1 0 1H19v2.5a.5.5 0 1 1-1 0V6h-2.5a.5.5 0 0 1 0-1H18zM7 7a1 1 0 0 1 1-1h3.5a.5.5 0 0 0 0-1H8a2 2 0 0 0-2 2v14a.5.5 0 0 0 .805.396L12.5 17l5.695 4.396A.5.5 0 0 0 19 21v-8.5a.5.5 0 0 0-1 0v7.485l-5.195-4.012a.5.5 0 0 0-.61 0L7 19.985z"
                                ></path>
                            </svg>
                            <span>{5}</span>
                        </div>

                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                fill="white"
                                className="cursor-pointer"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    fill="currentColor"
                                    fillRule="evenodd"
                                    d="M15.218 4.931a.4.4 0 0 1-.118.132l.012.006a.45.45 0 0 1-.292.074.5.5 0 0 1-.3-.13l-2.02-2.02v7.07c0 .28-.23.5-.5.5s-.5-.22-.5-.5v-7.04l-2 2a.45.45 0 0 1-.57.04h-.02a.4.4 0 0 1-.16-.3.4.4 0 0 1 .1-.32l2.8-2.8a.5.5 0 0 1 .7 0l2.8 2.79a.42.42 0 0 1 .068.498m-.106.138.008.004v-.01zM16 7.063h1.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-11c-1.1 0-2-.9-2-2v-10a2 2 0 0 1 2-2H8a.5.5 0 0 1 .35.15.5.5 0 0 1 .15.35.5.5 0 0 1-.15.35.5.5 0 0 1-.35.15H6.4c-.5 0-.9.4-.9.9v10.2a.9.9 0 0 0 .9.9h11.2c.5 0 .9-.4.9-.9v-10.2c0-.5-.4-.9-.9-.9H16a.5.5 0 0 1 0-1"
                                    clipRule="evenodd"
                                ></path>
                            </svg>
                            <span>{3}</span>
                        </div>
                    </div>
                </div>
            </div>

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px] flex flex-col gap-y-[24px]"
            >
                <div className="markdown">
                    <ReactMarkdown>{note?.Content}</ReactMarkdown>
                </div>

                <div className="flex items-center justify-between">
                    <div className="flex items-center gap-x-[20px]">
                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                onClick={handleLike}
                                onKeyDown={(e) => {
                                    if (e.key === "Enter" || e.key === " ") {
                                        handleLike();
                                    }
                                }}
                                role="button"
                                aria-label="like note"
                                tabIndex={0}
                                width="28"
                                height="28"
                                className="cursor-pointer"
                                viewBox="0 0 24 24"
                                fill={`${liked ? "white" : "none"}`}
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
                                        d="M8 10V20M8 10L4 9.99998V20L8 20M8 10L13.1956 3.93847C13.6886 3.3633 14.4642 3.11604 15.1992 3.29977L15.2467 3.31166C16.5885 3.64711 17.1929 5.21057 16.4258 6.36135L14 9.99998H18.5604C19.8225 9.99998 20.7691 11.1546 20.5216 12.3922L19.3216 18.3922C19.1346 19.3271 18.3138 20 17.3604 20L8 20"
                                        stroke="#ffffff"
                                        strokeWidth="1.5"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    ></path>{" "}
                                </g>
                            </svg>
                            <span>{note?.LikesCount}</span>
                        </div>

                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="white"
                                className="cursor-pointer"
                            >
                                <path d="M18.006 16.803c1.533-1.456 2.234-3.325 2.234-5.321C20.24 7.357 16.709 4 12.191 4S4 7.357 4 11.482c0 4.126 3.674 7.482 8.191 7.482.817 0 1.622-.111 2.393-.327.231.2.48.391.744.559 1.06.693 2.203 1.044 3.399 1.044.224-.008.4-.112.486-.287a.49.49 0 0 0-.042-.518c-.495-.67-.845-1.364-1.04-2.057a4 4 0 0 1-.125-.598zm-3.122 1.055-.067-.223-.315.096a8 8 0 0 1-2.311.338c-4.023 0-7.292-2.955-7.292-6.587 0-3.633 3.269-6.588 7.292-6.588 4.014 0 7.112 2.958 7.112 6.593 0 1.794-.608 3.469-2.027 4.72l-.195.168v.255c0 .056 0 .151.016.295.025.231.081.478.154.733.154.558.398 1.117.722 1.659a5.3 5.3 0 0 1-2.165-.845c-.276-.176-.714-.383-.941-.59z"></path>
                            </svg>
                            <span>{note?.CommentsCount}</span>
                        </div>
                    </div>

                    <div className="flex gap-x-[20px]">
                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="25"
                                height="25"
                                fill="white"
                                className="cursor-pointer"
                                viewBox="0 0 25 25"
                                aria-label="Add to list bookmark button"
                            >
                                <path
                                    fill="currentColor"
                                    d="M18 2.5a.5.5 0 0 1 1 0V5h2.5a.5.5 0 0 1 0 1H19v2.5a.5.5 0 1 1-1 0V6h-2.5a.5.5 0 0 1 0-1H18zM7 7a1 1 0 0 1 1-1h3.5a.5.5 0 0 0 0-1H8a2 2 0 0 0-2 2v14a.5.5 0 0 0 .805.396L12.5 17l5.695 4.396A.5.5 0 0 0 19 21v-8.5a.5.5 0 0 0-1 0v7.485l-5.195-4.012a.5.5 0 0 0-.61 0L7 19.985z"
                                ></path>
                            </svg>
                            <span>{5}</span>
                        </div>

                        <div className="flex items-center gap-x-[4px]">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                fill="white"
                                className="cursor-pointer"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    fill="currentColor"
                                    fillRule="evenodd"
                                    d="M15.218 4.931a.4.4 0 0 1-.118.132l.012.006a.45.45 0 0 1-.292.074.5.5 0 0 1-.3-.13l-2.02-2.02v7.07c0 .28-.23.5-.5.5s-.5-.22-.5-.5v-7.04l-2 2a.45.45 0 0 1-.57.04h-.02a.4.4 0 0 1-.16-.3.4.4 0 0 1 .1-.32l2.8-2.8a.5.5 0 0 1 .7 0l2.8 2.79a.42.42 0 0 1 .068.498m-.106.138.008.004v-.01zM16 7.063h1.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-11c-1.1 0-2-.9-2-2v-10a2 2 0 0 1 2-2H8a.5.5 0 0 1 .35.15.5.5 0 0 1 .15.35.5.5 0 0 1-.15.35.5.5 0 0 1-.35.15H6.4c-.5 0-.9.4-.9.9v10.2a.9.9 0 0 0 .9.9h11.2c.5 0 .9-.4.9-.9v-10.2c0-.5-.4-.9-.9-.9H16a.5.5 0 0 1 0-1"
                                    clipRule="evenodd"
                                ></path>
                            </svg>
                            <span>{3}</span>
                        </div>
                    </div>
                </div>
            </div>

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px] flex flex-col gap-y-[8px]"
            >
                <p className="font-bold text-[32px] text-center max-[619px]:text-[20px]">
                    Comments
                </p>
                <div className="flex flex-col">
                    <Input
                        minLength={1}
                        isRequired
                        inputType="string"
                        labelString="What are your thoughts?"
                        inputValue={comment}
                        handleChange={(value: string) => {
                            setComment(value);
                        }}
                        inputId="comment"
                        inputName="comment"
                    />
                </div>
                {/*TODO: Display comments*/}
            </div>

            {showDialoge && note ? (
                <NoteActionsDialoge
                    color={note.Color}
                    note={note}
                    handleClose={() => setShowDialoge(false)}
                />
            ) : undefined}
        </div>
    );
};

export default NotePage;
