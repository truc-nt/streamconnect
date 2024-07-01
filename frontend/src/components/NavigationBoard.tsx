'use client'

import React, { useState } from 'react';
import { Box, Button, Divider, Typography, Avatar } from '@mui/material';
import { Circle } from '@mui/icons-material';
import WhatshotIcon from '@mui/icons-material/Whatshot';
import ExploreIcon from '@mui/icons-material/Explore';
import PeopleAltIcon from '@mui/icons-material/PeopleAlt';


const formatViews = (views: number) => {
  return views.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
};

const NavigationBoard: React.FC = () => {
  const [activeButton, setActiveButton] = useState<string | null>(null);

  const handleButtonClick = (buttonName: string) => {
    setActiveButton(buttonName);
  };

  const mostViewedItems = [
    {
      username: 'User1',
      livestreamName: 'Livestream 1',
      views: 10000,
      avatarSrc: '/assets/avatar1.jpg',
    },
    {
      username: 'User2',
      livestreamName: 'Livestream 2',
      views: 8000,
      avatarSrc: '/assets/avatar2.jpg',
    },
    {
      username: 'User3',
      livestreamName: 'Livestream 3',
      views: 6000,
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
        onClick={() => handleButtonClick('explore')}
        sx={{
          color: 'white',
          textTransform: 'none',
          marginBottom: 2,
          justifyContent: 'flex-start',
          textAlign: 'left',
          backgroundColor: activeButton === 'explore' ? '#535561' : 'inherit',
          fontWeight: activeButton === 'explore' ? 'bold' : 'normal',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
          gap: 2
        }}
      >
        <ExploreIcon sx={{fontSize: '18px'}}/>Khám phá
      </Button>

      <Button
        fullWidth
        onClick={() => handleButtonClick('trending')}
        sx={{
          color: 'white',
          marginBottom: 2,
          textTransform: 'none',
          justifyContent: 'flex-start',
          textAlign: 'left',
          backgroundColor: activeButton === 'trending' ? '#535561' : 'inherit',
          fontWeight: activeButton === 'trending' ? 'bold' : 'normal',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
          gap: 2
        }}
      >
        <WhatshotIcon sx={{fontSize: '15px', borderRadius: 50, backgroundColor:'white', color: 'black'}}/>
        Xu hướng
      </Button>

      <Button
        fullWidth
        onClick={() => handleButtonClick('following')}
        sx={{
          color: 'white',
          marginBottom: 2,
          textTransform: 'none',
          justifyContent: 'flex-start',
          textAlign: 'left',
          backgroundColor: activeButton === 'following' ? '#535561' : 'inherit',
          fontWeight: activeButton === 'following' ? 'bold' : 'normal',
          '&:hover': {
            backgroundColor: '#4a4a4a',
          },
          gap: 2
        }}
      >
        <PeopleAltIcon sx={{fontSize: '15px', borderRadius: 50, backgroundColor:'white', color: 'black'}}/> Đang follow
      </Button>

      <Divider sx={{ width: '100%', marginY: 2, borderColor: '#4a4a4a' }} />

      <Typography variant="h6" sx={{ fontWeight: 'bold', pt: 2 }}>
        Xem nhiều nhất
      </Typography>

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
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <Typography sx={{ fontSize: '14px', fontWeight: '100', lineHeight: '20px' }}>
              {formatViews(item.views)}
            </Typography>
            <Circle sx={{ color: '#EF233C', fontSize: 'small' }} />
          </Box>
        </Button>
      ))}
    </Box>
  );
};

export default NavigationBoard;
