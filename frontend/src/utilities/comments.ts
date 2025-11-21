import { api } from "./api";

export interface Comment {
    ID: number;
    CreatedAt: string;
    UserID: number;
    NoteID: number;
    Content: string;
    LikesCount: number;
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

        return comments;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return undefined;
    }
};
