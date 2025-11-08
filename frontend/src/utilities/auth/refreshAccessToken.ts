import { useAuthStore } from "../../store/useAuthStore";

export async function refreshAccessToken() {
    const res = await fetch("/auth/refreshtoken", {
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
