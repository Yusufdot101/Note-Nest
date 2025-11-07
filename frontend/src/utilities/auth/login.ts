import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";

export const login = async (email: string, password: string, handleErrors: (error: string) => void): Promise<boolean> => {
    try {
        const res = await fetch(`${BASE_APIURL}/auth/login`, {
            method: "PUT",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: email, password: password })
        })

        const data = await res.json()
        if (!res.ok) {
            const error = data.error
            if (error) {
                handleErrors(error)
                return false
            }
            throw new Error(`HTTP error! status: ${res.status}`)
        }
        useAuthStore.getState().setAccessToken(data.token)
        useAuthStore.getState().setIsLoggedIn(true)
        return true
    } catch (error) {
        alert("an error occurred, please try again")
        console.error(error)
        return false
    }
}
