import type { Project } from "../components/ProjectCard";
import { api } from "./api";

export const fetchProjects = async (
    options: Map<string, string | number>,
): Promise<Project[]> => {
    try {
        const params = new URLSearchParams();
        for (const [key, value] of options) {
            params.append(key, String(value));
        }

        const res = await api(`/projects?${params.toString()}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        });

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
