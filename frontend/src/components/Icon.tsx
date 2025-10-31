const Icon = ({ src, href }: { src: string, href: string }) => {
    return (
        <div>
            <a href={href} target="_blank"><img src={src} alt="github logo" className="w-[75px] h-[75px] max-[619px]:w-[70px] max-[619px]:h-[70px]" /></a>
        </div>
    )
}

export default Icon
