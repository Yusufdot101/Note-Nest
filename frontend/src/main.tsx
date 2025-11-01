import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import Signup from './pages/Signup.tsx'
import Layout from './components/Layout.tsx'

const router = createBrowserRouter([
    {
        element: <Layout />,
        children: [
            { path: "/", element: <App /> },
            { path: "/signup", element: <Signup /> },
        ]
    }
])

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <RouterProvider router={router} />
    </StrictMode>,
)
