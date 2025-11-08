const Icon = ({
    src,
    href,
    alt,
}: {
    src: string;
    href: string;
    alt: string;
}) => {
    return (
        <div>
            <a href={href} target="_blank">
                <img
                    src={src}
                    alt={alt}
                    className="w-[75px] h-[75px] max-[619px]:w-[70px] max-[619px]:h-[70px]"
                />
            </a>
        </div>
    );
};

export default Icon;
