import React, { useEffect, useRef, useState } from "react";
import type { Note } from "../components/NoteCard";
import { useNavigate, useParams } from "react-router-dom";
import {
    fetchNote,
    fetchNoteOwner,
    likeUnlinkeNote,
    noteIsLiked,
    noteIsSaved,
    saveUnsaveNote,
} from "../utilities/note";
import ReactMarkdown from "react-markdown";
import { useAuthStore } from "../store/useAuthStore";
import type { Project } from "../components/ProjectCard";
import { fetchProject } from "../utilities/project";
import NoteActionsDialoge from "../components/NoteActionsDialoge";
import Comments from "../components/Comments";
import LikeButton from "../components/LikeButton";
import SaveButton from "../components/SaveButton";
import ShareButton from "../components/ShareButton";
import CommentButton from "../components/CommentButton";
import PublicOrPrivate from "../components/PublicOrPrivate";
import Menu from "../components/Menu";

const NotePage = () => {
    const [note, setNote] = useState<Note>();
    const [commentsCount, setCommentsCount] = useState(0);
    const [noteOwner, setNoteOwner] = useState("");
    const [project, setProject] = useState<Project>();
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
        setNote((prev) => {
            if (!prev) return prev;
            return {
                ...prev,
                LikesCount: liked ? prev.LikesCount - 1 : prev.LikesCount + 1,
            };
        });
    };

    const [saved, setSaved] = useState(false);
    const handleSaved = async () => {
        if (!noteid) return;
        const success = await saveUnsaveNote(
            +noteid,
            saved ? "unsave" : "save",
        );
        if (!success) return;
        setSaved((prev) => !prev);
        setNote((prev) => {
            if (!prev) return prev;
            return {
                ...prev,
                SavesCount: saved ? prev.SavesCount - 1 : prev.SavesCount + 1,
            };
        });
    };

    const userid = useAuthStore((state) => state.userID);

    const accessToken = useAuthStore((state) => state.accessToken);
    useEffect(() => {
        const setupNote = async () => {
            if (!noteid) return;
            const note = await fetchNote(+noteid);
            if (!note) return;
            setNote(note);
            setCommentsCount(note.CommentsCount);
        };

        const setupNoteOwner = async () => {
            if (!noteid) return;
            const noteOwner = await fetchNoteOwner(+noteid);
            if (!noteOwner) return;
            setNoteOwner(noteOwner);
        };

        const setupProject = async () => {
            if (!projectid) return;
            const project = await fetchProject(+projectid);
            if (!project) return;
            setProject(project);
        };

        const setupLiked = async () => {
            if (!noteid) return;
            const liked = await noteIsLiked(+noteid);
            setLike(liked);
        };

        const setupSaved = async () => {
            if (!noteid) return;
            const saved = await noteIsSaved(+noteid);
            setSaved(saved);
        };

        setupLiked();
        setupSaved();

        setupNote();
        setupNoteOwner();
        setupProject();
    }, [noteid, projectid, accessToken]);

    const handleMenuClick = (
        e: React.MouseEvent<SVGElement> | React.KeyboardEvent<SVGElement>,
    ) => {
        e.stopPropagation();
        setShowDialoge((prev) => !prev);
    };

    const commentsRef = useRef<HTMLDivElement | null>(null);

    // in case the clipboard api doesnt work
    function fallbackCopy(text: string): boolean {
        const textarea = document.createElement("textarea");
        textarea.value = text;
        textarea.style.position = "fixed";
        textarea.style.opacity = "0";

        document.body.appendChild(textarea);

        textarea.focus();
        textarea.select();

        try {
            document.execCommand("copy");
            return true;
        } catch (err) {
            console.error("fallback failed", err);
            return false;
        } finally {
            document.body.removeChild(textarea);
        }
    }

    async function copyToClipboard(text: string): Promise<boolean> {
        if (
            navigator.clipboard &&
            typeof navigator.clipboard.writeText === "function"
        ) {
            try {
                await navigator.clipboard.writeText(text);
                return true;
            } catch (e) {
                console.error("clipboard API failed:", e);
            }
        }

        // fallback always works
        return fallbackCopy(text);
    }

    const handleShare = async () => {
        const success = await copyToClipboard(window.location.toString());
        if (!success) {
            alert("an error occurred and could not copy to clipboard");
            return;
        }
        alert("copied link to clipboard!");
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
                    <PublicOrPrivate
                        visibility={note?.Visibility ?? "private"}
                    />
                    <span>
                        <Menu
                            userId={userid}
                            projectUserId={project?.UserID}
                            handleClick={handleMenuClick}
                        />
                    </span>
                </div>
            </div>

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px]"
            >
                <p>By: {noteOwner}</p>
                <p>Created: {new Date(note?.CreatedAt || "").toDateString()}</p>
                <div className="flex items-center justify-between">
                    <div className="flex items-center gap-x-[20px]">
                        <LikeButton
                            onToggle={handleLike}
                            count={note?.LikesCount ?? 0}
                            liked={liked}
                        />

                        <CommentButton
                            handleClick={() => {
                                if (!commentsRef.current) return;
                                commentsRef.current.scrollIntoView({
                                    behavior: "smooth",
                                });
                                commentsRef.current.focus();
                            }}
                            count={commentsCount}
                        />
                    </div>

                    <div className="flex gap-x-[20px]">
                        <SaveButton
                            onToggle={handleSaved}
                            count={note?.SavesCount ?? 0}
                            saved={saved}
                        />
                        <ShareButton handleClick={handleShare} />
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
                        <LikeButton
                            onToggle={handleLike}
                            count={note?.LikesCount ?? 0}
                            liked={liked}
                        />

                        <CommentButton
                            handleClick={() => {
                                if (!commentsRef.current) return;
                                commentsRef.current.scrollIntoView({
                                    behavior: "smooth",
                                });
                                commentsRef.current.focus();
                            }}
                            count={commentsCount}
                        />
                    </div>

                    <div className="flex gap-x-[20px]">
                        <SaveButton
                            onToggle={handleSaved}
                            count={note?.SavesCount ?? 0}
                            saved={saved}
                        />

                        <ShareButton handleClick={handleShare} />
                    </div>
                </div>
            </div>

            <div
                style={{ border: `1px solid ${note?.Color}` }}
                className="bg-primary p-[12px] rounded-[8px] flex flex-col gap-y-[8px]"
            >
                {noteid && (
                    <Comments
                        commentsRef={commentsRef}
                        handleCommentsCountChange={(newCount) =>
                            setCommentsCount(newCount)
                        }
                        noteID={+noteid}
                    />
                )}
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
