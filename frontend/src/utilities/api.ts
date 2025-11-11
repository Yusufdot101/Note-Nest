import { useAuthStore } from "../store/useAuthStore";
import { refreshAccessToken } from "./auth/refreshAccessToken";

export const BASE_APIURL =
    import.meta.env.VITE_API_URL || "http://localhost:8080";

export const api = async (path: string, options: RequestInit = {}) => {
    const { accessToken } = useAuthStore.getState();
    const url = path.startsWith("http") ? path : `${BASE_APIURL}${path}`;

    try {
        let res = await fetch(url, {
            ...options,
            credentials: "include",
            headers: {
                ...(options.headers || {}),
                Authorization: accessToken ? `Bearer ${accessToken}` : "",
            },
        });

        if (res.status === 401) {
            await refreshAccessToken();
            const newToken = useAuthStore.getState().accessToken;

            res = await fetch(url, {
                ...options,
                credentials: "include",
                headers: {
                    ...(options.headers || {}),
                    Authorization: newToken ? `Bearer ${newToken}` : "",
                },
            });

            if (!res.ok) {
                useAuthStore.getState().setIsLoggedIn(false); // because the refresh token didn't refresh access token successfully
                // alert("please login to use this feature.");
                return undefined;
            }
        }
        return res;
    } catch (error) {
        useAuthStore.getState().clearAccessToken();
        console.error(error);
        return undefined;
    }
};
