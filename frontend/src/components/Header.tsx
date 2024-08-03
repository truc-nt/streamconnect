"use client";

import React, { useState } from "react";
import {
  AppBar,
  Box,
  IconButton,
  InputBase,
  Toolbar,
  Avatar,
  Menu,
  MenuItem,
  Divider,
} from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import TelegramIcon from "@mui/icons-material/Telegram";
import NotificationsIcon from "@mui/icons-material/Notifications";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import { useRouter, usePathname } from "next/navigation";

interface HeaderProps {
  showName?: boolean;
}

const Header: React.FC<HeaderProps> = ({ showName }) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const router = useRouter();
  const pathname = usePathname();

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleCartClick = () => {
    router.push("/cart");
  };

  const handleNameClick = () => {
    router.push("/");
  };

  return (
    <AppBar position="static" sx={{ backgroundColor: "transparent" }}>
      <Toolbar
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          width: "100%",
        }}
      >
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            width: "70%",
            justifyContent: "flex-end",
          }}
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              backgroundColor: "background.default",
              borderRadius: 1,
              width: "100%",
              padding: "0 10px",
              height: 40,
            }}
          >
            <IconButton sx={{ padding: 0, color: "white" }}>
              <SearchIcon />
            </IconButton>
            <InputBase
              placeholder="Search for something..."
              sx={{
                ml: 1,
                flex: 1,
                color: "white",
                fontWeight: "200",
                fontSize: "14px",
              }}
              inputProps={{ "aria-label": "search" }}
            />
          </Box>
        </Box>
        {/* <Button
          sx={{
            height: 40,
            backgroundColor: '#08D2ED',
            color: 'white',
            '&:hover': {
              backgroundColor: '#06b1cc',
            },
            px: 2
          }}
        >
          Đăng nhập
        </Button> */}
        <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
          <IconButton sx={{ color: "white" }}>
            <TelegramIcon />
          </IconButton>
          <IconButton sx={{ color: "white" }}>
            <NotificationsIcon />
          </IconButton>
          <IconButton sx={{ color: "white" }}>
            <ShoppingCartIcon />
          </IconButton>
          <IconButton onClick={handleMenuOpen} sx={{ padding: 0 }}>
            <Avatar alt="User Avatar" src="/path-to-avatar.jpg" />
          </IconButton>
          <Menu
            anchorEl={anchorEl}
            open={Boolean(anchorEl)}
            onClose={handleMenuClose}
            PaperProps={{
              sx: {
                //backgroundColor: "#282A39",
                color: "#F4F4F5",
                width: "175px",
                mt: 1.5,
                "& .MuiMenuItem-root": {
                  fontSize: "15px",
                  marginX: "8px",
                  borderRadius: "4px",
                  "&:hover": {
                    backgroundColor: "#535561",
                  },
                  "&.Mui-selected": {
                    backgroundColor: "#535561",
                  },
                },
                "& .MuiDivider-root": {
                  //backgroundColor: "#F4F4F5",
                },
              },
            }}
          >
            <MenuItem onClick={handleMenuClose}>Hồ sơ</MenuItem>
            <MenuItem onClick={handleMenuClose}>Livestream</MenuItem>
            <MenuItem onClick={handleMenuClose}>Cửa hàng</MenuItem>
            <Divider />
            <MenuItem onClick={handleMenuClose}>Đăng xuất</MenuItem>
          </Menu>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
