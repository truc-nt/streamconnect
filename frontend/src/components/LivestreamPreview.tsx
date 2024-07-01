'use client';

import React from 'react';
import { Box, Avatar, Typography, Badge, IconButton } from '@mui/material';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ReactPlayer from 'react-player';
import { Fullscreen, Visibility, VolumeOff } from '@mui/icons-material';

const LivestreamPreview: React.FC = () => {
  return (
    <Box
      sx={{
        mt: 2,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
      }}
    >
      <Box sx={{ position: 'relative', width: '100%', height: '540px' }}>
        <ReactPlayer
          url="https://www.youtube.com/watch?v=VBKNoLcj8jA"
          width="100%"
          height="100%"
          controls
          playing
          muted
        />

        <Box sx={{ position: 'absolute', top: 16, left: 16, display: 'flex', alignItems: 'center' }}>
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
          sx={{ position: 'absolute', top: 16, right: 16, zIndex: 1, marginRight: 2, marginTop: 2 }}
        >
          <Box />
        </Badge>

        <Box sx={{ position: 'absolute', bottom: 16, right: 16, display: 'flex', gap: 1 }}>
          <IconButton sx={{ color: 'white' }}>
            <FavoriteIcon />
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
