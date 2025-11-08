import { useEffect, useState } from "react";
proje;
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
    const [projectVisibility, setProjectVisibility] = useState("");

    const [showNewProjectErrors, setShowNewProjectErrors] = useState(false);
    const [newProjectErrors, setNewProjectErrors] = useState<string[]>([]);
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
        const handleErrors = (errors: Record<string, string>) => {
            setShowNewProjectErrors(true);
            const errorMessages = Object.entries(errors).map(
                ([key, val]) => `${key}: ${val}`,
            );
            setNewProjectErrors(errorMessages);
        };
        const success = await newProject(
            projectName,
            projectDescription,
            projectVisibility,
            handleErrors,
        );
        if (!success) return;
        // navigate to the home page when the the account is created
        navigate("/");
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
            <p className="text-accent text-[32px] font-semibold text-center">
                CREATE NEW PROJECT
            </p>
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
                    <p
                        aria-label={"project name error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectNameError"
                    >
                        {projectNameError}
                    </p>
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

                <div
                    className={`w-full text-center py-[12px] rounded-[8px] bg-red-500 mx-auto ${!showNewProjectErrors ? "hidden" : ""}`}
                >
                    {newProjectErrors.map((error) => (
                        <p key={error}>{error}</p>
                    ))}
                </div>
            </form>
        </div>
    );
};

export default NewProject;
