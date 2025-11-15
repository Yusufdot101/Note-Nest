import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Layout from "./components/Layout";
import Signup from "./pages/Signup";
import Login from "./pages/Login";
import Home from "./pages/Home";
import { useEffect } from "react";
import { initAuth } from "./utilities/auth/initAuth";
import NewProject from "./pages/NewProject";
import AllProjects from "./pages/AllProjects";
import ProjectPage from "./pages/Project";
import EditProject from "./pages/EditProject";
import NoteCreation from "./pages/NoteCreation";

const router = createBrowserRouter([
    {
        element: <Layout />,
        children: [
            { path: "/", element: <Home /> },
            { path: "/signup", element: <Signup /> },
            { path: "/login", element: <Login /> },
            { path: "/projects/new", element: <NewProject /> },
            { path: "/projects", element: <AllProjects /> },
            { path: "/projects/:id", element: <ProjectPage /> },
            { path: "/projects/:id/edit", element: <EditProject /> },
            {
                path: "/projects/:projectid/notes/new",
                element: <NoteCreation />,
            },
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
