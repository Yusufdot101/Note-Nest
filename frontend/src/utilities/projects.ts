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

export const fetchProjects = async (): Promise<Project[]> => {
    try {
        const params = new URLSearchParams(window.location.search);
        const user = params.get("user");
        const res = await api(
            `/projects${user === null ? "" : `?user=${user}`}`,
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
        return data.projects;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return [];
    }
};

export const fetchOneProject = async (): Promise<Project | null> => {
    try {
        const url = new URL(window.location.toString());
        const segments = url.pathname.split("/").filter(Boolean);
        const projectID = segments.at(-1);
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
