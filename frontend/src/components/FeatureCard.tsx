import type { LucideProps } from "lucide-react";

type Props = {
    icon: React.ForwardRefExoticComponent<
        Omit<LucideProps, "ref"> & React.RefAttributes<SVGSVGElement>
    >;
    title: string;
    description: string;
    color: string;
};

const FeatureCard = ({ icon, title, description, color }: Props) => {
    const Icon = icon;
    return (
        <div
            tabIndex={0}
            className="border-[1px] border-white rounded-[8px] p-[12px] group flex flex-col gap-y-[12px]"
        >
            <Icon
                className={`w-8 h-8 ${color} mb-4 group-hover:scale-110 group-focus:scale-110 transition-transform`}
            />

            <div>
                <h3 className="text-lg font-bold text-foreground">{title}</h3>
            </div>

            <div>
                <p className="text-sm text-muted-foreground">{description}</p>
            </div>
        </div>
    );
};

export default FeatureCard;
