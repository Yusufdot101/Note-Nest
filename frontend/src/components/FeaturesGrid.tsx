import {
    Lock,
    Globe,
    MessageSquare,
    Heart,
    Share2,
    Folder,
} from "lucide-react";
import FeatureCard from "./FeatureCard";

const features = [
    {
        icon: Folder,
        title: "Organize Hierarchically",
        description:
            "Create projects, folders, and groups to organize your findings exactly how you think.",
        color: "text-accent",
    },
    {
        icon: Lock,
        title: "Privacy Controls",
        description:
            "Choose privacy for projects and notes. Keep personal notes to yourself or share publicly.",
        color: "text-accent",
    },
    {
        icon: Globe,
        title: "Public Discovery",
        description:
            "Share your public projects with the world and discover amazing content from others.",
        color: "text-accent",
    },
    {
        icon: Heart,
        title: "Like & Engage",
        description:
            "Show appreciation for great work by liking notes you find valuable.",
        color: "text-accent",
    },
    {
        icon: MessageSquare,
        title: "Rich Comments",
        description:
            "Discuss, provide feedback, and collaborate through thoughtful comments on public entries.",
        color: "text-accent",
    },
    {
        icon: Share2,
        title: "Easy Sharing",
        description: "Share specific notes with custom links.",
        color: "text-accent",
    },
];

const FeaturesGrid = () => {
    return (
        <section id="features" className="w-full flex flex-col gap-y-[24px]">
            <div className="text-center space-y-[12px]">
                <h2 className="text-4xl md:text-5xl font-bold text-foreground text-balance">
                    Everything you need
                </h2>
                <p className="text-lg">
                    Powerful features to capture, organize, and share your
                    knowledge with complete privacy control.
                </p>
            </div>

            <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-[12px]">
                {features.map((feature, index) => {
                    return (
                        <FeatureCard
                            key={index}
                            icon={feature.icon}
                            color={feature.color}
                            title={feature.title}
                            description={feature.description}
                        />
                    );
                })}
            </div>
        </section>
    );
};

export default FeaturesGrid;
