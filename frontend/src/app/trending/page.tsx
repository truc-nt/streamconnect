import CategorySlider from "@/components/CategorySlider";
import LivestreamPreview from "@/components/LivestreamPreview";
import LivestreamsList from "@/components/LivestreamsList";

export default function Home() {
    return (
      <>
        <CategorySlider />
        <LivestreamPreview/>
        <LivestreamsList />
      </>
    );
  }
  