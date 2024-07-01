import React from 'react';
import { AppBar, Box, Button, IconButton, InputBase, Toolbar } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';

const Header: React.FC = () => {
  return (
    <AppBar sx={{ backgroundColor: 'black', height: '70px', width: 'calc(100% - 272px)', ml: '272px' }}>
      <Toolbar sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%' }}>
        <Box sx={{ display: 'flex', alignItems: 'center', width: '70%', justifyContent: 'flex-end' }}>
          <Box
            sx={{
              display: 'flex',
              alignItems: 'center',
              backgroundColor: '#282a39',
              borderRadius: 1,
              width: '100%',
              padding: '0 10px',
              height: 40,
            }}
          >
            <IconButton sx={{ padding: 0, color: 'white' }}>
              <SearchIcon />
            </IconButton>
            <InputBase
              placeholder="Search for something..."
              sx={{ ml: 1, flex: 1, color: 'white', fontWeight: '200', fontSize: '14px' }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </Box>
        </Box>
        <Button
          sx={{
            height: 40,
            backgroundColor: '#08D2ED',
            color: 'white',
            '&:hover': {
              backgroundColor: '#06b1cc',
            },
            px: 2,
          }}
        >
          Đăng nhập
        </Button>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
