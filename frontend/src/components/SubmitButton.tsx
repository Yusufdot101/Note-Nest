interface submitButtonProps {
    text: string
}
const SubmitButton = ({ text }: submitButtonProps) => {
    return (

        <button className="w-full py-[12px] rounded-[8px] cursor-pointer bg-accent mx-auto">{text}</button>
    )
}

export default SubmitButton
