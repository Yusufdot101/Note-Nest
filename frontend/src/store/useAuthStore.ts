import { create } from "zustand";

type AuthState = {
  accessToken: string | null;
  isLoggedIn: boolean;
  setAccessToken: (token: string | null) => void;
  setIsLoggedIn: (isLoggedIn: boolean) => void;
  clearAccessToken: () => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  accessToken: null,
  isLoggedIn: false,
  setAccessToken: (token) => set({ accessToken: token }),
  setIsLoggedIn: (isLoggedIn) => set({ isLoggedIn: isLoggedIn }),
  clearAccessToken: () => set({ accessToken: null, isLoggedIn: false }),
}));
