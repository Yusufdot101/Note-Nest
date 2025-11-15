import { useEffect, useState } from "react";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import { useNavigate } from "react-router-dom";
import {
    getProjectDescriptionErrorMessages,
    getProjectNameErrorMessages,
    getProjectVisibilityErrorMessages,
} from "../utilities/inputValidation";
import { newProject } from "../utilities/project";
import ColorPicker from "../components/ColorPicker";

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
        <div
            style={{ border: `1px solid ${projectColor}` }}
            className="bg-primary flex flex-col w-full py-[32px] min-[620px]:text-2xl px-[12px]"
        >
            <div className="flex h-[35px] items-center justify-center gap-[8px]">
                <p className="text-accent text-[32px] max-[619px]:text-[24px] font-semibold text-center">
                    NEW PROJECT
                </p>
                <ColorPicker color={projectColor} setColor={setProjectColor} />
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
                    <p
                        aria-label={"project name error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectNameError"
                    >
                        {projectNameError}
                    </p>
                </div>
                <div className="flex flex-col">
                    <label htmlFor={"projectDescription"}>
                        Project Description
                    </label>
                    <textarea
                        name="projectDescription"
                        value={projectDescription}
                        onChange={(e) => setProjectDescription(e.target.value)}
                        id="projectDescription"
                        className="w-[100%] h-[100px] min-h-[50px] max-[619px]:min-h-[40px] bg-white rounded-[8px] min-h-[50px] p-[8px] outline-none text-black"
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
                    <div className="flex items-center gap-[10px]">
                        <div className="flex items-center gap-[8px]">
                            <label htmlFor={"private"}>Private</label>
                            <input
                                type="radio"
                                name="projectVisibility"
                                id="private"
                                value={"private"}
                                className="w-[30px] h-[30px] max-[619px]:w-[20px] accent-accent"
                                checked={projectVisibility === "private"}
                                onChange={(e) =>
                                    setProjectVisibility(e.target.value)
                                }
                            />
                        </div>
                        <div className="flex items-center gap-[8px]">
                            <label htmlFor={"public"}>Public</label>
                            <input
                                type="radio"
                                name="projectVisibility"
                                id="public"
                                value={"public"}
                                className="w-[30px] h-[30px] max-[619px]:w-[20px] accent-accent"
                                checked={projectVisibility === "public"}
                                onChange={(e) =>
                                    setProjectVisibility(e.target.value)
                                }
                            />
                        </div>
                    </div>
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
