import { useState } from "react";
import ColorPicker from "./ColorPicker";
import { likeUnlinkeComment, type Comment } from "../utilities/comments";
import { useAuthStore } from "../store/useAuthStore";
import SubmitButton from "./SubmitButton";
import LikeButton from "./LikeButton";

interface CommentCardProps {
    comment: Comment;
    handleMenuClick?: (
        e: React.MouseEvent<SVGElement>,
        comment: Comment,
    ) => void;
    isEditing: boolean;
    handleSaveEdit: (newContent: string) => void;
    handleCancelEdit: (commentID: number) => void;
}

const CommentCard = ({
    comment,
    handleMenuClick,
    isEditing,
    handleSaveEdit,
    handleCancelEdit,
}: CommentCardProps) => {
    // TODO: enable liking comments and also add color comments
    const [color, setColor] = useState("white");
    const [liked, setLike] = useState(comment.IsLiked);
    const [likesCount, setLikesCount] = useState(comment.LikesCount);

    const [content, setContent] = useState(comment.Content);

    const handleLike = async () => {
        const success = await likeUnlinkeComment(
            comment.ID,
            liked ? "unlike" : "like",
        );
        if (!success) return;
        setLikesCount((prev) => (liked ? prev - 1 : prev + 1));
        setLike((prev) => !prev);
    };

    const userid = useAuthStore((state) => state.userID);

    return (
        <div
            style={{ border: `1px solid ${color}` }}
            className="text-text bg-primary p-[12px] rounded-[8px] flex flex-col gap-[12px] h-fit -[300px]"
            role="group"
            tabIndex={0}
        >
            <div className="flex items-center justify-between gap-[4px]">
                <div className="flex items-center gap-[8px]">
                    <div className="h-[35px]">
                        <ColorPicker
                            color={color}
                            handleChange={
                                comment.UserID === userid
                                    ? (value) => {
                                          setColor(value);
                                      }
                                    : undefined
                            }
                        />
                    </div>

                    <div>
                        <span
                            style={{ color: color }}
                            className="text-[16px] max-[629px]:text-[12px] font-bold w-full line-clamp-1"
                        >
                            {/* Hardcoded for now */}
                            {comment.Username}
                        </span>
                        <span
                            style={{ color: color }}
                            className="text-[16px] max-[629px]:text-[12px] w-full line-clamp-1"
                        >
                            {new Date(comment.CreatedAt).toDateString()}
                        </span>
                    </div>
                </div>

                <div className="flex gap-[12px] items-center">
                    <span>{comment.Edited ? "Edited" : ""}</span>
                    <span>
                        <svg
                            fill="currentColor"
                            version="1.1"
                            id="Icons"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 32 32"
                            className={`${handleMenuClick ? "" : "hidden"} w-[30px] h-[30px] hover:text-accent active:text-text duration-300`}
                            onClick={(e) => {
                                handleMenuClick!(e, comment);
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

            <div hidden={!isEditing} className="flex flex-col gap-[4px]">
                <textarea
                    name="projectDescription"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                    id="projectDescription"
                    className="w-[100%] min-h-[50px] h-fit bg-white rounded-[8px] p-[8px] outline-none text-black"
                />
                <div className="flex max-[619px]:flex-col gap-[4px]">
                    <SubmitButton
                        disabled={content.trim() === comment.Content}
                        handleSubmit={() => handleSaveEdit(content)}
                        aria_label="Save changes"
                        text="Save Changes"
                    />
                    <SubmitButton
                        handleSubmit={() => {
                            setContent(comment.Content);
                            handleCancelEdit(comment.ID);
                        }}
                        aria_label="Cancel changes"
                        text="Cancel Changes"
                        bgColor="red"
                    />
                </div>
            </div>
            <div hidden={isEditing}>{comment.Content}</div>

            <div className="flex flex-col gap-[12px] font-bold">
                <div className="flex flex-col gap-[4px]">
                    <div className="flex gap-[12px] font-semibold">
                        <LikeButton
                            onToggle={handleLike}
                            liked={liked}
                            count={likesCount}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default CommentCard;
