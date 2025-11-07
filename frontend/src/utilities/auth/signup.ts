import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";

export const signup = async (username: string, email: string, password: string, handleErrors: (errors: Record<string, string>) => void): Promise<boolean> => {
    try {
        const res = await fetch(`${BASE_APIURL}/auth/signup`, {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name: username, email: email, password: password })
        })

        const data = await res.json()
        if (!res.ok) {
            if (!res.ok) {
                const errors = data.error
                if (errors) {
                    handleErrors(errors)
                    return false
                }
                throw new Error(`HTTP error! status: ${res.status}`)
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
