import { create } from "zustand";

type AuthState = {
    accessToken: string | null;
    isLoggedIn: boolean;
    userID: number | undefined;
    setAccessToken: (token: string | null) => void;
    setIsLoggedIn: (isLoggedIn: boolean) => void;
    clearAccessToken: () => void;
    setUserID: (userID: number) => void;
};

export const useAuthStore = create<AuthState>((set) => ({
    accessToken: null,
    isLoggedIn: false,
    userID: undefined,
    setAccessToken: (token) => set({ accessToken: token }),
    setIsLoggedIn: (isLoggedIn) => set({ isLoggedIn: isLoggedIn }),
    clearAccessToken: () =>
        set({ accessToken: null, isLoggedIn: false, userID: undefined }),
    setUserID: (userID: number) => set({ userID: userID }),
}));
