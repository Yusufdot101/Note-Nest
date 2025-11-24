import SubmitButton from "./SubmitButton";
import { useNavigate } from "react-router-dom";

const CTA = () => {
    const navigate = useNavigate();
    return (
        <section className="py-20 px-2 max-[619px]:px-3">
            <div className="max-w-4xl mx-auto">
                <div className="bg-gradient-to-br from-primary/10 to-accent/20 border border-primary/20 rounded-2xl p-12 md:p-16 text-center space-y-8">
                    <h2 className="text-4xl md:text-5xl font-bold text-foreground text-balance">
                        Ready to build your nest?
                    </h2>
                    <p className="text-lg max-w-2xl mx-auto">
                        Join thousands of creators organizing and sharing their
                        knowledge with Note Nest.
                    </p>
                    <div className="flex gap-[8px] max-[619px]:flex-col">
                        <SubmitButton
                            aria_label="login"
                            bgColor=""
                            text={"Get Started Now"}
                            handleSubmit={() => {
                                navigate("/login");
                            }}
                        />

                        <SubmitButton
                            aria_label="login"
                            bgColor="grey"
                            text={"Explore Public Content"}
                            handleSubmit={() => {
                                navigate("/notes");
                            }}
                        />
                    </div>
                </div>
            </div>
        </section>
    );
};

export default CTA;
