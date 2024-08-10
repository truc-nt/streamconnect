"use client";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import { Stack, Autocomplete, TextField, Box } from "@mui/material";

import {
  VisibilityOutlined,
  DeleteOutlined,
  CachedOutlined,
} from "@mui/icons-material";
import {
  useGetExternalShops,
  useGetExternalProducts,
} from "@/hook/external_shop";
import { getExternalProducts } from "@/api/external_shop";
import { useState, useEffect } from "react";
import { IExternalShop } from "@/api/shop";

const columns: GridColDef[] = [
  {
    field: "external_product_name",
    headerName: "Tên",
    width: 300,
  },
  {
    field: "product_name",
    headerName: "Tên hệ thống",
  },
  {
    field: "total_stock",
    headerName: "Số lượng",
  },
  {
    field: "",
    headerName: "Giá",
    renderCell: ({ row }) => (
      <span>
        {row.min_price} - {row.max_price}
      </span>
    ),
  },
  {
    field: "updated_at",
    headerName: "Ngày cập nhật",
    renderCell: ({ value }) => new Date(value).toLocaleString(),
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
  const {
    data: externalShops,
    error,
    isLoading,
  } = useGetExternalShops(process.env.NEXT_PUBLIC_SHOP_ID);

  const [externalShop, setExternalShop] = useState<IExternalShop | null>(null);
  const [externalProducts, setExternalProducts] = useState([]);
  //const [selectedExternalShopName, setSelectedExternalShopName] = useState("");
  //console.log(externalShop, selectedExternalShopName);

  useEffect(() => {
    const fetchData = async () => {
      //setExternalShop(externalShops?.[0]!);
      try {
        const res = await getExternalProducts(
          externalShops?.[0]!.id_external_shop!,
        );
        setExternalProducts(res?.data);
      } catch (error) {
        console.error(error);
      }
      //setExternalProducts(res);
    };
    if (!externalShop) return;
    fetchData();
  }, [externalShop]);

  if (isLoading) return <></>;
  return (
    <Stack gap={2}>
      <Autocomplete
        value={externalShop}
        onChange={(event: any, newValue: any) => {
          setExternalShop(newValue);
        }}
        inputValue={externalShop?.name || ""}
        onInputChange={(event, newInputValue) => {
          //setSelectedExternalShopName(newInputValue);
        }}
        disablePortal
        autoHighlight
        getOptionLabel={(option) => option.name}
        options={externalShops || []}
        renderOption={(props, option) => {
          const { key, ...optionProps } = props;
          return (
            <Box key={key} component="li" {...optionProps}>
              {option.name}
            </Box>
          );
        }}
        sx={{
          width: "300px",
        }}
        renderInput={(params) => (
          <TextField
            {...params}
            label="Chọn cửa hàng"
            inputProps={{
              ...params.inputProps,
            }}
            variant="outlined"
          />
        )}
      />
      {externalProducts && (
        <DataGrid
          rows={externalProducts}
          columns={columns}
          getRowId={(row) => row.external_product_external_id}
        />
      )}
    </Stack>
  );
}
