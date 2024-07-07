'use client';

import React, { useRef } from 'react';
import { Box, Typography, IconButton } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';

const products = [
  { id: 1, name: 'Product 1', price: 100000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
  { id: 2, name: 'Product 2', price: 200000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
  { id: 3, name: 'Product 3', price: 300000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
  { id: 4, name: 'Product 4', price: 400000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
  { id: 5, name: 'Product 5', price: 500000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
  { id: 6, name: 'Product 6', price: 600000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg' },
];

const ProductList: React.FC = () => {
  const formatPrice = (price: number) => {
    return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
  };

  const scrollContainerRef = useRef<HTMLDivElement>(null);

  const scrollLeft = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: -300, behavior: 'smooth' });
    }
  };

  const scrollRight = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: 300, behavior: 'smooth' });
    }
  };

  return (
    <Box
      sx={{
        width: '100%',
        mx: 'auto',
        my: 3,
        display: 'flex',
        alignItems: 'center',
      }}
    >
      <IconButton
        onClick={scrollLeft}
        sx={{
          backgroundColor: '#01E0EE',
          borderRadius: '50%',
          '&:hover': { backgroundColor: '#01E0EE' },
          marginRight: 2,
        }}
      >
        <FontAwesomeIcon icon={faArrowLeft} style={{ fontSize: '24px', color: 'black' }} />
      </IconButton>

      <Box
        ref={scrollContainerRef}
        sx={{
          overflowX: 'auto',
          display: 'flex',
          gap: 2,
          msOverflowStyle: 'none',
          scrollbarWidth: 'none',
          '::-webkit-scrollbar': { display: 'none' },
        }}
      >
        {products.map((product) => (
          <Box
            key={product.id}
            sx={{
              flex: '0 0 auto',
              width: '30%',
              backgroundColor: '#333',
              padding: 2,
              borderRadius: 2,
              display: 'flex',
              alignItems: 'center',
              border: '1px solid white', 
            }}
          >
            <img
              src={product.image}
              alt={product.name}
              style={{ width: '50%', height: 'auto', borderRadius: '8px', marginRight: '16px' }}
            />
            <Box sx={{ display: 'flex', flexDirection: 'column' }}>
              <Typography
                variant="body1"
                sx={{
                  color: 'white',
                  fontSize: '14px',
                  fontWeight: 'normal',
                  whiteSpace: 'nowrap',
                  overflow: 'hidden',
                  textOverflow: 'ellipsis',
                  maxWidth: '100%',
                }}
              >
                {product.name}
              </Typography>
              <Typography variant="body2" sx={{ color: 'white', fontSize: '18px' }}>
                đ {formatPrice(Number(product.price))}
              </Typography>
            </Box>
          </Box>
        ))}
      </Box>

      <IconButton
        onClick={scrollRight}
        sx={{
          backgroundColor: '#01E0EE',
          borderRadius: '50%',
          '&:hover': { backgroundColor: '#01E0EE' },
          marginLeft: 2,
        }}
      >
        <FontAwesomeIcon icon={faArrowRight} style={{ fontSize: '24px', color: 'black' }} />
      </IconButton>
    </Box>
  );
};

export default ProductList;
