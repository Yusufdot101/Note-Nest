import { api } from "./api";

export const newNote = async (
    projectID: number,
    title: string,
    content: string,
    visibility: string,
    color: string,
): Promise<boolean> => {
    try {
        const res = await api(`/projects/${projectID}/notes`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                title,
                content,
                visibility,
                color,
            }),
        });
        if (!res) {
            return false;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            if (errors) {
                console.error(errors);
                return false;
            }
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};
