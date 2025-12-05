import type { Metadata } from "../utilities/projects";

type Props = {
    options: Map<string, string | number>;
    handleOptionsChange: (key: string, value: string | number) => void;
    metadata: Metadata;
};

function PageNumbers({ options, handleOptionsChange, metadata }: Props) {
    const range = Array.from(
        { length: metadata?.TotalPages ?? 1 },
        (_, i) => i + 1,
    );
    const currentPageNumber = options.get("page") as number;
    const lastPage = metadata?.TotalPages;
    return (
        <div className="flex mx-auto gap-[4px]">
            {range.map((pageNumber) => {
                // numbers we show fully (1, last, near current)
                const isShown =
                    Math.abs(currentPageNumber - pageNumber) <= 1 ||
                    [1, lastPage].includes(pageNumber);

                // numbers that get "." instead
                const isEllipse = Math.abs(currentPageNumber - pageNumber) <= 4;
                const isCurrentPage = pageNumber === currentPageNumber;

                return (
                    <span
                        key={pageNumber}
                        className={`${isCurrentPage ? "text-accent mt-[-8px]" : ""} ${isEllipse ? "" : isShown ? "" : "hidden"} cursor-pointer hover:text-accent duration-300`}
                        onClick={() => {
                            if (isCurrentPage) return;
                            handleOptionsChange("page", pageNumber);
                        }}
                        tabIndex={0}
                        role="button"
                        aria-label={`go to page ${pageNumber}`}
                        onKeyDown={(e) => {
                            if (e.key === "Enter" || e.key === " ") {
                                e.preventDefault();
                                if (isCurrentPage) return;
                                handleOptionsChange("page", pageNumber);
                            }
                        }}
                    >
                        {isShown ? pageNumber : isEllipse ? "." : ""}
                    </span>
                );
            })}
        </div>
    );
}

export default PageNumbers;
