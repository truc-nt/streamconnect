"use client";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import Chip from "@mui/material/Chip";
import { Dialog, DialogTitle } from "@mui/material";
import {
  List,
  ListItemButton,
  ListItemAvatar,
  ListItem,
  ListItemText,
} from "@mui/material";
import Avatar from "@mui/material/Avatar";
import {
  VisibilityOutlined,
  DeleteOutlined,
  CachedOutlined,
} from "@mui/icons-material";
import { useState } from "react";
import { useGetExternalShops } from "@/hook/external_shop";
import { syncExternalVariants } from "@/api/external_shop";
import CardHeader from "@mui/material/CardHeader";
import Link from "next/link";

export default function Page() {
  return (
    <Stack gap={2}>
      <Stack direction="row" spacing={2}>
        <Link href="/seller/livestreams/create">
          <Button variant="contained">Thêm livestream</Button>
        </Link>
      </Stack>
    </Stack>
  );
}
