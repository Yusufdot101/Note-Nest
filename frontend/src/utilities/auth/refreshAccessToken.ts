import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";

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
    useAuthStore.getState().setAccessToken(data.access_token);
}
