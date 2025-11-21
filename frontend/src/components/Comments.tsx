import { useState } from "react";
import Input from "./Input";
import { newComment } from "../utilities/comments";
import SubmitButton from "./SubmitButton";

interface CommentsProps {
    noteID: number;
}

const Comments = ({ noteID }: CommentsProps) => {
    const [comment, setComment] = useState("");
    const handlePost = async () => {
        if (comment.trim() === "") return;
        const success = await newComment(noteID, comment);
        if (!success) return;
        setComment("");
    };

    return (
        <div>
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
            {/*TODO: Display comments*/}
        </div>
    );
};

export default Comments;
