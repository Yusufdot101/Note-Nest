import { CheckCircle2 } from "lucide-react";

type Props = {
    title: string;
    description: string;
    step: number;
};

const StepsCard = ({ title, description, step }: Props) => {
    return (
        <div
            tabIndex={0}
            className="group border-[1px] border-white rounded-[8px] p-[12px] group flex flex-col gap-y-[12px]"
        >
            <div className="flex items-center gap-4">
                <div className="flex-shrink-0 w-10 h-10 bg-primary/20 rounded-full flex items-center justify-center">
                    <CheckCircle2 className="w-8 h-8 text-accent group-hover:scale-110 group-focus:scale-110 transition-transform" />
                </div>
                <span className="text-[24px] font-bold text-accent">
                    Step {step}
                </span>
            </div>
            <h3 className="text-lg font-bold text-foreground mb-2">{title}</h3>
            <p className="text-sm text-muted-foreground">{description}</p>
        </div>
    );
};

export default StepsCard;
