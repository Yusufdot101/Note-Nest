import { createBrowserRouter, RouterProvider } from "react-router-dom"
import Layout from "./components/Layout"
import Signup from "./pages/Signup"
import Login from "./pages/Login"
import Home from "./pages/Home"
import { useEffect } from "react"
import { initAuth } from "./utilities/auth/initAuth"

const router = createBrowserRouter([
    {
        element: <Layout />,
        children: [
            { path: "/", element: <Home /> },
            { path: "/signup", element: <Signup /> },
            { path: "/login", element: <Login /> },
        ]
    }
])

const App = () => {
    useEffect(() => {
        initAuth()
    }, [])
    return (
        <>
            <RouterProvider router={router} />
        </>
    )
}

export default App
