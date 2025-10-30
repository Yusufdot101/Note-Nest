import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import Signup from './components/Signup.tsx'
import Header from './components/Header.tsx'

const router = createBrowserRouter([
    { path: "/", element: <App /> },
    { path: "/signup", element: <Signup /> },
])

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Header />
        <RouterProvider router={router} >
        </RouterProvider>
    </StrictMode>,
)
