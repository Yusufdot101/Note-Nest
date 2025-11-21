import { api } from "./api";

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
