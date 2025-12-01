import githubLogo from "./../assets/github.svg";
import xLogo from "./../assets/X.svg";
import linkedInLogo from "./../assets/linkedIn.svg";
import Icon from "./Icon";

const Footer = () => {
    return (
        <div className="bg-primary border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[12px] w-full min-w-[300px] flex flex-col items-center text-text gap-[8px]">
            <div className="flex h-[80px] max-[619px]:h-[72px] flex-wrap justify-center gap-x-[8px]">
                <Icon
                    src={githubLogo}
                    href={"https://github.com/Yusufdot101"}
                    alt={"GitHub logo"}
                    height="100%"
                />
                <Icon src={xLogo} href={""} alt={"X logo"} height="100%" />
                <Icon
                    src={linkedInLogo}
                    href={""}
                    alt={"LinkedIn logo"}
                    height="100%"
                />
            </div>
            <div className="w-full text-center">
                <p className="text-[20px] max-[619px]:text-[12px]">
                    Email: yusuf.mohamed.work@gmail.com
                </p>
                <p className="text-[20px] max-[619px]:text-[12px]">
                    COPYRIGHT © 2025 by Yusuf Mohamed
                </p>
            </div>
        </div>
    );
};

export default Footer;
