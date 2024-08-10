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
  Stack,
  Button,
} from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import TelegramIcon from "@mui/icons-material/Telegram";
import NotificationsIcon from "@mui/icons-material/Notifications";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";
import { useRouter, usePathname } from "next/navigation";
import Link from "next/link";

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

  return (
    <AppBar
      position="static"
      sx={{ backgroundColor: "transparent", boxShadow: "none" }}
    >
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
        <Stack
          direction="row"
          spacing={1}
          justifyContent="center"
          alignItems="center"
        >
          {true ? (
            <>
              <IconButton>
                <TelegramIcon />
              </IconButton>
              <IconButton>
                <NotificationsIcon />
              </IconButton>
              <IconButton onClick={() => router.push("/cart")}>
                <ShoppingCartIcon />
              </IconButton>
              <IconButton
                onClick={(event: React.MouseEvent<HTMLElement>) => {
                  setAnchorEl(event.currentTarget);
                }}
                sx={{ padding: 0 }}
              >
                <Avatar alt="User Avatar" src="/path-to-avatar.jpg" />
              </IconButton>
              <Menu
                anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={() => setAnchorEl(null)}
                anchorOrigin={{
                  vertical: "bottom",
                  horizontal: "right",
                }}
                transformOrigin={{
                  vertical: "top",
                  horizontal: "right",
                }}
                /*sx={{
                  width: "300px",
                  //mt: 1.5,
                  py: 10,
                  "& .MuiMenuItem-root": {
                    fontSize: "15px",
                    marginX: "20px",
                    borderRadius: "4px",
                    "&:hover": {
                      backgroundColor: "#535561",
                    },
                    "&.Mui-selected": {
                      backgroundColor: "#535561",
                    },
                  },
                }}*/
                PaperProps={{
                  sx: {
                    width: "175px",
                    mt: 1.5,
                    "& .MuiMenuItem-root": {
                      fontSize: "15px",
                      marginX: "8px",
                      borderRadius: "4px",
                    },
                  },
                }}
              >
                <MenuItem onClick={handleMenuClose}>Hồ sơ</MenuItem>
                <MenuItem onClick={handleMenuClose}>Livestream</MenuItem>
                <MenuItem onClick={() => router.push("/seller")}>
                  Cửa hàng
                </MenuItem>
                <Divider />
                <MenuItem onClick={handleMenuClose}>Đăng xuất</MenuItem>
              </Menu>
            </>
          ) : (
            <>
              <Button variant="contained" color="secondary">
                Đăng kí
              </Button>
              <Button variant="contained">Đăng nhập</Button>
            </>
          )}
        </Stack>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
