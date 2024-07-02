import CategorySlider from '@/components/CategorySlider';
import LivestreamPreview from '@/components/LivestreamPreview';
import SuggestLivestreams from '@/components/SuggestLivestreams';

export default function Home() {
  return (
    <>
      <CategorySlider />
      <LivestreamPreview />
      <SuggestLivestreams />
    </>
  );
}
