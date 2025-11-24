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
        <section id="how-it-works" className="py-20 px-4 sm:px-6 lg:px-8">
            <div className="max-w-7xl mx-auto">
                <div className="text-center space-y-4 mb-16">
                    <h2 className="text-4xl md:text-5xl font-bold text-foreground text-balance">
                        How it works
                    </h2>
                    <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
                        Getting started with Note Nest is simple and intuitive.
                    </p>
                </div>

                <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
                    {steps.map((step, index) => (
                        <div key={index} className="relative">
                            <StepsCard
                                key={index}
                                title={step.title}
                                description={step.description}
                                step={index + 1}
                            />
                            {index < steps.length - 1 && (
                                <div className="hidden lg:block absolute -right-4 top-12 w-8 h-0.5 bg-border"></div>
                            )}
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
};

export default HowItWorks;
