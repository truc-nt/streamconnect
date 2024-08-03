// Livestream Component
import React from 'react';
import { Box, Typography } from '@mui/material';
import LivestreamFrame from './components/LivestreamFrame';
import ChatSection from './components/ChatSection';
import ProductList from './components/ProductList';

const Livestream: React.FC = () => {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
      <Box sx={{ display: 'flex', flex: 1, overflow: 'hidden', gap: 2.5 }}>
        <Box sx={{ flex: 2.5, display: 'flex', alignItems: 'stretch' }}>
          <LivestreamFrame />
        </Box>
        <Box sx={{ flex: 1, display: 'flex' }}>
          <ChatSection />
        </Box>
      </Box>
      <Typography variant="h6" sx={{ color: 'white', fontSize: '20px', fontWeight: 'bold', marginTop: 2 }}>
        Livestream name
      </Typography>
      <Box sx={{ flexShrink: 0, overflowX: 'hidden' }}>
        <ProductList />
      </Box>
    </Box>
  );
};

export default Livestream;
