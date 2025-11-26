import StepsCard from "./StepsCard";

const steps = [
    {
        title: "Create Your Space",
        description: "Set up projects to organize your thoughts and findings.",
    },
    {
        title: "Add Notes",
        description:
            "Add rich notes to your projects with formatting, links, and media.",
    },
    {
        title: "Set Privacy",
        description:
            "Choose what stays private and what you want to share with the world.",
    },
    {
        title: "Engage & Connect",
        description:
            "Like, comment, and share with other creators in the Note Nest community.",
    },
];

const HowItWorks = () => {
    return (
        <section
            id="how-it-works"
            className="w-full flex flex-col gap-y-[24px]"
        >
            <div className="text-center space-y-[12px]">
                <h2 className="text-4xl md:text-5xl font-bold text-balance">
                    How it works
                </h2>
                <p className="text-lg">
                    Getting started with Note Nest is simple and intuitive.
                </p>
            </div>

            <div className="grid md:grid-cols-2 lg:grid-cols-2 gap-[12px]">
                {steps.map((step, index) => (
                    <StepsCard
                        key={index}
                        title={step.title}
                        description={step.description}
                        step={index + 1}
                    />
                ))}
            </div>
        </section>
    );
};

export default HowItWorks;
