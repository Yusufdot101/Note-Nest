import { useEffect, useState } from "react";
import {
    getConfirmPasswordErrorMessages,
    getPasswordErrorMessages,
} from "../utilities/inputValidation";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { useNavigate } from "react-router-dom";
import { resetPassword } from "../utilities/auth/forgotPassword";

const ResetPassword = () => {
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const [showError, setShowError] = useState(false);
    const [showResetPasswordErrors, setShowResetPasswordError] =
        useState(false);

    const [resetPasswordError, setResetPasswordError] = useState<string[]>([]);
    const [passwordError, setPasswordError] = useState("");
    const [confirmPasswordError, setConfirmPasswordError] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setShowError(true);
        if (passwordError || confirmPasswordError) {
            return;
        }
        // use the api
        setShowResetPasswordError(false);
        setResetPasswordError([]);

        const handleErrors = (errors: Record<string, string>) => {
            setShowResetPasswordError(true);
            const errorMessages = Object.entries(errors).map(
                ([key, val]) => `${key}: ${val}`,
            );
            setResetPasswordError(errorMessages);
        };

        const message = await resetPassword(password, handleErrors);
        if (!message) return;
        alert(message);
        // // navigate to the home page when the the account is created
        navigate("/login");
    };

    useEffect(() => {
        setPasswordError(getPasswordErrorMessages(password.trim()));
    }, [password]);
    useEffect(() => {
        setConfirmPasswordError(
            getConfirmPasswordErrorMessages(
                password.trim(),
                confirmPassword.trim(),
            ),
        );
    }, [password, confirmPassword]);

    return (
        <div className="w-full bg-primary flex flex-col border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">
                Reset Password
            </p>
            <form
                onSubmit={(e) => handleSubmit(e)}
                className="flex flex-col text-text gap-y-[8px]"
            >
                <div className="flex flex-col">
                    <Input
                        labelString={"New Password"}
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

                <div className="flex flex-col">
                    <Input
                        labelString={"Confirm Password"}
                        inputType={"password"}
                        inputName={"confirmPassword"}
                        isRequired
                        minLength={8}
                        maxLength={72}
                        inputValue={confirmPassword}
                        inputId={"confirmPassword"}
                        handleChange={(value) => {
                            setConfirmPassword(value.trimStart());
                        }}
                    />
                    <p
                        aria-label={"confirm password error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="confirmPasswordError"
                    >
                        {confirmPasswordError}
                    </p>
                </div>

                <SubmitButton
                    aria_label={"reset password"}
                    handleSubmit={() => {}}
                    text={"Reset Password"}
                    type="submit"
                />

                <div
                    className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showResetPasswordErrors ? "hidden" : ""}`}
                >
                    {resetPasswordError.map((error) => (
                        <p key={error}>{error}</p>
                    ))}
                </div>
            </form>
        </div>
    );
};

export default ResetPassword;
