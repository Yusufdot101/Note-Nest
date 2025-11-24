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
        <section
            id="features"
            className="py-20 px-4 sm:px-6 lg:px-8 bg-secondary/30"
        >
            <div className="max-w-7xl mx-auto">
                <div className="text-center space-y-4 mb-16">
                    <h2 className="text-4xl md:text-5xl font-bold text-foreground text-balance">
                        Everything you need
                    </h2>
                    <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
                        Powerful features to capture, organize, and share your
                        knowledge with complete privacy control.
                    </p>
                </div>

                <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
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
            </div>
        </section>
    );
};

export default FeaturesGrid;
