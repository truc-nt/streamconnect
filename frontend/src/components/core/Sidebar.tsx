"use client";

import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  Box,
  Button,
  Divider,
  Typography,
  Avatar,
  Stack,
  Container,
} from "@mui/material";
import { Circle } from "@mui/icons-material";
import WhatshotIcon from "@mui/icons-material/Whatshot";
import ExploreIcon from "@mui/icons-material/Explore";
import PeopleAltIcon from "@mui/icons-material/PeopleAlt";
import ListSubheader from "@mui/material/ListSubheader";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import DraftsIcon from "@mui/icons-material/Drafts";
import SendIcon from "@mui/icons-material/Send";
import Navigation from "@/components/core/Navigation";

const Sidebar = () => {
  return (
    <Stack
      spacing={2}
      sx={{
        py: 2,
        px: 3,
        height: "100%",
        backgroundColor: "background.default",
        color: "#fff",
      }}
    >
      <Typography variant="h6" sx={{ fontWeight: "bold" }}>
        NAME
      </Typography>
      <Navigation />
    </Stack>
  );
};

export default Sidebar;
