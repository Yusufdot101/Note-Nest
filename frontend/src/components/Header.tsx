import logo from "./../assets/logo.svg"
import hamburgerMenu from "./../assets/hamburgerMenu.svg"
const Header = () => {
    return (
        <header className="flex justify-between items-center">
            <div className="flex items-center gap-[12px]">
                <img src={logo} alt="logo" className="w-[100px] h-[100px] max-[619px]:w-[70px] max-[619px]:h-[70px]" />
                <span className="text-text font-semibold max-[619px]:text-[20px]  min-[620px]:text-[32px]">NOTE NEST</span>
            </div>
            <img src={hamburgerMenu} alt="hamburger menu" className="w-[50px] h-[50x] max-[619px]:w-[30px] max-[619px]:h-[30px]" />
        </header>
    )
}

export default Header
