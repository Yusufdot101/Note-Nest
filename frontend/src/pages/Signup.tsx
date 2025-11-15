import { useEffect, useState } from "react";
import {
    getEmailErrorMessages,
    getPasswordErrorMessages,
    getUsernameErrorMessages,
} from "../utilities/inputValidation";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { Link, useNavigate } from "react-router-dom";
import { signup } from "../utilities/auth/signup";

const Signup = () => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [showError, setShowError] = useState(false);
    const [showSignupErrors, setSignupShowErrors] = useState(false);

    const [signupErrors, setSignupErrors] = useState<string[]>([]);
    const [usernameError, setUsernameError] = useState("");
    const [emailError, setEmailError] = useState("");
    const [passwordError, setPasswordError] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setShowError(true);
        if (usernameError || emailError || passwordError) {
            return;
        }
        // use the api
        setSignupShowErrors(false);
        setSignupErrors([]);
        const handleErrors = (errors: Record<string, string>) => {
            setSignupShowErrors(true);
            const errorMessages = Object.entries(errors).map(
                ([key, val]) => `${key}: ${val}`,
            );
            setSignupErrors(errorMessages);
        };
        const success = await signup(username, email, password, handleErrors);
        if (!success) return;
        // navigate to the home page when the the account is created
        navigate("/");
    };

    useEffect(() => {
        setUsernameError(getUsernameErrorMessages(username));
    }, [username]);
    useEffect(() => {
        setEmailError(getEmailErrorMessages(email));
    }, [email]);
    useEffect(() => {
        setPasswordError(getPasswordErrorMessages(password));
    }, [password]);

    return (
        <div className="bg-primary flex flex-col w-full border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">
                SIGN UP
            </p>
            <form
                onSubmit={(e) => handleSubmit(e)}
                className="flex flex-col text-text gap-y-[8px]"
            >
                <div className="flex flex-col">
                    <Input
                        labelString={"Username"}
                        inputType={"text"}
                        inputName={"username"}
                        isRequired
                        minLength={2}
                        inputValue={username}
                        inputId={"username"}
                        handleChange={(value) => setUsername(value)}
                    />
                    <p
                        aria-label={"username error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="usernameError"
                    >
                        {usernameError}
                    </p>
                </div>
                <div className="flex flex-col">
                    <Input
                        labelString={"Email"}
                        inputType={"email"}
                        inputName={"email"}
                        isRequired
                        inputValue={email}
                        inputId={"email"}
                        handleChange={(value) => setEmail(value)}
                    />
                    <p
                        aria-label={"email error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="emailError"
                    >
                        {emailError}
                    </p>
                </div>
                <div className="flex flex-col">
                    <Input
                        labelString={"Password"}
                        inputType={"password"}
                        inputName={"password"}
                        isRequired
                        minLength={8}
                        maxLength={72}
                        inputValue={password}
                        inputId={"password"}
                        handleChange={(value) =>
                            setPassword(value.replaceAll(" ", ""))
                        }
                    />
                    <p
                        aria-label={"password error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="passwordError"
                    >
                        {passwordError}
                    </p>
                </div>
                <p>
                    Already have an account?{" "}
                    <Link to={"/login"} className="text-accent">
                        Login here
                    </Link>
                </p>
                <SubmitButton
                    aria_label={"sign up"}
                    handleSubmit={() => {}}
                    text={"Sign Up"}
                />
                <div
                    className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showSignupErrors ? "hidden" : ""}`}
                >
                    {signupErrors.map((error) => (
                        <p key={error}>{error}</p>
                    ))}
                </div>
            </form>
        </div>
    );
};

export default Signup;
