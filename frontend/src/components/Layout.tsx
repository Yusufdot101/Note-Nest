import { Outlet } from "react-router-dom"
import Header from "./Header"
import Footer from "./Footer"

const Layout = () => {
    return (
        <>
            <div className="bg-background min-h-screen flex flex-col p-[32px] gap-[32px]">
                <Header />
                <Outlet />
                <Footer />
            </div>
        </>
    )
}

export default Layout
