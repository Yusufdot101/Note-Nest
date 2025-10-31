import { useEffect, useState } from "react"
import { getEmailErrorMessages, getPasswordErrorMessages, getUsernameErrorMessages } from "../utilities/inputValidation"
import Input from "../components/Input"
import SubmitButton from "../components/SubmitButton"

const Signup = () => {
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const [showError, setShowError] = useState(false)

    const [usernameError, setUsernameError] = useState("")
    const [emailError, setEmailError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setShowError(true)
        if (usernameError || emailError || passwordError) {
            return
        }
        // use the api
    }

    useEffect(() => {
        setUsernameError(getUsernameErrorMessages(username))
    }, [username])
    useEffect(() => {
        setEmailError(getEmailErrorMessages(email))
    }, [email])
    useEffect(() => {
        setPasswordError(getPasswordErrorMessages(password))
    }, [password])

    return (
        <div className="bg-primary flex flex-col mx-uto w-full shadow-[0px_0px_4px_1px_white] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">SIGN UP</p>
            <form onSubmit={(e) => handleSubmit(e)} className="flex flex-col text-text gap-y-[8px]">
                <div className="flex flex-col">
                    <Input lableStrig={"Username"} inputType={"text"} isRequired inputValue={username} inputId={"username"} handleChange={(value) => setUsername(value)} />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="usernameError">{usernameError}</p>
                </div>
                <div className="flex flex-col">
                    <Input lableStrig={"Email"} inputType={"email"} isRequired inputValue={email} inputId={"email"} handleChange={(value) => setEmail(value)} />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="emailError">{emailError}</p>
                </div>
                <div className="flex flex-col">
                    <Input lableStrig={"Password"} inputType={"password"} isRequired inputValue={password} inputId={"password"} handleChange={(value) => setPassword(value.replaceAll(" ", ""))} />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="passwordError">{passwordError}</p>
                </div>
                <p>Already have an account? <a href="#" className="text-accent">Login here</a></p>
                <SubmitButton text={"Sign Up"} />
            </form >
        </div >
    )
}

export default Signup
