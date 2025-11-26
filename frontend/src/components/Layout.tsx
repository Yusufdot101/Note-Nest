import { Outlet } from "react-router-dom";
import Header from "./Header";
import Footer from "./Footer";

const Layout = () => {
    return (
        <>
            <div className="flex flex-col gap-y-[24px] relative text-text bg-primary min-h-screen bg-primary py-[32px] min-[620px]:text-2xl px-[5vw]">
                <Header />
                <Outlet />
                <Footer />
            </div>
        </>
    );
};

export default Layout;
