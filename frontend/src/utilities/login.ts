const BASE_APIURL = import.meta.env.VITE_API_URL || "http://localhost:8080"
export const handleLogin = async (email: string, password: string, handleErrors: (error: string) => void) => {
    try {
        const res = await fetch(`${BASE_APIURL}/users/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: email, password: password })
        })

        const data = await res.json()
        if (!res.ok) {
            const error = data.error
            if (error) {
                handleErrors(error)
                return
            }
            throw new Error(`HTTP error! status: ${res.status}`)
        }

        const token = data.token as string
        localStorage.setItem("sessionToken", token)
        // navigate to the home page when the the account is created
        window.location.replace("/")
    } catch (error) {
        alert("an error occurred, please try again")
        console.error(error)
    }
}
