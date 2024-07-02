'use client';

import React from 'react';
import { Box, Avatar, Typography, Badge, IconButton } from '@mui/material';
import ReactPlayer from 'react-player';
import { FeaturedVideo, Fullscreen, Visibility, VolumeOff } from '@mui/icons-material';
import { keyframes } from '@emotion/react';

const gradientAnimation = keyframes`
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
`;

const LivestreamPreview: React.FC = () => {
  return (
    <Box
      sx={{
        mt: 3,
        mb: 4,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
      }}
    >
      <Box
        sx={{
          position: 'relative',
          width: '100%',
          height: '400px',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          '&::before': {
            content: '""',
            position: 'absolute',
            top: '-8px',
            left: '-8px',
            right: '-8px',
            bottom: '-8px',
            zIndex: 1,
            borderRadius: '8px',
            // background: 'linear-gradient(45deg, red, orange, yellow, green, indigo, violet)',
            background: 'linear-gradient(45deg, #08d2ed, #282a39)',
            // background: 'linear-gradient(45deg, #8b0000, black)',
            backgroundSize: '400% 400%',
            animation: `${gradientAnimation} 15s ease infinite`,
          },
          '&::after': {
            content: '""',
            position: 'absolute',
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            zIndex: 1,
            borderRadius: '8px',
            padding: '8px', 
            boxSizing: 'border-box',
            background: 'inherit',
            backgroundClip: 'padding-box',
          },
        }}
      >
        <Box
          sx={{
            position: 'relative',
            width: '100%',
            height: '100%',
            borderRadius: '8px',
            overflow: 'hidden',
            zIndex: 2,
          }}
        >
          <ReactPlayer
            url="https://www.youtube.com/watch?v=VBKNoLcj8jA"
            width="100%"
            height="100%"
            controls
            playing
            muted
          />
        </Box>

        <Box sx={{ position: 'absolute', top: 16, left: 16, display: 'flex', alignItems: 'center', zIndex: 3 }}>
          <Avatar alt="Username" src="/assets/your-avatar-image.jpg" />
          <Box sx={{ ml: 2 }}>
            <Typography variant="subtitle1" sx={{ fontWeight: 'bold', color: 'white' }}>
              Username
            </Typography>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <Typography variant="body2" sx={{ color: '#D8D8D8' }}>
                Livestream Name
              </Typography>
              <Visibility sx={{ fontSize: '16px', color: '#D8D8D8' }} />
              <Typography variant="body2" sx={{ color: '#D8D8D8' }}>
                1234
              </Typography>
            </Box>
          </Box>
        </Box>

        <Badge
          badgeContent="Live"
          color="error"
          sx={{ position: 'absolute', top: 16, right: 16, zIndex: 3, marginRight: 2, marginTop: 2 }}
        >
          <Box />
        </Badge>

        <Box sx={{ position: 'absolute', bottom: 16, right: 16, display: 'flex', gap: 1, zIndex: 3 }}>
          <IconButton sx={{ color: 'white' }}>
            <FeaturedVideo />
          </IconButton>
          <IconButton sx={{ color: 'white' }}>
            <Fullscreen />
          </IconButton>
          <IconButton sx={{ color: 'white' }}>
            <VolumeOff />
          </IconButton>
        </Box>
      </Box>
    </Box>
  );
};

export default LivestreamPreview;
