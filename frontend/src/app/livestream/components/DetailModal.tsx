import React, { useState, useEffect } from 'react';
import { Modal, Box, Typography, Avatar, Button } from '@mui/material';

interface Product {
  id: number;
  name: string;
  image: string;
  description: string;
  categories: string[];
  categoryValues: {
    [category: string]: string[];
  };
  prices: {
    platform: string;
    category1: string;
    category2: string;
    price: number;
  }[];
}

interface DetailModalProps {
  open: boolean;
  onClose: () => void;
  product: Product;
}

const DetailModal: React.FC<DetailModalProps> = ({ open, onClose, product }) => {
  const formatPrice = (price?: number) => {
    if (typeof price === 'number') {
      return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
    }
    return '0';
  };

  const [selectedCategory1, setSelectedCategory1] = useState<string | null>(null);
  const [selectedCategory2, setSelectedCategory2] = useState<string | null>(null);

  useEffect(() => {
    if (!open) {
      setSelectedCategory1(null);
      setSelectedCategory2(null);
    }
  }, [open]);

  const handleCategory1Click = (category: string) => {
    setSelectedCategory1(category);
    setSelectedCategory2(null); // Reset category2 when category1 changes
  };

  const handleCategory2Click = (category: string) => {
    setSelectedCategory2(category);
  };

  const getPriceRange = () => {
    const filteredPrices = product.prices.filter(p => {
      return (!selectedCategory1 || p.category1 === selectedCategory1) &&
        (!selectedCategory2 || p.category2 === selectedCategory2 || !p.category2);
    });

    if (filteredPrices.length === 0) return '0';

    const prices = filteredPrices.map(p => p.price);
    const minPrice = Math.min(...prices);
    const maxPrice = Math.max(...prices);

    return minPrice === maxPrice
      ? `${formatPrice(minPrice)} đ`
      : `${formatPrice(minPrice)} đ - ${formatPrice(maxPrice)} đ`;
  };

  const getExactPrice = (platform: string) => {
    if (selectedCategory1 === null && selectedCategory2 === null) {
      const priceForPlatform = product.prices.find(p => p.platform === platform);
      if (priceForPlatform) {
        return `${formatPrice(priceForPlatform.price)} đ`;
      }
    }

    const filteredPrices = product.prices.filter(p =>
      p.platform === platform &&
      p.category1 === (selectedCategory1 || '') &&
      (!p.category2 || p.category2 === (selectedCategory2 || ''))
    );

    if (filteredPrices.length > 0) {
      const price = filteredPrices[0].price;
      return `${formatPrice(price)} đ`;
    }

    return getPriceRange();
  };

  const platforms = Array.from(new Set(product.prices.map(price => price.platform)));

  const hasCategories = product.categories.length > 0;

  return (
    <Modal open={open} onClose={onClose} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <Box
        sx={{
          backgroundColor: '#282a39',
          border: '1px solid white',
          borderRadius: 2,
          width: '60%',
          height: '550px',
          overflowY: 'auto',
          p: 2,
          display: 'flex',
          flexDirection: 'column',
          '::-webkit-scrollbar': { display: 'none' },
          msOverflowStyle: 'none',
          scrollbarWidth: 'none',
        }}
      >
        {/* Top Section */}
        <Box sx={{ display: 'flex', mb: 2 }}>
          {/* Product Image */}
          <Box sx={{ width: '70%', mr: 2, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <img src={product.image} alt={product.name} style={{ width: '100%', height: 'auto', borderRadius: '8px' }} />
          </Box>
          {/* Product Details */}
          <Box sx={{ display: 'flex', flexDirection: 'column', width: '50%' }}>
            <Typography variant="h6" sx={{ color: 'white' }}>{product.name}</Typography>
            <Typography variant="h6" sx={{ color: '#08D2ED' }}>{getPriceRange()}</Typography>
            {hasCategories && (
              <Box
                sx={{
                  overflowY: 'auto',
                  '::-webkit-scrollbar': { display: 'none' },
                  msOverflowStyle: 'none',
                  scrollbarWidth: 'none',
                }}
              >
                {product.categories.map((category, index) => (
                  <Box key={index} sx={{ display: 'flex', flexDirection: 'column', mt: 1, backgroundColor: '#54576b', p: 1, borderRadius: 1, width: '100%' }}>
                    <Typography variant="body2" sx={{ color: 'white', mb: 1 }}>{category}</Typography>
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                      {product.categoryValues[category].map((value, idx) => (
                        <Button
                          key={idx}
                          variant={selectedCategory1 === value || selectedCategory2 === value ? 'contained' : 'outlined'}
                          sx={{
                            backgroundColor: selectedCategory1 === value || selectedCategory2 === value ? '#08d1ed' : '#e2e2e2',
                            color: selectedCategory1 === value || selectedCategory2 === value ? 'white' : 'black',
                            '&.Mui-contained:hover': { backgroundColor: '#08d1ed' },
                            '&.Mui-outlined:hover': { backgroundColor: '#e2e2e2', borderColor: 'transparent' },
                            '&:hover': {
                              backgroundColor: selectedCategory1 === value || selectedCategory2 === value ? '#08d1ed' : '#e2e2e2',
                              borderColor: 'transparent',
                            },
                            '&.Mui-contained': { backgroundColor: '#08d1ed', color: 'white' },
                            '&.Mui-outlined': { borderColor: 'transparent' },
                          }}
                          onClick={() => category === product.categories[0] ? handleCategory1Click(value) : handleCategory2Click(value)}
                        >
                          {value}
                        </Button>
                      ))}
                    </Box>
                  </Box>
                ))}
              </Box>
            )}
            {/* Platforms Prices Section */}
            <Box sx={{ display: 'flex', flexDirection: 'column', backgroundColor: '#54576b', p: 1, borderRadius: 1, my: 1 }}>
              <Typography variant="body2" sx={{ color: 'white', mb: 1 }}>Giá các sàn</Typography>
              <Box sx={{ display: 'flex', overflowX: 'auto', gap: 2, '::-webkit-scrollbar': { display: 'none' }, msOverflowStyle: 'none', scrollbarWidth: 'none' }}>
                {platforms.map((platform, index) => (
                  <Box key={index} sx={{ display: 'flex', backgroundColor: '#282a39', p: 1, borderRadius: 1, border: '1px solid white', minWidth: '350px' }}>
                    {/* Left Section */}
                    <Box sx={{ display: 'flex', flexDirection: 'column', flex: 1 }}>
                      <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                        <Avatar sx={{ mr: 1 }}>S</Avatar>
                        <Typography variant="body2" sx={{ color: 'white' }}>{platform}</Typography>
                      </Box>
                      <Typography variant="body2" sx={{ color: '#08D2ED' }}>
                        {getExactPrice(platform)}
                      </Typography>
                    </Box>

                    {/* Right Section */}
                    <Box sx={{ display: 'flex', flexDirection: 'column', justifyContent: 'space-between', ml: 2 }}>
                      <Button variant="contained" sx={{ backgroundColor: '#01E0EE', color: 'black', fontSize: '12px', mb: 1, '&:hover': { backgroundColor: '#08d1ed', borderColor: 'transparent' } }}>Thêm vào giỏ hàng</Button>
                      <Button variant="contained" sx={{ backgroundColor: '#01E0EE', color: 'black', fontSize: '12px', '&:hover': { backgroundColor: '#08d1ed', borderColor: 'transparent' } }}>Mua ngay</Button>
                    </Box>
                  </Box>
                ))}
              </Box>
            </Box>
          </Box>
        </Box>

        {/* Bottom Section */}
        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2, backgroundColor: '#54576b', p: 2, borderRadius: 1 }}>
          <Avatar sx={{ mr: 2 }}>U</Avatar>
          <Typography variant="body2" sx={{ color: 'white' }}>{product.name}</Typography>
          <Button
            variant="contained"
            sx={{
              ml: 'auto',
              backgroundColor: '#01E0EE',
              color: 'black',
              fontSize: '12px',
              '&:hover': { backgroundColor: '#08d1ed', borderColor: 'transparent' },
            }}
          >
            Truy cập
          </Button>
        </Box>

        {/* Product Description */}
        <Box sx={{ backgroundColor: '#54576b', p: 2, borderRadius: 1 }}>
          <Typography variant="body2" sx={{ color: 'white' }}>{product.description}</Typography>
        </Box>
      </Box>
    </Modal>
  );
};

export default DetailModal;

