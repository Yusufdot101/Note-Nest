import { useEffect, useRef } from "react";
import logo from "./../assets/logo.svg";
import { Link, useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/useAuthStore";
import { logout } from "../utilities/auth/logout";
interface SideMenuProps {
    handleClose: () => void;
    menuIsOpen: boolean;
}

const SideMenu = ({ menuIsOpen, handleClose }: SideMenuProps) => {
    const navigate = useNavigate();
    const ref = useRef<HTMLDivElement>(null);
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
    const handleLogout = () => {
        logout();
        navigate("/");
    };

    const navigationLinks = [
        {
            url: "/",
            text: "Home",
        },
        {
            url: "/notes",
            text: "Notes",
        },
        {
            url: "/projects/new",
            text: "New Project",
        },
    ];
    useEffect(() => {
        function handleClick(e: MouseEvent) {
            if (
                ref.current &&
                !ref.current.contains(e.target as Node) &&
                !["svg", "path"].includes((e.target as HTMLElement).tagName)
            ) {
                handleClose();
            }
        }

        document.addEventListener("click", handleClick);
        return () => document.removeEventListener("click", handleClick);
    }, [handleClose]);

    return (
        <div
            ref={ref}
            className={`${menuIsOpen ? "" : "mr-[-100%]"} transition-all duration-300 fixed right-0 top-0 min-w-[250px] w-[40vw] h-screen bg-primary text-text px-[12px] py-[32px] shadow-[0px_0px_4px_1px_white] flex flex-col gap-[12px] z-10`}
        >
            <div className="flex justify-between items-center">
                <Link to={"/"} onClick={handleClose}>
                    <div className="flex items-center gap-[12px] cursor-pointer">
                        <img
                            src={logo}
                            alt="logo"
                            className="w-[70px] h-[70px] max-[619px]:w-[50px] max-[619px]:h-[50px]"
                        />
                        <span className="text-text font-semibold max-[619px]:text-[20px]  min-[620px]:text-[32px] hover:text-accent duration-300">
                            NNest
                        </span>
                    </div>
                </Link>
                <svg
                    viewBox="0 0 96 67"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    className="text-text hover:text-accent cursor-pointer active:text-text duration-300 w-[50px] h-[50px] max-[619px]:w-[30px] max-[619px]:h-[30px]"
                    onClick={handleClose}
                >
                    <path
                        d="M3.5 3.5H92.5M3.5 33.3568H92.5M3.5 63.2137H92.5"
                        stroke="currentColor"
                        strokeWidth="7"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                    />
                </svg>
            </div>
            <ul className="flex flex-col gap-[8px]">
                {navigationLinks.map((link) => (
                    <Link
                        to={link.url}
                        key={`${link.url}-${link.text}`}
                        onClick={handleClose}
                    >
                        <li className="bg-[#747474] p-[8px] hover:text-accent duration-300">
                            {link.text}
                        </li>
                    </Link>
                ))}
                {isLoggedIn ? (
                    <button
                        className="cursor-pointer text-left"
                        onClick={() => {
                            handleLogout();
                            handleClose();
                        }}
                    >
                        <li className="bg-[#747474] p-[8px] hover:text-accent duration-300">
                            Logout
                        </li>
                    </button>
                ) : (
                    <Link to={"/login"} onClick={handleClose}>
                        <li className="bg-[#747474] p-[8px] hover:text-accent duration-300">
                            Login
                        </li>
                    </Link>
                )}
            </ul>
        </div>
    );
};

export default SideMenu;
