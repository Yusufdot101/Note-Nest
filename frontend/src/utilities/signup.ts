const BASE_APIURL = import.meta.env.VITE_API_URL || "http://localhost:8080"
export const handleSignup = async (username: string, email: string, password: string, handleErrors: (errors: Record<string, string>) => void) => {
    try {
        const res = await fetch(`${BASE_APIURL}/users/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name: username, email: email, password: password })
        })

        if (!res.ok) {
            const data = await res.json()
            const errors = await data.error
            if (errors) {
                handleErrors(errors)
                return
            }
            throw new Error(`HTTP error! status: ${res.status}`)
        }
        // navigate to the home page when the the account is created
        window.location.replace("/")
    } catch (error) {
        alert("an error occured, please try again")
        console.warn(error)
    }
}
