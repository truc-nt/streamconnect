import { GridColDef } from "@mui/x-data-grid";
import DataGrid from "@/components/core/DataGrid";
import { Stack } from "@mui/material";
import {
  VisibilityOutlined,
  DeleteOutlined,
  CachedOutlined,
} from "@mui/icons-material";

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

export default function ExternalProductDataGrid() {
  return <DataGrid rows={[]} columns={columns} />;
}
