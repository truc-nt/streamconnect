"use client";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
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

const columns: GridColDef[] = [
  {
    field: "name",
    headerName: "Tên",
  },
  {
    field: "status",
    headerName: "Trạng thái",
  },
  {
    field: "updatedAt",
    headerName: "Ngày cập nhật",
  },
  {
    field: "actions",
    headerName: "",
    renderCell: () => (
      <Stack direction="row" spacing={0.5}>
        <VisibilityOutlined />
        <CachedOutlined />
        <DeleteOutlined />
      </Stack>
    ),
  },
];

export default function Page() {
  const [open, setOpen] = useState(false);
  return (
    <>
      <Stack spacing={1}>
        <Stack direction="row" spacing={2}>
          <Button variant="contained" color="secondary">
            Đồng bộ
          </Button>
          <Button variant="contained" onClick={() => setOpen(true)}>
            Liên kết cửa hàng
          </Button>
        </Stack>
        <DataGrid rows={[]} columns={columns} />
      </Stack>
      <Dialog
        open={open}
        onClose={() => setOpen(false)}
        fullWidth
        maxWidth="xs"
      >
        <DialogTitle>Liên kết shop</DialogTitle>
        <List>
          <ListItem>
            <ListItemButton>
              <ListItemAvatar>
                <Avatar
                  alt="Shopify Logo"
                  src="/assets/imgs/logo-shopify.jpg"
                  sx={{ padding: 1, backgroundColor: "white" }}
                />
              </ListItemAvatar>
              <ListItemText primary="Shopify" />
            </ListItemButton>
          </ListItem>
        </List>
      </Dialog>
    </>
  );
}
