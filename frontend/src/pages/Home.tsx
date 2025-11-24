import HeroSection from "../components/HeroSection";
import FeaturesGrid from "../components/FeaturesGrid";
import HowItWorks from "../components/HowItWorks";
import CTA from "../components/CTASection";
import { useAuthStore } from "../store/useAuthStore";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

const Home = () => {
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
    const navigate = useNavigate();
    useEffect(() => {
        if (isLoggedIn) navigate("/notes");
    }, [isLoggedIn, navigate]);

    return (
        <div className="flex flex-col items-center relative text-text bg-primary p-[12px] h-fit rounded-[8px] border-[1px] border-white">
            <HeroSection />
            <FeaturesGrid />
            <HowItWorks />
            <CTA />
        </div>
    );
};

export default Home;
