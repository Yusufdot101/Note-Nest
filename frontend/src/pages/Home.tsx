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
        <div className="flex flex-col gap-y-[64px] h-fit w-full min-w-[300px] border-[1px] border-solid border-[#ffffff] rounded-[8px] p-[24px]">
            <HeroSection />
            <FeaturesGrid />
            <HowItWorks />
            <CTA />
        </div>
    );
};

export default Home;
