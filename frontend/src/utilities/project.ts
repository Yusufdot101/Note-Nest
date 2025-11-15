import type { Project } from "../components/ProjectCard";
import { api } from "./api";

export const newProject = async (
    projectName: string,
    projectDescription: string,
    projectVisibility: string,
    projectColor: string,
): Promise<boolean> => {
    try {
        const res = await api("/projects", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: projectName,
                description: projectDescription,
                visibility: projectVisibility,
                color: projectColor,
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

export const updateProject = async (
    projectID: number,
    projectName: string,
    projectDescription: string,
    projectVisibility: string,
    projectColor: string,
): Promise<boolean> => {
    try {
        const res = await api(`/projects/${projectID}`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: projectName,
                description: projectDescription,
                visibility: projectVisibility,
                color: projectColor,
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

export const fetchProject = async (
    projectID: number,
): Promise<Project | null> => {
    try {
        const res = await api(`/projects/${projectID}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        });
        if (!res) {
            return null;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        return data.project;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return null;
    }
};

export const deleteProject = async (projectID: number): Promise<boolean> => {
    try {
        if (!confirm("are you sure you want to delete this project?")) {
            return false;
        }

        if (!projectID) {
            alert("Invalid project ID");
            return false;
        }

        const res = await api(`/projects/${projectID}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
        });

        if (!res) {
            return false;
        }

        if (!res.ok) {
            const data = await res.json();
            console.error(data.error);
            return false;
        }

        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};
