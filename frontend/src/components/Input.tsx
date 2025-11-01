interface inputProps {
    labelString: string
    inputType: string
    isRequired: boolean
    minLength?: number
    maxLength?: number
    inputValue: string
    inputId: string
    inputName: string
    handleChange: (value: string) => void
}
const Input = ({ labelString, inputType, inputName, isRequired, minLength, maxLength, inputValue, inputId, handleChange }: inputProps) => {
    return (
        <>
            <label htmlFor={inputId}>{labelString}</label>
            <input required={isRequired} minLength={minLength} maxLength={maxLength} type={inputType} id={inputId} name={inputName} value={inputValue} onChange={(e) => handleChange(e.target.value)} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
        </>
    )
}

export default Input
