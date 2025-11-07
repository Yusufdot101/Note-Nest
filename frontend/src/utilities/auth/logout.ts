import { useAuthStore } from "../../store/useAuthStore"
import { BASE_APIURL } from "../api"

export const logout = async () => {
    try {
        await fetch(`${BASE_APIURL}/auth/logout`, {
            credentials: "include",
            method: "PUT",
        })

        useAuthStore.getState().setIsLoggedIn(false)
        useAuthStore.getState().setAccessToken(null)
    } catch (error) {
        alert("an error occurred, please try again")
        console.error(error)
    }
}
