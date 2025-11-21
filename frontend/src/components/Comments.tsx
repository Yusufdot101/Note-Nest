import { useCallback, useEffect, useState } from "react";
import Input from "./Input";
import { fetchComments, newComment, type Comment } from "../utilities/comments";
import SubmitButton from "./SubmitButton";
import CommentCard from "./CommentCard";
import CommentActionsDialog from "./CommentsActionsDialog";
import { useAuthStore } from "../store/useAuthStore";

interface CommentsProps {
    noteID: number;
}

const Comments = ({ noteID }: CommentsProps) => {
    const [comment, setComment] = useState("");
    const [clickedComment, setClickedComment] = useState<Comment>();
    const [comments, setComments] = useState<Comment[]>([]);
    const [showDialog, setShowDialog] = useState(false);

    const userid = useAuthStore((state) => state.userID);

    const setupComments = useCallback(async () => {
        const comments = await fetchComments(noteID);
        if (!comments) return;
        setComments(comments);
    }, [noteID]);

    useEffect(() => {
        setupComments();
    }, [setupComments]);

    const handlePost = async () => {
        if (comment.trim() === "") return;
        const success = await newComment(noteID, comment);
        if (!success) return;
        setComment("");
        setupComments();
    };

    return (
        <div className="flex flex-col gap-[12px]">
            <p className="font-bold text-[32px] text-center max-[619px]:text-[20px]">
                Comments
            </p>
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    e.stopPropagation();
                    handlePost();
                }}
                className="flex flex-col gap-[4px] text-[20px]"
            >
                <div>
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
                <SubmitButton
                    aria_label="add comment"
                    type="submit"
                    text="Post Comment"
                    handleSubmit={() => {}}
                />
            </form>

            <div className="flex flex-col gap-y-[8px]">
                {comments.map((comment) => (
                    <CommentCard
                        key={comment.ID}
                        comment={comment}
                        handleMenuClick={
                            comment.UserID === userid
                                ? (e) => {
                                      e.stopPropagation();
                                      setShowDialog(true);
                                      setClickedComment(comment);
                                  }
                                : undefined
                        }
                    />
                ))}
            </div>

            {showDialog && comments && (
                <CommentActionsDialog
                    comment={clickedComment ? clickedComment : undefined}
                    color="white"
                    handleClose={() => {
                        setShowDialog(false);
                    }}
                />
            )}
        </div>
    );
};

export default Comments;
