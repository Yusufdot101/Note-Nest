import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
    plugins: [react(), tailwindcss()],
    test: {
        environment: "jsdom",
        globals: true,
        setupFiles: "./src/setupTests.ts",
    },
    server: {
        host: true, // binds to 0.0.0.0
        port: 3000,
        allowedHosts: [
            "archlinux.local", // hostname
            "localhost",
        ],
    },
});
