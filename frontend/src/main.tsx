import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import Signup from './components/Signup.tsx'
import Header from './components/Header.tsx'
import Footer from './components/Footer.tsx'

const router = createBrowserRouter([
    { path: "/", element: <App /> },
    { path: "/signup", element: <Signup /> },
])

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <div className="bg-background min-h-screen flex flex-col p-[32px] gap-[32px]">
            <Header />
            <RouterProvider router={router} >
            </RouterProvider>
            <Footer />
        </div>
    </StrictMode>,
)
