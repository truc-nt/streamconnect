import React from 'react';
import { Box, Button, Divider, Typography, Avatar } from '@mui/material';

const NavigationBoard: React.FC = () => {
  // Mock data for the most viewed items
  const mostViewedItems = [
    {
      username: 'User1',
      livestreamName: 'Livestream 1',
      views: 1000,
      avatarSrc: '/assets/avatar1.jpg',
    },
    {
      username: 'User2',
      livestreamName: 'Livestream 2',
      views: 800,
      avatarSrc: '/assets/avatar2.jpg',
    },
    {
      username: 'User3',
      livestreamName: 'Livestream 3',
      views: 600,
      avatarSrc: '/assets/avatar3.jpg',
    },
  ];

  return (
    <Box
      sx={{
        height: '100vh',
        width: '272px',
        backgroundColor: '#282A39',
        color: 'white',
        py: 2,
        px: 3,
      }}
    >
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          marginBottom: 4,
          textTransform: 'none',
        }}
      >
        <Typography variant="h6" sx={{ fontWeight: 'bold' }}>
          NAME
        </Typography>
      </Box>

      <Button
        fullWidth
        sx={{
          color: 'white',
          textTransform: 'none',
          marginBottom: 2,
          justifyContent: 'flex-start',
          textAlign: 'left',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
        }}
      >
        Khám phá
      </Button>

      <Button
        fullWidth
        sx={{
          color: 'white',
          marginBottom: 2,
          textTransform: 'none',
          justifyContent: 'flex-start',
          textAlign: 'left',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
        }}
      >
        Xu hướng
      </Button>

      <Button
        fullWidth
        sx={{
          color: 'white',
          marginBottom: 2,
          textTransform: 'none',
          justifyContent: 'flex-start',
          textAlign: 'left',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
        }}
      >
        Đang follow
      </Button>

      <Divider sx={{ width: '100%', marginY: 2, borderColor: '#4a4a4a' }} />

      <Typography variant="h6" sx={{ fontWeight: 'bold', pt: 2 }}>
        Xem nhiều nhất
      </Typography>

      {/* Render most viewed items */}
      {mostViewedItems.map((item, index) => (
        <Button
          key={index}
          fullWidth
          sx={{
            display: 'flex',
            justifyContent: 'flex-start',
            textTransform: 'none',
            color: 'white',
            marginTop: 3,
            '&:hover': {
              backgroundColor: '#4a4a4a',
            },
          }}
        >
          <Avatar src={item.avatarSrc} sx={{ width: 36, height: 36, marginRight: 1 }} />
          <Box sx={{ display: 'flex', flexDirection: 'column', justifyContent: 'center', flexGrow: 1, marginRight: 1 }}>
            <Typography sx={{ fontWeight: '400', fontSize: '14px', lineHeight: '20px', textAlign: 'left' }}>
              {item.username}
            </Typography>
            <Typography sx={{ fontWeight: '100', fontSize: '14px', color: '#9A9A9A', lineHeight: '20px', textAlign: 'left' }}>
              {item.livestreamName}
            </Typography>
          </Box>
          <Typography sx={{ fontSize: '14px', fontWeight: '100', lineHeight: '20px' }}>{item.views} views</Typography>
        </Button>
      ))}
    </Box>
  );
};

export default NavigationBoard;
