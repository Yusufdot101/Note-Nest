import SubmitButton from "./SubmitButton";
import { useNavigate } from "react-router-dom";

const CTA = () => {
    const navigate = useNavigate();
    return (
        <section className="w-full flex flex-col bg-gradient-to-br from-primary/10 to-accent/20 border border-primary/20 rounded-2xl p-[8px] text-center space-y-[12px]">
            <h2 className="text-4xl md:text-3xl font-bold text-balance">
                Ready to build your nest?
            </h2>
            <p className="text-lg">
                Join thousands of creators organizing and sharing their
                knowledge with Note Nest.
            </p>
            <div className="flex gap-[8px] max-[700px]:flex-col">
                <SubmitButton
                    aria_label="login"
                    bgColor=""
                    text={"Get Started Now"}
                    handleSubmit={() => {
                        navigate("/login");
                    }}
                />

                <SubmitButton
                    aria_label="explore public content"
                    bgColor="grey"
                    text={"Explore Public Content"}
                    handleSubmit={() => {
                        navigate("/notes");
                    }}
                />
            </div>
        </section>
    );
};

export default CTA;
