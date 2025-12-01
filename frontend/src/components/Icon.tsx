import { useNavigate } from "react-router-dom";

const Icon = ({
    src,
    href,
    alt,
    background,
    minWidth,
    height = "50px",
}: {
    src: string;
    href: string;
    alt: string;
    background?: string;
    height?: string;
    minWidth?: string;
}) => {
    const navigate = useNavigate();
    return (
        <div
            style={{
                background: background,
                height: height,
                minWidth: minWidth,
            }}
            className="rounded-[8px] py-[8px] flex items-center justify-center cursor-pointer flex-1"
            onClick={() => {
                navigate(href);
            }}
        >
            <img src={src} alt={alt} className="h-full" />
        </div>
    );
};

export default Icon;
