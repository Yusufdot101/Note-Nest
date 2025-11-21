import { useState } from "react";
import Input from "./Input";
import { newComment } from "../utilities/comments";

interface CommentsProps {
    noteID: number;
}

const Comments = ({ noteID }: CommentsProps) => {
    console.log(noteID);
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
                className="flex flex-col"
            >
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
            </form>
            {/*TODO: Display comments*/}
        </div>
    );
};

export default Comments;
