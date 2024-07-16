'use client';

import React, { useRef, useState } from 'react';
import { Box, Typography, IconButton } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';
import DetailModal from './DetailModal'; // Import DetailModal

const products = [
  { id: 1, name: 'Product 1', price: 100000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 1', categories1: ['Màu đỏ', 'Màu xanh', 'Màu vàng', 'Màu trắng'], categories2: ['Size L', 'Size XL'] },
  { id: 2, name: 'Product 2', price: 200000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 2', categories1: ['Màu đỏ', 'Màu xanh'], categories2: ['Size L', 'Size XL'] },
  { id: 3, name: 'Product 3', price: 300000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 3', categories1: ['Màu đỏ', 'Màu xanh', 'Màu vàng'], categories2: ['Size L', 'Size XL'] },
  { id: 4, name: 'Product 4', price: 400000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 4', categories1: ['Màu đỏ', 'Màu xanh', 'Màu vàng'], categories2: ['Size L', 'Size XL'] },
  { id: 5, name: 'Product 5', price: 500000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 5', categories1: ['Màu đỏ', 'Màu xanh', 'Màu vàng'], categories2: ['Size L', 'Size XL'] },
  { id: 6, name: 'Product 6', price: 600000, image: 'https://assets.tmecosys.com/image/upload/t_web767x639/img/recipe/ras/Assets/5C4B5768-8901-433D-8A8E-7A2E912BB22E/Derivates/49b89a69-8241-460b-9d56-8b4438b5636d.jpg', description: 'Mô tả sản phẩm 6', categories1: ['Màu đỏ', 'Màu xanh', 'Màu vàng'], categories2: ['Size L', 'Size XL'] },
];

const ProductList: React.FC = () => {
  const [selectedProduct, setSelectedProduct] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

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

  const handleProductClick = (product) => {
    setSelectedProduct(product);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
  };

  return (
    <>
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
              onClick={() => handleProductClick(product)}
              sx={{
                flex: '0 0 auto',
                width: '30%',
                backgroundColor: '#333',
                padding: 2,
                borderRadius: 2,
                display: 'flex',
                alignItems: 'center',
                border: '1px solid white',
                cursor: 'pointer',
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

      {selectedProduct && (
        <DetailModal
          open={isModalOpen}
          onClose={handleCloseModal}
          product={selectedProduct}
        />
      )}
    </>
  );
};

export default ProductList;
