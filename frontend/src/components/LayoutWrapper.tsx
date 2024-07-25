'use client';

import React from 'react';
import { Box } from '@mui/material';
import { usePathname } from 'next/navigation';
import NavigationBoard from '@/components/NavigationBoard';
import Header from '@/components/Header';

const LayoutWrapper: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const pathname = usePathname();
  const isCheckoutPage = pathname === '/cart';

  return (
    <Box sx={{ display: 'flex', height: '100vh', overflow: 'hidden' }}>
      {!isCheckoutPage && (
        <Box sx={{ width: 272, flexShrink: 0 }}>
          <NavigationBoard />
        </Box>
      )}
      <Box sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column' }}>
        <Header showName={isCheckoutPage} />
        <Box
          sx={{
            flexGrow: 1,
            padding: 2,
            backgroundColor: '#1c1c1c',
            overflowY: 'auto',
            '::-webkit-scrollbar': { display: 'none' },
            msOverflowStyle: 'none',
            scrollbarWidth: 'none',
          }}
        >
          <Box sx={{ marginX: 1.5 }}>
            {children}
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

export default LayoutWrapper;
