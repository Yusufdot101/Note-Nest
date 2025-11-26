import React, { useCallback, useEffect, useState } from "react";
import Input from "./Input";
import {
    deleteComment,
    fetchComments,
    newComment,
    updateComment,
    type Comment,
} from "../utilities/comments";
import SubmitButton from "./SubmitButton";
import CommentCard from "./CommentCard";
import CommentActionsDialog from "./CommentsActionsDialog";
import { useAuthStore } from "../store/useAuthStore";

interface CommentsProps {
    noteID: number;
    handleCommentsCountChange: (newCount: number) => void;
    commentsRef: React.Ref<HTMLDivElement | null>;
}

const Comments = ({
    noteID,
    handleCommentsCountChange,
    commentsRef,
}: CommentsProps) => {
    const [comment, setComment] = useState("");
    const [clickedComment, setClickedComment] = useState<Comment>();
    const [currentEditingID, setCurrentEditingID] = useState<
        number | undefined
    >();

    const [comments, setComments] = useState<Comment[]>([]);
    const [showDialog, setShowDialog] = useState(false);

    const userid = useAuthStore((state) => state.userID);

    const setupComments = useCallback(async () => {
        const comments = await fetchComments(noteID);
        if (!comments) return;
        setComments(comments);
        handleCommentsCountChange(comments.length);
    }, [noteID, handleCommentsCountChange]);

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

    const handleEdit = async (commentID: number, newContent: string) => {
        if (newContent.trim() === "") return;
        const success = await updateComment(commentID, newContent);
        if (!success) return;
        setupComments();
        setCurrentEditingID(undefined);
    };

    const handleDelete = async (commentID: number) => {
        if (!confirm("Are you sure you want to delete this comment?")) return;
        const success = await deleteComment(commentID);
        if (!success) return;
        setupComments();
    };

    return (
        <div
            ref={commentsRef}
            tabIndex={-1}
            className="flex flex-col gap-[12px]"
        >
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
                        handleSaveEdit={(newContent: string) => {
                            handleEdit(comment.ID, newContent);
                        }}
                        handleCancelEdit={(commentID) => {
                            setCurrentEditingID((prev) =>
                                prev === commentID ? undefined : prev,
                            );
                        }}
                        isEditing={currentEditingID === comment.ID}
                        key={comment.ID}
                        comment={comment}
                        handleMenuClick={
                            comment.UserID === userid
                                ? (e, comment) => {
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
                    handleClickEdit={(id: number) => {
                        setCurrentEditingID(id);
                        setShowDialog(false);
                    }}
                    handleClickDelete={(id: number) => {
                        setShowDialog(false);
                        handleDelete(id);
                    }}
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
