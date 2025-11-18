import searchIcon from "../assets/searchIcon.svg";
import Input from "./Input";

interface SearchBarProps {
    searchValue: string;
    setSearchValue: React.Dispatch<React.SetStateAction<string>>;
    placeholder: string;
    handleSearch: () => void;
}

const SearchBar = ({
    searchValue,
    setSearchValue,
    placeholder,
    handleSearch,
}: SearchBarProps) => {
    return (
        <div className="w-full flex gap-[8px] h-[50px]">
            <Input
                ariaLabel="search value"
                inputType="text"
                inputId="searchValue"
                inputName="searchValue"
                placeholder={placeholder}
                inputValue={searchValue}
                handleChange={(value: string) => setSearchValue(value)}
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
                src={searchIcon}
                alt="search icon"
                className="h-full cursor-pointer"
            />
        </div>
    );
};

export default SearchBar;
