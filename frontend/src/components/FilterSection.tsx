type Props = {
    options: Map<string, string | number>;
    handleOptionsChange: (key: string, value: string | number) => void;
};

function FilterSection({ options, handleOptionsChange }: Props) {
    return (
        <div className="text-text bg-primary gap-[8px] flex flex-wrap">
            <div className="flex-1 flex gap-[4px] text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white min-w-[240px] max-[619px]:min-w-[164px]">
                <label htmlFor="visibility">Visibility: </label>
                <select
                    onChange={(e) => {
                        const target = e.target as HTMLSelectElement;
                        handleOptionsChange("visibility", target.value);
                    }}
                    value={options.get("visibility") as string}
                    name="visibility"
                    id="visibility"
                    className="cursor-pointer flex-1"
                >
                    <option value="">All</option>
                    <option value="private">Private</option>
                    <option value="public">Public</option>
                </select>
            </div>

            <div className="flex-1 flex gap-[4px] text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white min-w-[240px] max-[619px]:min-w-[164px]">
                <label htmlFor="sort">Sort: </label>
                <select
                    onChange={(e) => {
                        const target = e.target as HTMLSelectElement;
                        handleOptionsChange("sort", target.value);
                    }}
                    value={options.get("sort") as string}
                    name="sort"
                    id="sort"
                    className="cursor-pointer flex-1"
                >
                    <option value="">Default</option>
                    <option value="likes_count">Likes</option>
                    <option value="comments_count">Comments</option>
                    <option value="created_at">Date</option>
                </select>
            </div>

            <div className="flex-1 flex gap-[4px] text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white min-w-[240px] max-[619px]:min-w-[164px]">
                <label htmlFor="order">Order:</label>
                <select
                    onChange={(e) => {
                        const target = e.target as HTMLSelectElement;
                        handleOptionsChange("order", target.value);
                    }}
                    value={options.get("order") as string}
                    name="order"
                    id="order"
                    className="cursor-pointer flex-1"
                >
                    <option value="">Default</option>
                    <option value="ascending">Ascending</option>
                    <option value="descending">Descending</option>
                </select>
            </div>
        </div>
    );
}

export default FilterSection;
