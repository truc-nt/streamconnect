import React from 'react';
import { Modal, Box, Typography, Avatar, Button } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faAngleRight } from '@fortawesome/free-solid-svg-icons';

interface Product {
  id: number;
  name: string;
  price: number;
  image: string;
  description: string;
  categories1: string[];
  categories2: string[];
}

interface DetailModalProps {
  open: boolean;
  onClose: () => void;
  product: Product;
}

const DetailModal: React.FC<DetailModalProps> = ({ open, onClose, product }) => {
  const formatPrice = (price: number) => {
    return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.');
  };

  return (
    <Modal open={open} onClose={onClose} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <Box
        sx={{
          backgroundColor: '#282a39',
          border: '1px solid white',
          borderRadius: 2,
          width: '60%',
          maxHeight: '90%',
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
          <Box sx={{ width: '50%', pr: 2 }}>
            <img
              src={product.image}
              alt={product.name}
              style={{ width: '100%', borderRadius: '8px' }}
            />
          </Box>
          {/* Product Details */}
          <Box sx={{ width: '50%', pl: 2, display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
            <Typography
              variant="h5"
              sx={{ color: 'white', mb: 1.5 }}
            >
              {product.name}
            </Typography>
            <Typography
              variant="h5"
              sx={{ color: '#08D2ED', mb: 2 }}
            >
              {formatPrice(product.price)} đ - {formatPrice(product.price * 2)} đ
            </Typography>
            <Box
              sx={{
                backgroundColor: '#54576b',
                fontFamily: 'Roboto',
                borderRadius: 2,
                p: 2,
                mb: 2,
                textAlign: 'left',
              }}
            >
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 1.5 }}>
                <Typography
                  variant="body1"
                  sx={{ color: 'white' }}
                >
                  Chọn loại hàng ({product.categories1.length} màu, {product.categories2.length} size)
                </Typography>
                <FontAwesomeIcon icon={faAngleRight} style={{ color: '#E2E2E2' }} />
              </Box>
              <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap', justifyContent: 'flex-start' }}>
                {product.categories1.slice(0, 3).map((category, index) => (
                  <Box
                    key={index}
                    sx={{
                      backgroundColor: '#E2E2E2',
                      color: '#54576B',
                      borderRadius: 1,
                      p: 1,
                    }}
                  >
                    {category}
                  </Box>
                ))}
                {product.categories1.length > 3 && (
                  <Box
                    sx={{
                      backgroundColor: '#E2E2E2',
                      color: '#54576B',
                      borderRadius: 1,
                      p: 1,
                    }}
                  >
                    +{product.categories1.length - 3}
                  </Box>
                )}
              </Box>
            </Box>
          </Box>
        </Box>

        {/* Bottom Section */}
        <Box sx={{ mb: 2 }}>
          <Box
            sx={{
              backgroundColor: '#54576B',
              p: 2,
              borderRadius: 2,
              display: 'flex',
              justifyContent: 'space-between',
              alignItems: 'center',
              mb: 2,
            }}
          >
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Avatar src="/assets/user-avatar.jpg" alt="User" sx={{ mr: 2 }} />
              <Typography sx={{ color: 'white' }}>Username</Typography>
            </Box>
            <Button
              variant="outlined"
              sx={{
                color: 'white',
                borderColor: 'white',
              }}
            >
              Truy cập
            </Button>
          </Box>

          <Box
            sx={{
              backgroundColor: '#54576B',
              p: 2,
              borderRadius: 2,
              color: 'white',
            }}
          >
            <Typography variant="h6" sx={{ mb: 1 }}>
              Giới thiệu về sản phẩm này
            </Typography>
            <Typography sx={{ fontWeight: '300', fontSize: '16px' }}>{product.description}</Typography>
          </Box>
        </Box>
      </Box>
    </Modal>
  );
};

export default DetailModal;
