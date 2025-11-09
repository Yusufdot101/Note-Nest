import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Layout from "./components/Layout";
import Signup from "./pages/Signup";
import Login from "./pages/Login";
import Home from "./pages/Home";
import { useEffect } from "react";
import { initAuth } from "./utilities/auth/initAuth";
import NewProject from "./pages/NewProject";
import Projects from "./pages/Projects";

const router = createBrowserRouter([
    {
        element: <Layout />,
        children: [
            { path: "/", element: <Home /> },
            { path: "/signup", element: <Signup /> },
            { path: "/login", element: <Login /> },
            { path: "/projects/new", element: <NewProject /> },
            { path: "/projects", element: <Projects /> },
        ],
    },
]);

const App = () => {
    useEffect(() => {
        initAuth();
    }, []);
    return (
        <>
            <RouterProvider router={router} />
        </>
    );
};

export default App;
