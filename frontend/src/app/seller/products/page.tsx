"use client";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import { Stack, Autocomplete, TextField, Box } from "@mui/material";

import {
  VisibilityOutlined,
  DeleteOutlined,
  CachedOutlined,
} from "@mui/icons-material";
import { useGetExternalShops } from "@/hook/external_shop";
import { useState } from "react";

const columns: GridColDef[] = [
  {
    field: "name",
    headerName: "Tên",
  },
  {
    field: "sku",
    headerName: "Sku",
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
    field: "quantity",
    headerName: "Số lượng",
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

const rows = [
  {
    id: 1,
    name: "Tên 1",
    sku: "Sku 1bcgivegfvievierbvivribervibeivbeviebvierivribvribvibbirvbirvivribibvribrvbivbibvj",
    status: "Đang hoạt đô",
  },
  {
    id: 2,
    name: "Tên 2",
    sku: "Sku 2",
  },
  {
    id: 3,
    name: "Tên 3",
    sku: "Sku 3",
  },
  {
    id: 4,
    name: "Tên 4",
    sku: "Sku 4",
  },
  {
    id: 5,
    name: "Tên 5",
    sku: "Sku 5",
  },
  /*{
    id: 6,
    name: "Tên 6",
    sku: "Sku 6",
  },
  {
    id: 7,
    name: "Tên 7",
    sku: "Sku 7",
  },
  {
    id: 8,
    name: "Tên 8",
    sku: "Sku 8",
  },
  {
    id: 9,
    name: "Tên 9",
    sku: "Sku 9",
  },
  {
    id: 10,
    name: "Tên 10",
    sku: "Sku 10",
  },*/
];

export default function Page() {
  return <DataGrid rows={[]} columns={[]} />;
}
