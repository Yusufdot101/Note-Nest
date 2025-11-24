import { Sparkles } from "lucide-react";

const HeroSection = () => {
    return (
        <section className="pb-16 px-4 sm:px-6 lg:px-8 mt-16">
            <div className="max-w-7xl mx-auto">
                <div className="grid md:grid-cols-2 gap-12 items-center">
                    <div className="space-y-2">
                        <div className="space-y-2">
                            <div className="inline-flex items-center gap-2 bg-accent/10 text-accent px-4 py-2 rounded-full">
                                <Sparkles className="w-4 h-4" />
                                <span className="text-sm font-medium">
                                    Organize. Share. Discover.
                                </span>
                            </div>

                            <h1 className="text-5xl md:text-6xl font-bold text-foreground text-balance leading-tight">
                                Your personal knowledge{" "}
                                <span className="text-primary">nest</span>
                            </h1>

                            <p className="text-lg text-balance">
                                Capture your findings, organize them
                                beautifully, and share with a community of
                                curious minds. Keep what matters private and let
                                your best work shine.
                            </p>
                        </div>
                    </div>

                    <div className="relative">
                        <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-accent/20 rounded-2xl blur-3xl -z-10"></div>
                        <div className="bg-card border border-border rounded-2xl p-8 space-y-4 shadow-xl">
                            <div className="space-y-3">
                                <div className="h-3 bg-background rounded w-3/4"></div>
                                <div className="h-3 bg-background rounded w-3/4"></div>
                                <div className="h-3 bg-background rounded w-3/4"></div>
                            </div>
                            <div className="pt-4 border-t border-border">
                                <div className="space-y-2">
                                    <div className="flex gap-2">
                                        <div className="w-8 h-8 bg-accent/30 rounded"></div>
                                        <div className="flex-1 space-y-1">
                                            <div className="h-2 bg-accent/20 rounded w-20"></div>
                                            <div className="h-2 bg-accent/10 rounded w-16"></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
};

export default HeroSection;
