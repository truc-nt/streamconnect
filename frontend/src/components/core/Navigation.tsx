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
  Collapse,
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
  ExpandLess,
  ExpandMore,
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
    label: "Sản phẩm",
    icon: Category,
    children: [
      {
        path: "/seller/products",
        label: "Hệ thống",
      },
      {
        path: "/seller/external_products",
        label: "Liên kết",
      },
    ],
  },
];

interface INavigationItem {
  path?: string;
  label: string;
  icon?: React.ElementType<any>;
  children?: INavigationItem[];
}

const NavigationItem = ({ path, label, icon, children }: INavigationItem) => {
  const [open, setOpen] = useState(false);

  const pathname = usePathname();

  const handleClick = () => {
    setOpen(!open);
  };

  return (
    <>
      <NavigationLink path={path}>
        <ListItemButton
          style={{
            backgroundColor: pathname === path ? "#535561" : "transparent",
            borderRadius: 8,
          }}
          selected={pathname === path}
          onClick={children ? handleClick : undefined}
        >
          {icon && (
            <ListItemIcon style={{ minWidth: "40px" }}>
              <SvgIcon
                component={icon}
                style={{
                  borderRadius: 50,
                  backgroundColor: "white",
                  color: "black",
                  padding: "2px",
                  fontSize: "18px",
                }}
              />
            </ListItemIcon>
          )}
          <ListItemText primary={<Typography>{label}</Typography>} />
          {children && (open ? <ExpandLess /> : <ExpandMore />)}
        </ListItemButton>
      </NavigationLink>
      {children && (
        <Collapse
          in={open}
          timeout="auto"
          unmountOnExit
          style={{ paddingLeft: "10px" }}
        >
          <List component="nav">
            {children.map((child, index) => (
              <NavigationItem key={index} {...child} />
            ))}
          </List>
        </Collapse>
      )}
    </>
  );
};

const NavigationLink = ({
  path,
  children,
}: {
  path: string | undefined;
  children: React.ReactNode;
}) => {
  return path ? (
    <Link
      href={path}
      passHref
      style={{ textDecoration: "none", color: "inherit" }}
    >
      {children}
    </Link>
  ) : (
    <>{children}</>
  );
};

const Navigation = () => {
  const activeSegment = useSelectedLayoutSegment();

  let items = mainNavigator;

  if (activeSegment === "user") {
    items = userNavigator;
  } else if (activeSegment === "seller") {
    items = sellerNavigator;
  }

  return (
    <List component="nav">
      {items.map((item, index) => (
        <NavigationItem key={index} {...item} />
      ))}
    </List>
  );
};

export default Navigation;
