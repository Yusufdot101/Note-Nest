const Signup = () => {
    return (
        <div className="bg-primary flex flex-col mx-uto w-full shadow-[0px_0px_4px_1px_white] py-[32px] min-[620px]:text-2xl px-[12px]">
            <p className="text-accent text-[32px] font-semibold text-center">SIGN UP</p>
            <form action="" className="flex flex-col text-text gap-y-[8px]">
                <div className="flex flex-col">
                    <label htmlFor="username">Username</label>
                    <input type="text" id="username" className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="email">Email</label>
                    <input type="text" id="email" className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="password">password</label>
                    <input type="text" id="password" className="bg-white p-[8px] rounded-[8px] h-[50px] outline-none text-black" />
                </div>
                <p>Already have an account? <a href="#" className="text-accent">Login here</a></p>
                <button className="w-full py-[12px] rounded-[8px] cursor-pointer bg-accent mx-auto">Sign Up</button>
            </form>
        </div>
    )
}

export default Signup
