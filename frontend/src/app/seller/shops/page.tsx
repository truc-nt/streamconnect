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
import { setOpen as setOpenAlert } from "@/store/alert";
import { useAppDispatch } from "@/store/store";

export default function Page() {
  const [open, setOpen] = useState(false);
  const { data, error } = useGetExternalShops(process.env.NEXT_PUBLIC_SHOP_ID);
  const dispatch = useAppDispatch();
  const columns: GridColDef[] = [
    {
      field: "name",
      headerName: "Tên",
      renderCell: (data) => {
        return (
          <Stack
            direction="row"
            spacing={1}
            justifyContent="center"
            alignItems="center"
          >
            <Avatar
              alt="Shopify Logo"
              src={"/assets/imgs/logo-shopify.jpg"}
              sx={{ padding: 1, backgroundColor: "white" }}
            />
            <p>{data?.value}</p>
          </Stack>
        );
      },
    },
    {
      field: "status",
      headerName: "Trạng thái",
      renderCell: (data) => {
        switch (data.value) {
          case "active":
            return (
              <Chip label="Đang hoạt động" color="success" variant="outlined" />
            );
          case "inactive":
            return (
              <Chip label="Ngưng hoạt động" color="error" variant="outlined" />
            );
        }
      },
    },
    {
      field: "created_at",
      headerName: "Ngày kết nối",
      renderCell: (data) => {
        return new Date(data.value).toLocaleString();
      },
    },
    {
      field: "updated_at",
      headerName: "Ngày cập nhật",
      renderCell: (data) => {
        return new Date(data.value).toLocaleString();
      },
    },
    {
      field: "actions",
      headerName: "",
      renderCell: (data) => (
        <Stack direction="row" spacing={0.5}>
          <VisibilityOutlined />
          <CachedOutlined
            onClick={async () => {
              try {
                const res = await syncExternalVariants(
                  data.row.id_external_shop,
                );
                dispatch(
                  setOpenAlert({
                    message: "Đồng bộ thành công",
                    type: "success",
                  }),
                );
              } catch (error) {
                dispatch(
                  setOpenAlert({
                    message: "Đồng bộ không thành công",
                    type: "error",
                  }),
                );
                console.log(error);
              }
            }}
          />
          <DeleteOutlined />
        </Stack>
      ),
    },
  ];

  return (
    <Stack gap={2}>
      <Stack direction="row" spacing={2}>
        <Button variant="contained" color="secondary">
          Đồng bộ
        </Button>
        <Button variant="contained" onClick={() => setOpen(true)}>
          Liên kết cửa hàng
        </Button>
      </Stack>
      <DataGrid
        rows={data}
        columns={columns}
        getRowId={(row) => row.id_external_shop}
      />
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
    </Stack>
  );
}
