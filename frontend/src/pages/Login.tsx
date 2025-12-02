import google from "../assets/google.svg";
import { useEffect, useState } from "react";
import {
    getEmailErrorMessages,
    getPasswordErrorMessages,
} from "../utilities/inputValidation";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { Link, useNavigate } from "react-router-dom";
import { login } from "../utilities/auth/login";
import Icon from "../components/Icon";

const continueWithOptions = [
    {
        name: "google",
        imgSrc: google,
        href: "http://127.0.0.1:8080/auth/start/google",
    },
];

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const [showError, setShowError] = useState(false);
    const [showLoginErrors, setShowLoginError] = useState(false);

    const [loginError, setLoginError] = useState<string[]>([]);
    const [emailError, setEmailError] = useState("");
    const [passwordError, setPasswordError] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setShowError(true);
        if (emailError || passwordError) {
            return;
        }
        // use the api
        setShowLoginError(false);
        setLoginError([]);

        const handleErrors = (errors: Record<string, string>) => {
            setShowLoginError(true);
            const errorMessages = Object.entries(errors).map(
                ([key, val]) => `${key}: ${val}`,
            );
            setLoginError(errorMessages);
        };

        const success = await login(email, password, handleErrors);
        if (!success) return;

        // navigate to the home page when the the account is created
        navigate("/");
    };

    useEffect(() => {
        setEmailError(getEmailErrorMessages(email.trim()));
    }, [email]);
    useEffect(() => {
        setPasswordError(getPasswordErrorMessages(password.trim()));
    }, [password]);

    return (
        <div className="w-full bg-primary flex flex-col gap-[24px] border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">
                LOGIN
            </p>
            <form
                onSubmit={(e) => handleSubmit(e)}
                className="flex flex-col text-text gap-y-[8px]"
            >
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
                        handleChange={(value) => {
                            setPassword(value.trimStart());
                        }}
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
                    <Link to={"/forgot-password"} className="text-accent">
                        Forgot Password?
                    </Link>
                </p>
                <SubmitButton
                    aria_label={"login"}
                    handleSubmit={() => {}}
                    text={"Login"}
                    type="submit"
                />
                <p>
                    Don't have an account?{" "}
                    <Link to={"/signup"} className="text-accent">
                        Register here
                    </Link>
                </p>
                <div
                    className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showLoginErrors ? "hidden" : ""}`}
                >
                    {loginError.map((error) => (
                        <p key={error}>{error}</p>
                    ))}
                </div>
            </form>

            <div className="relative">
                <hr className="text-accent border-[1px] opacity-75" />
                <p className="absolute top-[50%] left-[50%] -translate-x-[50%] -translate-y-[50%] bg-primary px-[12px] w-fit text-nowrap">
                    or continue with
                </p>
            </div>

            <div className="flex flex-col gap-[12px]">
                <div className="flex gap-[8px] flex-wrap h-fit">
                    {continueWithOptions.map((option) => (
                        <Icon
                            key={option.href}
                            src={option.imgSrc}
                            href={option.href}
                            alt={`continue with ${option.name}`}
                            background="white"
                            minWidth="100px"
                        />
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Login;
