import githubLogo from "./../assets/github.svg"
import xLogo from "./../assets/X.svg"
import linkedInLogo from "./../assets/linkedIn.svg"

const Footer = () => {
    return (
        <div className="bg-primary shadow-[0px_0px_4px_1px_white] py-[12px] flex flex-col items-center text-text gap-[8px]">
            <div className="flex justify-center gap-x-[8px]">
                <a href="https://github.com/Yusufdot101" target="_blank"><img src={githubLogo} alt="github logo" className="w-[75px] h-[75px] max-[619px]:w-[70px] max-[619px]:h-[70px]" /></a>
                <a href="#"><img src={xLogo} alt="x logo" className="w-[75px] h-[75px] max-[619px]:w-[70px] max-[619px]:h-[70px]" /></a>
                <a href="#"><img src={linkedInLogo} alt="linkedin logo" className="w-[75px] h-[75px] max-[619px]:w-[70px] max-[619px]:h-[70px]" /></a>
            </div>
            <div className="w-full text-center">
                <p className="max-[619px]:text-[0.75rem]">Email: yusuf.mohamed.work@gmail.com</p>
                <p className="max-[619px]:text-[0.75rem]">COPYRIGHT © 2025 by Yusuf Mohamed</p>
            </div>
        </div>
    )
}

export default Footer
