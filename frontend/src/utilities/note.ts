import type { Note } from "../components/NoteCard";
import { api } from "./api";

export const newNote = async (
    projectID: number,
    title: string,
    content: string,
    visibility: string,
    color: string,
    handleError: (errors: Record<string, string>) => void,
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
                handleError(errors);
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

export const fetchNotes = async (
    options: Map<string, string | number>,
): Promise<Note[]> => {
    try {
        const params = new URLSearchParams();
        for (const [key, value] of options) {
            params.append(key, String(value));
        }

        const queryString = params.toString();
        const res = await api(`/notes${queryString ? `?${queryString}` : ""}`);

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

export const fetchNote = async (noteID: number): Promise<Note | undefined> => {
    try {
        const res = await api(`/notes/${noteID}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        });
        if (!res) {
            return;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return data.note;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return undefined;
    }
};

export const fetchNoteOwner = async (noteID: number): Promise<string> => {
    try {
        const res = await api(`/notes/${noteID}/username`);
        if (!res) {
            return "";
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return data.username;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return "";
    }
};

export const deleteNote = async (noteID: number): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}`, {
            method: "DELETE",
        });
        if (!res) {
            return false;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

export const toggleVisibility = async (
    noteID: number,
    newVisibility: string,
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/visibility`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                visibility: newVisibility,
            }),
        });
        if (!res) {
            return false;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};

export const editNote = async (
    noteid: number,
    title: string,
    content: string,
    handleError: (errors: Record<string, string>) => void,
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteid}/content`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                title,
                content,
            }),
        });
        if (!res) {
            return false;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            if (errors) {
                handleError(errors);
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

export const editNoteColor = async (
    noteid: number,
    newColor: string,
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteid}/color`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                color: newColor,
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

export const likeUnlinkeNote = async (
    noteID: number,
    action: "like" | "unlike",
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/like`, {
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

export const noteIsLiked = async (noteID: number): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/like`);

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

export const saveUnsaveNote = async (
    noteID: number,
    action: "save" | "unsave",
): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/save`, {
            method: action === "save" ? "POST" : "DELETE",
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

export const noteIsSaved = async (noteID: number): Promise<boolean> => {
    try {
        const res = await api(`/notes/${noteID}/save`);

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

export const fetchSavedNotes = async (): Promise<Note[]> => {
    try {
        const res = await api(`/saved/notes`);

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
