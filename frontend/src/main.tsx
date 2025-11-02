import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import Signup from './pages/Signup.tsx'
import Layout from './components/Layout.tsx'
import Login from './pages/Login.tsx'

const router = createBrowserRouter([
    {
        element: <Layout />,
        children: [
            { path: "/", element: <App /> },
            { path: "/signup", element: <Signup /> },
            { path: "/login", element: <Login /> },
        ]
    }
])

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <RouterProvider router={router} />
    </StrictMode>,
)
