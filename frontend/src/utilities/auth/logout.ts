import { useAuthStore } from "../../store/useAuthStore";
import { BASE_APIURL } from "../api";

export const logout = async () => {
  try {
    const res = await fetch(`${BASE_APIURL}/auth/logout`, {
      credentials: "include",
      method: "PUT",
    });
    if (!res.ok) {
      throw new Error(`Logout failed: ${res.status} ${res.statusText}`);
    }
    useAuthStore.getState().setIsLoggedIn(false);
    useAuthStore.getState().setAccessToken(null);
  } catch (error) {
    alert("an error occurred, please try again");
    console.error(error);
  }
};
