import { BASE_APIURL } from "../api";

export const forgotPassword = async (
    email: string,
    handleErrors: (errors: Record<string, string>) => void,
): Promise<string> => {
    try {
        const res = await fetch(`${BASE_APIURL}/auth/forgot-password`, {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: email }),
        });

        const data = await res.json();
        if (!res.ok) {
            const error = data.error;
            if (error) {
                handleErrors(error);
                return "";
            }
            throw new Error(`HTTP error! status: ${res.status}`);
        }

        const message = data.message;
        return message;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return "";
    }
};

export const resetPassword = async (
    newPassword: string,
    handleErrors: (errors: Record<string, string>) => void,
): Promise<string> => {
    try {
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get("token");
        const res = await fetch(`${BASE_APIURL}/auth/reset-password`, {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ password: newPassword, token }),
        });

        const data = await res.json();
        if (!res.ok) {
            const error = data.error;
            if (error) {
                handleErrors(error);
                return "";
            }
            throw new Error(`HTTP error! status: ${res.status}`);
        }

        const message = data.message;
        return message;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return "";
    }
};
