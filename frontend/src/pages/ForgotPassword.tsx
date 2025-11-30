import { useEffect, useState } from "react";
import { getEmailErrorMessages } from "../utilities/inputValidation";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { useNavigate } from "react-router-dom";
import { forgotPassword } from "../utilities/auth/forgotPassword";

const ForgotPassword = () => {
    const [email, setEmail] = useState("");

    const [showError, setShowError] = useState(false);
    const [showForgotPasswordErrors, setShowForgotPasswordError] =
        useState(false);

    const [forgotPasswordError, setForgotPasswordError] = useState<string[]>(
        [],
    );
    const [emailError, setEmailError] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setShowError(true);
        if (emailError) {
            return;
        }
        // use the api
        setShowForgotPasswordError(false);
        setForgotPasswordError([]);

        const handleErrors = (errors: Record<string, string>) => {
            setShowForgotPasswordError(true);
            const errorMessages = Object.entries(errors).map(
                ([key, val]) => `${key}: ${val}`,
            );
            setForgotPasswordError(errorMessages);
        };

        const message = await forgotPassword(email, handleErrors);
        if (!message) return;
        alert(message);
        // navigate to the home page when the the account is created
        navigate("/login");
    };

    useEffect(() => {
        setEmailError(getEmailErrorMessages(email.trim()));
    }, [email]);

    return (
        <div className="w-full bg-primary flex flex-col border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">
                Forgot Password
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
                <SubmitButton
                    aria_label={"send password reset link"}
                    handleSubmit={() => {}}
                    text={"Send Reset Link"}
                    type="submit"
                />
                <div
                    className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showForgotPasswordErrors ? "hidden" : ""}`}
                >
                    {forgotPasswordError.map((error, index) => (
                        <p key={index}>{error}</p>
                    ))}
                </div>
            </form>
        </div>
    );
};

export default ForgotPassword;
