import { useEffect, useState } from "react"
import { getEmailErrorMessages, getPasswordErrorMessages, getUsernameErrorMessages } from "../utilities/inputValidation"

const Signup = () => {
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const [showError, setShowError] = useState(false)

    const [usernameError, setUsernameError] = useState("")
    const [emailError, setEmailError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    const handlerSumbit = (e: React.FormEvent<HTMLFormElement>) => {
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
            <form onSubmit={(e) => handlerSumbit(e)} className="flex flex-col text-text gap-y-[8px]">
                <div className="flex flex-col">
                    <label htmlFor="username">Username</label>
                    <input required type="text" id="username" name="username" value={username} onChange={(e) => setUsername(e.target.value)} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="usernameError">{usernameError}</p>
                </div>
                <div className="flex flex-col">
                    <label htmlFor="email">Email</label>
                    <input required type="email" id="email" name="email" value={email} onChange={(e) => setEmail(e.target.value.trim())} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="emailError">{emailError}</p>
                </div>
                <div className="flex flex-col">
                    <label htmlFor="password">Password</label>
                    <input required type="password" id="password" name="password" value={password} onChange={(e) => setPassword(e.target.value.replaceAll(" ", ""))} className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                    <p className={`text-red-500 ${!showError ? "hidden" : ""}`} id="passwordError">{passwordError}</p>
                </div>
                <p>Already have an account? <a href="#" className="text-accent">Login here</a></p>
                <button className="w-full py-[12px] rounded-[8px] cursor-pointer bg-accent mx-auto">Sign Up</button>
            </form >
        </div >
    )
}

export default Signup
