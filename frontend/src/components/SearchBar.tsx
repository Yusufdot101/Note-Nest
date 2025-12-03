import searchIcon from "../assets/searchIcon.svg";
import filterIcon from "../assets/filterIcon.svg";
import Input from "./Input";
import { useState } from "react";
import FilterSection from "./FilterSection";

interface SearchBarProps {
    options: Map<string, number | string>;
    handleOptionsChange: (key: string, value: string | number) => void;
    searchPlaceholder: string;
    handleSearch: () => void;
}

const SearchBar = ({
    options,
    handleOptionsChange,
    searchPlaceholder,
    handleSearch,
}: SearchBarProps) => {
    const [showFilterSection, setShowFilterSection] = useState(false);
    return (
        <form
            onSubmit={(e) => {
                e.preventDefault();
                handleSearch();
            }}
            className="w-full flex flex-col gap-[4px]"
        >
            <div className="flex gap-[4px] h-[50px]">
                <img
                    role="button"
                    tabIndex={0}
                    aria-label="search"
                    onKeyDown={(e) => {
                        if (e.key === "Enter" || e.key === " ") {
                            setShowFilterSection((prev) => !prev);
                        }
                    }}
                    onClick={(e) => {
                        e.preventDefault();
                        setShowFilterSection((prev) => !prev);
                    }}
                    src={filterIcon}
                    alt="search icon"
                    className="h-full max-[619px]:w-[40px] cursor-pointer bg-white rounded-[8px] p-[4px]"
                />
                <Input
                    ariaLabel="title"
                    inputType="text"
                    inputId="title"
                    inputName="title"
                    placeholder={searchPlaceholder}
                    inputValue={options.get("title") as string}
                    handleChange={(value) => {
                        handleOptionsChange("title", value);
                    }}
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
                    className="h-full max-[619px]:w-[40px] cursor-pointer bg-white rounded-[8px]"
                />
            </div>

            {showFilterSection && (
                <FilterSection
                    options={options}
                    handleOptionsChange={handleOptionsChange}
                />
            )}
        </form>
    );
};

export default SearchBar;
