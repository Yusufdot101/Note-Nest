type Props = {
    handleClick: (
        e: React.MouseEvent<SVGElement> | React.KeyboardEvent<SVGElement>,
    ) => void;
    userId: number | undefined;
    projectUserId: number | undefined;
};

const Menu = ({ handleClick, userId, projectUserId }: Props) => {
    return (
        <span>
            <svg
                role="button"
                aria-label="open note actions menu"
                tabIndex={0}
                onKeyDown={(e) => {
                    if (e.key === "Enter" || e.key === " ") {
                        handleClick(e);
                    }
                }}
                fill="currentColor"
                version="1.1"
                id="Icons"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 32 32"
                className={`${userId === projectUserId ? "" : "hidden"} w-[30px] h-[30px] hover:text-accent active:text-text duration-300`}
                onClick={(e) => {
                    handleClick(e);
                }}
            >
                <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                <g
                    id="SVGRepo_tracerCarrier"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                ></g>
                <g id="SVGRepo_iconCarrier">
                    {" "}
                    <g>
                        {" "}
                        <path d="M16,10c1.7,0,3-1.3,3-3s-1.3-3-3-3s-3,1.3-3,3S14.3,10,16,10z"></path>{" "}
                        <path d="M16,13c-1.7,0-3,1.3-3,3s1.3,3,3,3s3-1.3,3-3S17.7,13,16,13z"></path>{" "}
                        <path d="M16,22c-1.7,0-3,1.3-3,3s1.3,3,3,3s3-1.3,3-3S17.7,22,16,22z"></path>{" "}
                    </g>{" "}
                </g>
            </svg>
        </span>
    );
};

export default Menu;
