const emailRegexPattern = new RegExp(
    "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
);

export const getUsernameErrorMessages = (username: string): string => {
    let message = "";
    const minLength = 2;
    if (username.trim().length < minLength) {
        message += `Username must be at least ${minLength} character(s).`;
    }
    return message;
};

export const getEmailErrorMessages = (email: string): string => {
    let message = "";
    if (!emailRegexPattern.test(email)) {
        message += "Email must be a valid email address.";
    }
    return message;
};

export const getPasswordErrorMessages = (password: string): string => {
    let message = "";
    const minLength = 8;
    if (password.trim().length < minLength) {
        message += `Password must be at least ${minLength} character(s).`;
    }
    return message;
};

export const getProjectNameErrorMessages = (projectName: string): string => {
    let message = "";
    const minLength = 1;
    if (projectName.trim().length < minLength) {
        message += `Project name must be at least ${minLength} character(s).`;
    }
    return message;
};

export const getProjectDescriptionErrorMessages = (
    projectDescription: string,
): string => {
    let message = "";
    const minLength = 0;
    if (projectDescription.trim().length < minLength) {
        message += `Project description must be at least ${minLength} character(s).`;
    }
    return message;
};

export const getProjectVisibilityErrorMessages = (
    projectVisibility: string,
): string => {
    let message = "";
    const minLength = 1;
    const allowedVisibilityValues = ["public", "private"];
    if (projectVisibility.trim().length < minLength) {
        message += `Project visibility must be at least ${minLength} character(s).`;
    }
    if (!allowedVisibilityValues.includes(projectVisibility.toLowerCase())) {
        message += `Project visibility must be either ${allowedVisibilityValues.join('"or"')}.`;
    }
    return message;
};
