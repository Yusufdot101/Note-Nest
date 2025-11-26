import { Sparkles } from "lucide-react";

const HeroSection = () => {
    return (
        <section className="w-full grid md:grid-cols-2 gap-[12px] items-center">
            <div className="space-y-2">
                <div className="inline-flex items-center gap-2 bg-accent/10 text-accent px-4 py-2 rounded-full">
                    <Sparkles className="w-4 h-4" />
                    <span className="text-sm font-medium">
                        Organize. Share. Discover.
                    </span>
                </div>

                <h1 className="text-5xl md:text-6xl font-bold text-balance leading-tight">
                    Your personal knowledge{" "}
                    <span className="text-accent">nest</span>
                </h1>

                <p className="text-lg text-balance">
                    Capture your findings, organize them beautifully, and share
                    with a community of curious minds. Keep what matters private
                    and let your best work shine.
                </p>
            </div>

            <div className="relative">
                <div className="bg-primary border border-border rounded-[8px] p-[24px] space-y-4 shadow-white shadow-sm">
                    <div className="pb-[12px] border-b border-border">
                        <div className="space-y-2">
                            <div className="flex gap-2">
                                <div className="w-8 h-8 bg-accent/50 rounded"></div>
                                <div className="flex-1 space-y-1">
                                    <div className="h-2 bg-accent/30 rounded w-24"></div>
                                    <div className="h-2 bg-accent/20 rounded w-20"></div>
                                    <div className="h-2 bg-accent/10 rounded w-16"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="space-y-[8px]">
                        <div className="h-3 bg-background rounded w-[90%]"></div>
                        <div className="h-3 bg-background rounded w-[80%]"></div>
                        <div className="h-3 bg-background rounded w-[75%]"></div>
                    </div>
                </div>
            </div>
        </section>
    );
};

export default HeroSection;
