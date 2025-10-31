const emailRegexPattern = new RegExp("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

export const getUsernameErrorMessages = (username: string): string => {
    let message = ""
    if (username.trim().length < 2) {
        message += "Username must be at least 2 characters."
    }
    return message
}

export const getEmailErrorMessages = (email: string): string => {
    let message = ""
    if (!emailRegexPattern.test(email)) {
        message += "Email must a valid email address."
    }
    return message
}

export const getPasswordErrorMessages = (password: string): string => {
    let message = ""
    if (password.trim().length < 8) {
        message += "Password must be at least 8 characters."
    }
    return message
}
