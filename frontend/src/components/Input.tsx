interface inputProps {
    labelString: string
    inputType: string
    isRequired: boolean
    inputValue: string
    inputId: string
    inputName: string
    handleChange: (value: string) => void
}
const Input = ({ labelString, inputType, inputName, isRequired, inputValue, inputId, handleChange }: inputProps) => {
    return (
        <>
            <label htmlFor={inputId}>{labelString}</label>
            <input required={isRequired} type={inputType} id={inputId} name={inputName} value={inputValue} onChange={(e) => handleChange(e.target.value)} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
        </>
    )
}

export default Input
