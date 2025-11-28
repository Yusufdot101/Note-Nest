import { api } from "./api";

export interface Comment {
    ID: number;
    Edited: boolean;
    CreatedAt: string;
    UserID: number;
    NoteID: number;
    Content: string;
    LikesCount: number;
    Username: string;
    IsLiked: boolean;
}

export const newComment = async (
    noteID: number,
    comment: string,
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/comments`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ content: comment }),
        });

        if (!res) return false;
        if (!res.ok) return false;

        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

export const fetchComments = async (
    noteID: number,
): Promise<Comment[] | undefined> => {
    try {
        const res = await api(`/notes/${noteID}/comments`);

        if (!res) return undefined;
        if (!res.ok) return undefined;

        const data = await res.json();
        const comments = data.comments;
        if (!Array.isArray(comments)) return undefined;

        const commentsTwo = await Promise.all(
            comments.map(async (comment: Comment) => {
                const isLiked = await commentIsLiked(comment.ID);
                return {
                    ...comment,
                    IsLiked: isLiked,
                };
            }),
        );

        return commentsTwo;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return undefined;
    }
};

export const updateComment = async (
    commentID: number,
    content: string,
): Promise<boolean> => {
    try {
        const res = await api(`/comments/${commentID}`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ content: content }),
        });

        if (!res) return false;
        if (!res.ok) return false;

        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

export const deleteComment = async (commentID: number): Promise<boolean> => {
    try {
        const res = await api(`/comments/${commentID}`, {
            method: "DELETE",
        });

        if (!res) return false;
        if (!res.ok) return false;

        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

export const likeUnlinkeComment = async (
    commentID: number,
    action: "like" | "unlike",
): Promise<boolean> => {
    try {
        const res = await api(`/comments/${commentID}/like`, {
            method: action === "like" ? "POST" : "DELETE",
        });

        if (!res) {
            return false;
        }

        if (!res.ok) {
            return false;
        }

        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

const commentIsLiked = async (commentID: number): Promise<boolean> => {
    try {
        const res = await api(`/comments/${commentID}/like`);

        if (!res) {
            return false;
        }

        if (!res.ok) {
            return false;
        }

        const data = await res.json();
        return data.state ?? false;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};
