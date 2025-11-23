import { useEffect, useRef } from "react";
import type { Comment } from "../utilities/comments";

interface CommentActionsDialogProps {
    comment?: Comment;
    color: string;
    handleClose: () => void;
    handleClickEdit: (id: number) => void;
}

const CommentActionsDialog = ({
    comment,
    color,
    handleClose,
    handleClickEdit,
}: CommentActionsDialogProps) => {
    // TODO: enable actions like editing and deleting comment
    const ref = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (!ref.current?.contains(e.target as Node)) {
                handleClose();
            }
        };

        const handleEscape = (e: KeyboardEvent) => {
            if (e.key === "Escape") {
                handleClose();
            }
        };
        document.addEventListener("click", handleClick);
        document.addEventListener("keydown", handleEscape);

        return () => {
            document.removeEventListener("click", handleClick);
            document.removeEventListener("keydown", handleEscape);
        };
    }, [handleClose]);

    return (
        <div
            ref={ref}
            className="fixed text-[28px] z-2 max-[629px]:text-[20px] top-0 mt-[30vh] left-0 right-0 m-auto h-fit w-[90vw] max-w-[700px] bg-primary border-[1px] border-solid border-[#ffffff] rounded-[8px] text-text p-[12px]"
        >
            <p
                style={{ color: color }}
                className="font-bold text-[32px] max-[629px]:text-[24px] p-[8px] text-center"
            >
                Comment Actions
            </p>
            <ul className="flex flex-col gap-[8px]">
                <li
                    onClick={() => {
                        if (!comment?.ID) return;
                        handleClickEdit(comment.ID);
                    }}
                    className="bg-[#747474] p-[8px] hover:opacity-80 duration-300 cursor-pointer"
                >
                    Edit Comment
                </li>
                <li className="bg-[#FF0000] p-[8px] hover:opacity-80 duration-300 cursor-pointer">
                    Delete Comment
                </li>
            </ul>
        </div>
    );
};

export default CommentActionsDialog;
