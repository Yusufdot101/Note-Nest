import { Link } from "react-router-dom";
import logo from "./../assets/logo.svg";
import SideMenu from "./SideMenu";
import { useState } from "react";
const Header = () => {
  const [menuIsOpen, setMenuIsOpen] = useState(false);
  const handleClose = () => setMenuIsOpen(false);
  return (
    <header className="flex justify-between items-center">
      <Link to={"/"}>
        <div className="flex items-center gap-[12px] cursor-pointer">
          <img
            src={logo}
            alt="logo"
            className="w-[100px] h-[100px] max-[619px]:w-[70px] max-[619px]:h-[70px]"
          />
          <span className="text-text font-semibold max-[619px]:text-[20px]  min-[620px]:text-[32px] hover:text-accent duration-300">
            NOTE NEST
          </span>
        </div>
      </Link>

      <svg
        viewBox="0 0 96 67"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        className="text-white hover:text-accent cursor-pointer active:text-white duration-300 w-[50px] h-[50px] max-[619px]:w-[30px] max-[619px]:h-[30px]"
        onClick={() => setMenuIsOpen((prev) => !prev)}
      >
        <path
          d="M3.5 3.5H92.5M3.5 33.3568H92.5M3.5 63.2137H92.5"
          stroke="currentColor"
          strokeWidth="7"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
      </svg>
      <SideMenu menuIsOpen={menuIsOpen} handleClose={handleClose} />
    </header>
  );
};

export default Header;
