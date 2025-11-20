import searchIcon from "../assets/searchIcon.svg";
import Input from "./Input";

interface SearchBarProps {
    searchValue: string;
    handleValueChange: (value: string) => void;
    placeholder: string;
    handleSearch: () => void;
}

const SearchBar = ({
    searchValue,
    handleValueChange,
    placeholder,
    handleSearch,
}: SearchBarProps) => {
    return (
        <form
            onSubmit={(e) => {
                e.preventDefault();
                handleSearch();
            }}
            className="w-full flex gap-[4px] h-[50px]"
        >
            <Input
                ariaLabel="search value"
                inputType="text"
                inputId="searchValue"
                inputName="searchValue"
                placeholder={placeholder}
                inputValue={searchValue}
                handleChange={handleValueChange}
            />
            <img
                role="button"
                tabIndex={0}
                aria-label="search"
                onKeyDown={(e) => {
                    if (e.key === "Enter" || e.key === " ") {
                        handleSearch();
                    }
                }}
                onClick={(e) => {
                    e.preventDefault();
                    handleSearch();
                }}
                src={searchIcon}
                alt="search icon"
                className="h-full cursor-pointer"
            />
        </form>
    );
};

export default SearchBar;
