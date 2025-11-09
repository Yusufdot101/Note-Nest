import { useEffect, useState } from "react";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { useNavigate } from "react-router-dom";
import {
    getProjectDescriptionErrorMessages,
    getProjectNameErrorMessages,
    getProjectVisibilityErrorMessages,
} from "../utilities/inputValidation";
import { newProject } from "../utilities/projects";

const NewProject = () => {
    const [projectName, setProjectName] = useState("");
    const [projectDescription, setProjectDescription] = useState("");
    const [projectVisibility, setProjectVisibility] = useState("private");
    const [projectColor, setProjectColor] = useState("#00FFFF");

    const [projectNameError, setProjectNameError] = useState("");
    const [projectDescriptionError, setProjectDescriptionError] = useState("");
    const [projectVisibilityError, setProjectVisibilityError] = useState("");
    const [showError, setShowError] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setShowError(true);
        if (projectNameError || projectVisibilityError) {
            return;
        }
        setShowError(true);
        // use the api
        const success = await newProject(
            projectName,
            projectDescription,
            projectVisibility,
            projectColor,
        );
        if (!success) return;
        // navigate to the home page when the the account is created
        navigate("/projects");
    };

    useEffect(() => {
        setProjectNameError(getProjectNameErrorMessages(projectName));
    }, [projectName]);
    useEffect(() => {
        setProjectDescriptionError(
            getProjectDescriptionErrorMessages(projectDescription),
        );
    }, [projectDescription]);
    useEffect(() => {
        setProjectVisibilityError(
            getProjectVisibilityErrorMessages(projectVisibility),
        );
    }, [projectVisibility]);

    return (
        <div className="bg-primary flex flex-col w-full shadow-[0px_0px_4px_1px_white] py-[32px] min-[620px]:text-2xl px-[12px]">
            <div className="flex items-center justify-center gap-[8px]">
                <p className="text-accent text-[32px] max-[619px]:text-[24px] font-semibold text-center">
                    NEW PROJECT
                </p>
                <div
                    className="relative w-[40px] max-[619px]:w-[35px] h-[30px] max-[619px]:h-[25px] rounded-lg"
                    style={{ backgroundColor: projectColor }}
                >
                    {" "}
                    <input
                        className="inline-block absolute cursor-pointer w-full h-full opacity-0"
                        type="color"
                        required
                        value={projectColor}
                        onChange={(e) => {
                            setProjectColor(e.target.value);
                        }}
                        onInput={() => {}}
                    />
                </div>
            </div>
            <form
                onSubmit={(e) => handleSubmit(e)}
                className="flex flex-col text-text gap-y-[8px]"
            >
                <div className="flex flex-col">
                    <Input
                        labelString={"Project Name"}
                        inputType={"text"}
                        inputName={"project name"}
                        isRequired
                        inputValue={projectName}
                        inputId={"projectName"}
                        handleChange={(value) => setProjectName(value)}
                    />
                </div>
                <div className="flex flex-col">
                    <Input
                        labelString={"Project Description"}
                        inputType={"text"}
                        inputName={"Project Description"}
                        inputValue={projectDescription}
                        inputId={"projectDescription"}
                        handleChange={(value) => setProjectDescription(value)}
                    />
                    <p
                        aria-label={"project description error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectDescriptionError"
                    >
                        {projectDescriptionError}
                    </p>
                </div>
                <div className="flex flex-col">
                    <Input
                        labelString={"Project Visibility"}
                        inputType={"text"}
                        inputName={"Project Visibility"}
                        isRequired
                        inputValue={projectVisibility}
                        inputId={"projectVisibility"}
                        handleChange={(value) => setProjectVisibility(value)}
                    />
                    <p
                        aria-label={"project visibility error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectVisibilityError"
                    >
                        {projectVisibilityError}
                    </p>
                </div>

                <SubmitButton
                    aria_label={"Create Project"}
                    handleSubmit={() => {}}
                    text={"Create Project"}
                />
            </form>
        </div>
    );
};

export default NewProject;
