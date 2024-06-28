"use client";

import React, { useState } from 'react';
import { Box, Button } from '@mui/material';

const categories = [
  'Dành cho bạn',
  'Bách hóa',
  'Thời trang',
  'Phong cách sống',
  'Điện tử',
  'Làm đẹp',
  'Giáo dục',
  'Sách'
];

const CategorySlider: React.FC = () => {
  const [activeCategory, setActiveCategory] = useState<string>(categories[0]);

  return (
    <Box sx={{ display: 'flex', flexDirection: 'row', overflowX: 'auto', paddingY: 1, whiteSpace: 'nowrap' }}>
      {categories.map((category) => (
        <Button
          key={category}
          onClick={() => setActiveCategory(category)}
          sx={{
            backgroundColor: activeCategory === category ? '#dddddd' : '#6B6B6B',
            color: activeCategory === category ? '#020202' : 'white',
            margin: '0 5px',
            px: 2,
            whiteSpace: 'nowrap',
            '&:hover': {
              backgroundColor: activeCategory === category ? '#dddddd' : '#5a5a5a',
            },
          }}
        >
          {category}
        </Button>
      ))}
    </Box>
  );
};

export default CategorySlider;
