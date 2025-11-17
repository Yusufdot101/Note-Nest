import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";
import { decodeJWT } from "../userIdFromJWT";

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
        const token = data.access_token;
        const { payload } = decodeJWT(token ?? "");

        if (!payload || !payload.sub) {
            console.error("invalid JWT payload");
            useAuthStore.getState().clearAccessToken();
            return;
        }

        const userId = +payload.sub;
        if (isNaN(userId)) {
            console.error("invalid user ID in JWT");
            useAuthStore.getState().clearAccessToken();
            return;
        }

        useAuthStore.getState().setAccessToken(token);
        useAuthStore.getState().setUserID(userId);
        useAuthStore.getState().setIsLoggedIn(true);
    } catch (error) {
        console.error(error);
    }
};
