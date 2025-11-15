import { useEffect, useState } from "react";
import Input from "../components/Input";
import SubmitButton from "../components/SubmitButton";
import {
    getProjectDescriptionErrorMessages,
    getProjectNameErrorMessages,
    getProjectVisibilityErrorMessages,
} from "../utilities/inputValidation";
import {
    deleteProject,
    fetchProject,
    updateProject,
} from "../utilities/project";
import { useNavigate, useParams } from "react-router-dom";
import ColorPicker from "../components/ColorPicker";

const EditProject = () => {
    const [projectName, setProjectName] = useState("");
    const [projectDescription, setProjectDescription] = useState("");
    const [projectVisibility, setProjectVisibility] = useState("");
    const [projectColor, setProjectColor] = useState("#00FFFF");

    const { id } = useParams();

    useEffect(() => {
        const setupProject = async () => {
            if (id === "") return;
            const project = await fetchProject(+id!);
            if (!project) return;
            setProjectName(project.Name);
            setProjectDescription(project.Description);
            setProjectVisibility(project.Visibility);
            setProjectColor(project.Color);
        };
        setupProject();
    }, [id]);

    const [projectNameError, setProjectNameError] = useState("");
    const [projectDescriptionError, setProjectDescriptionError] = useState("");
    const [projectVisibilityError, setProjectVisibilityError] = useState("");
    const [showError, setShowError] = useState(false);

    const navigate = useNavigate();

    const handleSave = async () => {
        setShowError(true);
        if (
            projectNameError ||
            projectDescriptionError ||
            projectVisibilityError ||
            !id
        )
            return;
        const success = await updateProject(
            +id,
            projectName,
            projectDescription,
            projectVisibility,
            projectColor,
        );
        if (!success) return;
        // Navigate to projects list after successful update
        navigate("/projects");
    };

    const handleCancel = () => {
        navigate(`/projects/${id}`);
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
            className="bg-primary flex flex-col w-full py-[32px] min-[620px]:text-2xl px-[12px] rounded-[8px]"
        >
            <div className="flex items-center justify-center gap-[8px] h-[35px]">
                <p className="text-accent text-[32px] max-[619px]:text-[24px] font-semibold text-center">
                    EDIT PROJECT
                </p>
                <ColorPicker color={projectColor} setColor={setProjectColor} />
            </div>
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                }}
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
                        id="projectDescription"
                        className="w-[100%] h-[100px] min-h-[50px] max-[619px]:min-h-[40px] bg-white rounded-[8px] p-[8px] outline-none text-black"
                        value={projectDescription}
                        onChange={(e) => setProjectDescription(e.target.value)}
                    />
                    <p
                        aria-label={"project description error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectDescriptionError"
                    >
                        {projectDescriptionError}
                    </p>
                </div>
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
                    <p
                        aria-label={"project visibility error"}
                        className={`text-red-500 ${!showError ? "hidden" : ""}`}
                        id="projectVisibilityError"
                    >
                        {projectVisibilityError}
                    </p>
                </div>

                <div className="flex gap-[8px]">
                    <SubmitButton
                        aria_label={"Save Changes"}
                        handleSubmit={() => {
                            handleSave();
                        }}
                        text={"Save"}
                    />
                    <SubmitButton
                        aria_label={"Cancel Changes"}
                        handleSubmit={() => {
                            if (!id) return;
                            handleCancel();
                        }}
                        text={"Cancel"}
                        bgColor="grey"
                    />
                </div>
                <SubmitButton
                    aria_label={"Delete Project"}
                    handleSubmit={async () => {
                        if (!id) return;
                        const success = await deleteProject(id);
                        if (!success) return;
                        navigate("/projects");
                    }}
                    text={"Delete Project"}
                    bgColor="red"
                />
            </form>
        </div>
    );
};

export default EditProject;
