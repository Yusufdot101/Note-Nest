interface LikeButtonProps {
    onToggle: () => Promise<void>;
    liked: boolean;
    count: number;
}

const LikeButton = ({ onToggle, liked, count }: LikeButtonProps) => {
    return (
        <div
            onClick={onToggle}
            onKeyDown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                    e.preventDefault();
                    onToggle();
                }
            }}
            role="button"
            aria-label="like note"
            tabIndex={0}
            className="flex items-center gap-x-[4px] cursor-pointer"
        >
            <svg
                width="28"
                height="28"
                viewBox="0 0 24 24"
                fill={`${liked ? "white" : "none"}`}
                xmlns="http://www.w3.org/2000/svg"
            >
                <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                <g
                    id="SVGRepo_tracerCarrier"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                ></g>
                <g id="SVGRepo_iconCarrier">
                    {" "}
                    <path
                        d="M8 10V20M8 10L4 9.99998V20L8 20M8 10L13.1956 3.93847C13.6886 3.3633 14.4642 3.11604 15.1992 3.29977L15.2467 3.31166C16.5885 3.64711 17.1929 5.21057 16.4258 6.36135L14 9.99998H18.5604C19.8225 9.99998 20.7691 11.1546 20.5216 12.3922L19.3216 18.3922C19.1346 19.3271 18.3138 20 17.3604 20L8 20"
                        stroke="#ffffff"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                    ></path>{" "}
                </g>
            </svg>
            <span>{count}</span>
        </div>
    );
};

export default LikeButton;
