"use client";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import { Stack, Autocomplete, TextField, Box, Button } from "@mui/material";

import {
  VisibilityOutlined,
  DeleteOutlined,
  CachedOutlined,
} from "@mui/icons-material";
import { useGetProducts } from "@/hook/shop";
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
    field: "price",
    headerName: "Giá",
  },
  {
    field: "stock",
    headerName: "Số lượng",
  },
  {
    field: "updated_at",
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
  const { data, error } = useGetProducts(process.env.NEXT_PUBLIC_SHOP_ID);

  return (
    <Stack gap={2}>
      <Stack direction="row" spacing={2}>
        <Button variant="contained">Thêm sản phẩm</Button>
      </Stack>
      <DataGrid
        rows={data}
        columns={columns}
        getRowId={(row) => row.id_product}
      />
    </Stack>
  );
}
