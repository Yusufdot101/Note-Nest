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
        const token = data.accessToken;
        const { payload } = decodeJWT(token ?? "");

        useAuthStore.getState().setAccessToken(token);
        useAuthStore.getState().setUserID(+payload.sub);
        useAuthStore.getState().setIsLoggedIn(true);
    } catch (error) {
        console.error(error);
    }
};
