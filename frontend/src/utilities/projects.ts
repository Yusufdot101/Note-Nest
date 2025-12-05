import type { Project } from "../components/ProjectCard";
import { api } from "./api";

export interface Metadata {
    TotalResources: number;
    PageSize: number;
    TotalPages: number;
    PageNumber: number;
}

export const fetchProjects = async (
    options: Map<string, string | number>,
): Promise<{ projects: Project[]; metadata: Metadata } | undefined> => {
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
            return;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            console.error(errors);
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const projects = data.projects;
        const metadata = data.metadata;
        if (!Array.isArray(projects) || !metadata) {
            return;
        }

        return { projects, metadata };
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return;
    }
};
