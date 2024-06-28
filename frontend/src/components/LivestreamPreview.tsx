'use client';

import React from 'react';
import { Box, Avatar, Typography, Badge, IconButton } from '@mui/material';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ShareIcon from '@mui/icons-material/Share';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import ReactPlayer from 'react-player';

const LivestreamPreview: React.FC = () => {
  return (
    <Box
      sx={{
        mt: 2,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center', 
        width: '100%'   
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
            <Typography variant="body2" sx={{ color: 'white' }}>
              Livestream Name
            </Typography>
            <Typography variant="caption" sx={{ color: 'white' }}>
              1234 views
            </Typography>
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
            <ShareIcon />
          </IconButton>
          <IconButton sx={{ color: 'white' }}>
            <MoreVertIcon />
          </IconButton>
        </Box>
      </Box>
    </Box>
  );
};

export default LivestreamPreview;
