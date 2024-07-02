import CategorySlider from '@/components/CategorySlider';
import LivestreamPreview from '@/components/LivestreamPreview';
import LivestreamsList from '@/components/LivestreamsList';
import { Typography } from '@mui/material';


export default function Home() {
  return (
    <>
      <CategorySlider />
      <LivestreamPreview />
      <Typography variant="h6" sx={{ fontSize: '20px', fontWeight: 'bold', color: 'white' }}>
        Livestreams Được Đề Xuất
      </Typography>
      <LivestreamsList />
    </>
  );
}
