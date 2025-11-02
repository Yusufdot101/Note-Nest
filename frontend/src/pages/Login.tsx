import { useEffect, useState } from "react"
import { getEmailErrorMessages, getPasswordErrorMessages } from "../utilities/inputValidation"
import Input from "../components/Input"
import SubmitButton from "../components/SubmitButton"
import { handleLogin } from "../utilities/login"

const Login = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const [showError, setShowError] = useState(false)
    const [showLoginErrors, setShowLoginError] = useState(false)

    const [loginError, setLoginError] = useState<string>("")
    const [emailError, setEmailError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setShowError(true)
        if (emailError || passwordError) {
            return
        }
        // use the api
        setShowLoginError(false)
        setLoginError("")
        const handleError = (error: string) => {
            setShowLoginError(true)
            setLoginError(error)
        }
        handleLogin(email, password, handleError)
    }

    useEffect(() => {
        setEmailError(getEmailErrorMessages(email))
    }, [email])
    useEffect(() => {
        setPasswordError(getPasswordErrorMessages(password))
    }, [password])

    return (
        <div className="bg-primary flex flex-col w-full shadow-[0px_0px_4px_1px_white] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">LOGIN</p>
            <form onSubmit={(e) => handleSubmit(e)} className="flex flex-col text-text gap-y-[8px]">
                <div className="flex flex-col">
                    <Input labelString={"Email"} inputType={"email"} inputName={"email"} isRequired inputValue={email} inputId={"email"} handleChange={(value) => setEmail(value)} />
                    <p aria-label={"email error"} className={`text-red-500 ${!showError ? "hidden" : ""}`} id="emailError">{emailError}</p>
                </div>
                <div className="flex flex-col">
                    <Input labelString={"Password"} inputType={"password"} inputName={"password"} isRequired minLength={8} maxLength={72} inputValue={password} inputId={"password"} handleChange={(value) => setPassword(value.replaceAll(" ", ""))} />
                    <p aria-label={"password error"} className={`text-red-500 ${!showError ? "hidden" : ""}`} id="passwordError">{passwordError}</p>
                </div>
                <p>Don't have an account? <a href="/signup" className="text-accent">Register here</a></p>
                <SubmitButton aria_label={"login"} handleSubmit={() => { }} text={"Login"} />
                <div className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showLoginErrors ? "hidden" : ""}`}>
                    {loginError}
                </div>
            </form >
        </div >
    )
}

export default Login
