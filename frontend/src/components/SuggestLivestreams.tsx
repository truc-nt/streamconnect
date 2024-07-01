'use client'

import React from 'react';
import { Box, Grid, Avatar, Typography, Badge, Tooltip, Button } from '@mui/material';

interface Livestream {
  imageUrl: string;
  avatarUrl: string;
  username: string;
  livestreamName: string;
  views: number;
}

// Mock data for livestreams
const mockLivestreams: Livestream[] = [
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User1',
    livestreamName: 'Livestream 1',
    views: 1000,
  },
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User2',
    livestreamName: 'Livestream 2',
    views: 800,
  },
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User3',
    livestreamName: 'Livestream 3',
    views: 600,
  },
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User3',
    livestreamName: 'Livestream 3',
    views: 600,
  },
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User3',
    livestreamName: 'Livestream 3',
    views: 600,
  },
  {
    imageUrl: 'https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png',
    avatarUrl: 'https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png',
    username: 'User3',
    livestreamName: 'Livestream 3',
    views: 600,
  },
];

const SuggestLivestreams: React.FC = () => {
  const handleLivestreamClick = (livestream: Livestream) => {
    console.log(`Clicked on ${livestream.livestreamName}`);
  };

  return (
    <Box>
      <Typography variant="h6" sx={{ fontSize: '20px', fontWeight: 'bold', color: 'white', marginY: 3 }}>
        Livestreams Được Đề Xuất
      </Typography>

      <Grid container spacing={2}>
        {mockLivestreams.map((livestream, index) => (
          <Grid key={index} item xs={12} sm={6} md={4} >
            <Button
              onClick={() => handleLivestreamClick(livestream)}
              sx={{ padding: 0, textTransform: 'none', cursor: 'pointer' }}
              fullWidth
            >
              <Box sx={{ position: 'relative', width: '100%', cursor: 'pointer', mb: 2 }}>
                {/* Livestream Thumbnail */}
                <img
                  src={livestream.imageUrl}
                  alt="Livestream Thumbnail"
                  style={{ width: '100%', height: 'auto', borderRadius: '8px' }}
                />

                {/* Viewer Count */}
                <Box sx={{ position: 'absolute', top: 8, left: 8, backgroundColor: 'rgba(0, 0, 0, 0.5)', padding: '4px 8px', borderRadius: '4px' }}>
                  <Typography variant="caption" sx={{ color: 'white' }}>
                    {livestream.views} views
                  </Typography>
                </Box>

                {/* Badge Live */}
                <Box sx={{ position: 'absolute', top: 8, right: 24 }}>
                  <Badge badgeContent="Live" color="error">
                    <Box />
                  </Badge>
                </Box>

                {/* Avatar and Livestream Info */}
                <Box sx={{ mt: 1, left: 0, width: '100%', textAlign: 'left', paddingLeft: '8px', paddingRight: '8px' }}>
                  <Avatar alt={livestream.username} src={livestream.avatarUrl} sx={{ width: 36, height: 36, marginRight: 1, display: 'inline-block' }} />
                  <Box style={{ display: 'inline-block', verticalAlign: 'top' }}>
                    <Typography variant="body2" sx={{ fontWeight: 'bold', color: 'white', display: 'block' }}>
                      {livestream.username}
                    </Typography>
                    <Typography variant="body2" sx={{ color: 'white', display: 'block' }}>
                      {livestream.livestreamName}
                    </Typography>
                  </Box>
                </Box>
              </Box>
            </Button>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default SuggestLivestreams;
