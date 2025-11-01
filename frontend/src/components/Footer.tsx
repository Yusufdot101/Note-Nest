import githubLogo from "./../assets/github.svg"
import xLogo from "./../assets/X.svg"
import linkedInLogo from "./../assets/linkedIn.svg"
import Icon from "./Icon"

const Footer = () => {
    return (
        <div className="bg-primary shadow-[0px_0px_4px_1px_white] py-[12px] flex flex-col items-center text-text gap-[8px]">
            <div className="flex justify-center gap-x-[8px]">
                <Icon src={githubLogo} href={"https://github.com/Yusufdot101"} alt={"GitHub logo"} />
                <Icon src={xLogo} href={""} alt={"X logo"} />
                <Icon src={linkedInLogo} href={""} alt={"LinkedIn logo"} />
            </div>
            <div className="w-full text-center">
                <p className="max-[619px]:text-[0.75rem]">Email: yusuf.mohamed.work@gmail.com</p>
                <p className="max-[619px]:text-[0.75rem]">COPYRIGHT © 2025 by Yusuf Mohamed</p>
            </div>
        </div>
    )
}

export default Footer
