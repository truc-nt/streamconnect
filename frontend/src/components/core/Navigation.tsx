"use client";

import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  Box,
  Button,
  Divider,
  Typography,
  Avatar,
  MenuItem,
  MenuList,
  SvgIcon,
} from "@mui/material";
import { Circle } from "@mui/icons-material";
import {
  Whatshot,
  Explore,
  PeopleAlt,
  Person,
  Inventory,
  Cable,
  BarChart,
  Store,
  Category,
  SmartDisplay,
} from "@mui/icons-material";
import ListSubheader from "@mui/material/ListSubheader";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import DraftsIcon from "@mui/icons-material/Drafts";
import SendIcon from "@mui/icons-material/Send";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";

import Link from "next/link";
import { useSelectedLayoutSegment, usePathname } from "next/navigation";

interface INavigationItem {
  path: string;
  label: string;
  icon: React.ElementType<any>;
}

const mainNavigator: INavigationItem[] = [
  {
    path: "/",
    label: "Khám phá",
    icon: Explore,
  },
  {
    path: "/trending",
    label: "Xu hướng",
    icon: Whatshot,
  },
  {
    path: "/following",
    label: "Đang theo dõi",
    icon: PeopleAlt,
  },
];

const userNavigator: INavigationItem[] = [
  {
    path: "/user/profile",
    label: "Thông tin tài khoản",
    icon: Person,
  },
  {
    path: "/user/orders",
    label: "Quản lý đơn hàng",
    icon: Inventory,
  },
  {
    path: "/user/external-acccounts",
    label: "Tài khoản liên kết",
    icon: Cable,
  },
];

const sellerNavigator: INavigationItem[] = [
  {
    path: "/seller/dashboard",
    label: "Kênh bán hàng",
    icon: BarChart,
  },
  {
    path: "/seller/livestreams",
    label: "Livestream",
    icon: SmartDisplay,
  },
  {
    path: "/seller/shops",
    label: "Cửa hàng",
    icon: Store,
  },
  {
    path: "/seller/products",
    label: "Sản phẩm",
    icon: Category,
  },
];

const Navigation = () => {
  const activeSegment = useSelectedLayoutSegment();
  const pathname = usePathname();

  let items = mainNavigator;

  if (activeSegment === "user") {
    items = userNavigator;
  } else if (activeSegment === "seller") {
    items = sellerNavigator;
  }

  return (
    <MenuList style={{ display: "flex", flexDirection: "column", gap: "15px" }}>
      {items.map((item) => (
        <Link
          key={item.path}
          href={item.path}
          style={{ textDecoration: "none", color: "inherit" }}
        >
          <MenuItem
            style={{
              backgroundColor:
                pathname === item.path ? "#535561" : "transparent",
              borderRadius: 8,
            }}
            selected={pathname === item.path}
          >
            <ListItemIcon
              style={{
                minWidth: "20px",
              }}
            >
              <SvgIcon
                component={item.icon}
                style={{
                  borderRadius: 50,
                  backgroundColor: "white",
                  color: "black",
                  fontSize: "13px",
                  padding: "1px",
                }}
              />
            </ListItemIcon>
            <ListItemText
              primary={
                <Typography
                  style={{
                    fontWeight: pathname === item.path ? "bold" : "normal",
                    fontSize: 14,
                  }}
                >
                  {item.label}
                </Typography>
              }
            />
          </MenuItem>
        </Link>
      ))}
    </MenuList>
  );
};

export default Navigation;
