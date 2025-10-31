interface inputProps {
    lableStrig: string
    inputType: string
    isRequired: boolean
    inputValue: string
    inputId: string
    handleChange: (value: string) => void
}
const Input = ({ lableStrig, inputType, isRequired, inputValue, inputId, handleChange }: inputProps) => {
    return (
        <>
            <label htmlFor={inputId}>{lableStrig}</label>
            <input required={isRequired} type={inputType} id={inputId} name="username" value={inputValue} onChange={(e) => handleChange(e.target.value)} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
        </>
    )
}

export default Input
