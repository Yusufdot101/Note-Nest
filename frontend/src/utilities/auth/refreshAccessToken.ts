import { useAuthStore } from "../../store/useAuthStore";

export async function refreshAccessToken() {
    const res = await fetch("/auth/refresh", {
        method: "PUT",
        credentials: "include" // important! sends cookie
    });

    if (!res.ok) return;

    const data = await res.json();
    useAuthStore.getState().setAccessToken(data.accessToken);
}
