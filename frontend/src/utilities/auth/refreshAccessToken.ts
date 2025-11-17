import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";
import { decodeJWT } from "../userIdFromJWT";

export async function refreshAccessToken() {
    const res = await fetch(`${BASE_APIURL}/auth/refreshtoken`, {
        method: "PUT",
        credentials: "include", // important! sends cookie
    });

    if (!res.ok) {
        console.error(
            `Failed to refresh access token: ${res.status} ${res.statusText}`,
        );
        return;
    }

    const data = await res.json();
    const token = data.access_token;
    const { payload } = decodeJWT(token ?? "");

    useAuthStore.getState().setAccessToken(token);
    useAuthStore.getState().setUserID(+payload.sub);
}
