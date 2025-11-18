import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";
import { decodeJWT } from "../userIdFromJWT";

export const login = async (
    email: string,
    password: string,
    handleErrors: (error: string) => void,
): Promise<boolean> => {
    try {
        const res = await fetch(`${BASE_APIURL}/auth/login`, {
            method: "PUT",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: email, password: password }),
        });

        const data = await res.json();
        if (!res.ok) {
            const error = data.error;
            if (error) {
                handleErrors(error);
                return false;
            }
            throw new Error(`HTTP error! status: ${res.status}`);
        }

        const token = data.access_token;
        const { payload } = decodeJWT(token ?? "");

        if (!payload || !payload.sub) {
            console.error("invalid JWT payload");
            useAuthStore.getState().clearAccessToken();
            return false;
        }

        const userId = +payload.sub;
        if (isNaN(userId)) {
            console.error("invalid user ID in JWT");
            useAuthStore.getState().clearAccessToken();
            return false;
        }

        useAuthStore.getState().setAccessToken(token);
        useAuthStore.getState().setUserID(userId);
        useAuthStore.getState().setIsLoggedIn(true);
        return true;
    } catch (error) {
        alert("an error occurred, please try again");
        console.error(error);
        return false;
    }
};
