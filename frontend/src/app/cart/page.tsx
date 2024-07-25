import React from 'react';
import { Box, Button, Grid, Typography } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';

const Cart: React.FC = () => {
  return (
    <>
      <Typography variant="h6" sx={{ color: 'white', fontSize: '20px', fontWeight: 'bold', textAlign: 'center', marginBottom: 2 }}>
        Giỏ hàng
      </Typography>
      <Grid container spacing={2}>
        {/* Left Section */}
        <Grid item xs={8}>
          <Box sx={{ border: '1px solid white', borderRadius: 2, padding: 1.5 }}>
            <Grid container spacing={2} alignItems="center">
              <Grid item xs={3}>
                <Typography variant="body2" sx={{ color: 'white', pl: 2 }}>
                  Sản phẩm
                </Typography>
              </Grid>
              <Grid item xs={2}>
                <Typography variant="body2" sx={{ color: 'white' }}>
                  Sàn TMĐT
                </Typography>
              </Grid>
              <Grid item xs={2}>
                <Typography variant="body2" sx={{ color: 'white' }}>
                  Đơn giá
                </Typography>
              </Grid>
              <Grid item xs={2}>
                <Typography variant="body2" sx={{ color: 'white' }}>
                  Số lượng
                </Typography>
              </Grid>
              <Grid item xs={2}>
                <Typography variant="body2" sx={{ color: 'white' }}>
                  Thành tiền
                </Typography>
              </Grid>
              <Grid item xs={1}>
                <DeleteIcon sx={{ color: 'white' }} />
              </Grid>
            </Grid>
          </Box>
        </Grid>

        {/* Right Section */}
        <Grid item xs={4}>
          <Box sx={{ border: '1px solid white', padding: 2, borderRadius: 2, marginBottom: 2 }}>
            <Grid container spacing={2}>
              <Grid item xs={6}>
                <Typography sx={{ color: 'white', fontWeight: 300, fontSize: '16px' }}>
                  Giao tới
                </Typography>
              </Grid>
              <Grid item xs={6}>
                <Typography variant="body2" sx={{ color: '#08D2ED', textAlign: 'right' }}>
                  Thay đổi
                </Typography>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="body2" sx={{ color: 'white' }}>
                  Họ tên nhận hàng | Số điện thoại
                </Typography>
              </Grid>
            </Grid>
          </Box>
          <Box sx={{ border: '1px solid white', padding: 2 }}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography variant="body2" sx={{ color: '#E2E2E2' }}>
                  Tạm tính
                </Typography>
              </Grid>
              <Grid item xs={12}>
                <Box sx={{ borderTop: '1px solid white', paddingTop: 2 }}>
                  <Typography variant="body2" sx={{ color: '#E2E2E2' }}>
                    Tổng tiền
                  </Typography>
                  <Typography variant="body2" sx={{ color: '#979797', mt: 2 }}>
                    (Giá chưa áp dụng các mã giảm giá)
                  </Typography>
                </Box>
              </Grid>
            </Grid>
          </Box>
          <Button
            variant="contained"
            fullWidth
            sx={{
              backgroundColor: '#08D2ED',
              color: 'white',
              fontWeight: 'bold',
              marginTop: 2,
              '&:hover': { backgroundColor: '#08d1ed', borderColor: 'transparent' }
            }}
          >
            Mua hàng
          </Button>
        </Grid>
      </Grid>
    </>
  );
};

export default Cart;
