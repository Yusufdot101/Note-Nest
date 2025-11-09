import type { ProjectCardProps } from "../components/ProjectCard";
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

export const fetchProjects = async (): Promise<ProjectCardProps[]> => {
    try {
        const params = new URLSearchParams(window.location.search);
        const user = params.get("user");
        const res = await api(
            `/projects${user == undefined ? "" : `?user=${user}`}`,
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
