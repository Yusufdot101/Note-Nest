import { api } from "./api";

export const newProject = async (
    projectName: string,
    projectDescription: string,
    projectVisibility: string,
    handleErrors: (errors: Record<string, string>) => void,
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
            }),
        });
        if (!res) {
            return false;
        }
        const data = await res.json();
        if (!res.ok) {
            const errors = data.error;
            if (errors) {
                handleErrors(errors);
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
