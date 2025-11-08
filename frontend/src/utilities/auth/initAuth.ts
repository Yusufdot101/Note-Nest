import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";

export const initAuth = async () => {
    try {
        const res = await fetch(`${BASE_APIURL}/auth/refreshtoken`, {
            method: "PUT",
            credentials: "include",
        });
        if (!res.ok) {
            useAuthStore.getState().setIsLoggedIn(false);
            return;
        }
        const data = await res.json();
        useAuthStore.getState().setAccessToken(data.accessToken);
        useAuthStore.getState().setIsLoggedIn(true);
    } catch (error) {
        console.error(error);
    }
};
