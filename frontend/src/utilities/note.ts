import type { Note } from "../components/NoteCard";
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

export const fetchNotes = async (projectID: number): Promise<Note[]> => {
    try {
        const params = new URLSearchParams(window.location.search);
        const user = params.get("user");
        const res = await api(
            `/notes?${user === null ? "" : `user=${user}&`}${projectID ? "" : `projectid=${projectID}&`}`,
            {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );
        if (!res) {
            return [];
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return data.notes;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return [];
    }
};
