import NavigationBoard from '@/components/NavigationBoard';
import Header from '@/components/Header';
import { Box } from '@mui/material';
import CategorySlider from '@/components/CategorySlider';
import LivestreamPreview from '@/components/LivestreamPreview';
import SuggestLivestreams from '@/components/SuggestLivestreams';

export default function Home() {
  return (
    <Box sx={{ display: 'flex', height: '100vh', overflow: 'hidden' }}>
      <Box sx={{ width: 272, flexShrink: 0 }}>
        <NavigationBoard />
      </Box>
      <Box sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column' }}>
        <Header />
        <Box
          sx={{
            flexGrow: 1,
            padding: 2,
            backgroundColor: '#1c1c1c',
            overflowY: 'auto',
            '::-webkit-scrollbar': { display: 'none' },
            '-ms-overflow-style': 'none',
            'scrollbar-width': 'none',
          }}
        >
          <Box sx={{ marginX: 1.5 }}>
            <CategorySlider />
            <LivestreamPreview />
            <SuggestLivestreams />
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
