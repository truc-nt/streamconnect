"use client";
import {
  Grid,
  Paper,
  Box,
  Typography,
  Stack,
  Divider,
  Button,
  TextField,
  FormControl,
  IconButton,
  Icon,
  Chip,
} from "@mui/material";
import DataGrid from "@/components/core/DataGrid";
import { GridColDef } from "@mui/x-data-grid";
import { Add } from "@mui/icons-material";
import { useState, useEffect } from "react";
import ProductInfo from "@/components/product/ProductInfo";

interface IChoiceItem {
  id_external_product: number;
  id_ecommerce: number;
  ecommerce: string;
  stock: number;
  price: number;
  option: Record<string, string>;
  quantity: number;
}

interface ChosenItem {
  id_external_product: number;
  id_ecommerce: number;
  ecommerce: string;
  stock: number;
  price: number;
  quantity: number;
  option: Record<string, string>;
}

const option_titles = {
  Title: ["A", "B", "c"],
  Color: ["red", "black"],
};

const externalProducts: IChoiceItem[] = [
  {
    id_external_product: 1,
    id_ecommerce: 1,
    ecommerce: "Shopify",
    stock: 50,
    price: 50,
    option: {
      Title: "A",
      Color: "red",
    },
    quantity: 0,
  },
  {
    id_external_product: 2,
    id_ecommerce: 1,
    ecommerce: "Shopify",
    stock: 30,
    price: 50,
    option: {
      Title: "B",
      Color: "red",
    },
    quantity: 0,
  },
];

const chosenColumns: GridColDef[] = [
  {
    field: "option",
    headerName: "Phân loại",
    renderCell: ({ row }) => (
      <Stack direction="row" spacing={1}>
        {Object.entries(row.option).map(([key, value]) => (
          <Chip key={key} label={`${key}: ${value}`} size="small" />
        ))}
      </Stack>
    ),
  },
  {
    field: "ecommerce",
    headerName: "Sàn thương mại",
  },
  {
    field: "price",
    headerName: "Giá",
  },
  {
    field: "quantity",
    headerName: "Số lượng",
  },
];

const ProductDetail = () => {
  const [chosenOption, setChosenOption] = useState<Record<string, string>>({});
  const [chosenData, setChosenData] = useState<ChosenItem[]>([]);
  const [choiceData, setChoiceData] = useState<IChoiceItem[]>([]);

  console.log(chosenOption);

  useEffect(() => {
    if (Object.keys(chosenOption).length < Object.keys(option_titles).length) {
      return;
    }
    const filteredChoiceData = externalProducts.filter((externalProduct) => {
      return Object.entries(chosenOption).every(([optionKey, optionValue]) => {
        return externalProduct.option[optionKey] === optionValue;
      });
    });

    setChoiceData(filteredChoiceData);
  }, [chosenOption]);

  const handleChangeOption = (key: string, value: string) => {
    setChosenOption((prevOptions) => ({
      ...prevOptions,
      [key]: value,
    }));
  };

  const handleChangeChoiceQuantity = (
    id_external_product: number,
    id_ecommerce: number,
    quantity: string,
  ) => {
    const parsedQuantity = parseInt(quantity, 10);

    if (isNaN(parsedQuantity) || parsedQuantity <= 0) {
      return;
    }

    const updatedChoiceData = choiceData.map((item) =>
      item.id_external_product === id_external_product &&
      item.id_ecommerce === id_ecommerce
        ? { ...item, quantity: parsedQuantity }
        : item,
    );
    setChoiceData(updatedChoiceData);
  };

  const handleAddChosen = (
    id_external_product: number,
    id_ecommerce: number,
  ) => {
    if (!chosenOption) return;
    const itemToAdd = choiceData.find(
      (item) =>
        item.id_external_product === id_external_product &&
        item.id_ecommerce === id_ecommerce,
    );

    if (!itemToAdd || !itemToAdd.stock) return;

    const updatedChoiceData = choiceData.filter(
      (item) =>
        !(
          item.id_external_product === id_external_product &&
          item.id_ecommerce === id_ecommerce
        ),
    );

    setChosenData((prevChosenData) => [
      ...prevChosenData,
      {
        ...itemToAdd,
        option: chosenOption,
      },
    ]);

    setChoiceData(updatedChoiceData);
  };

  const choiceColumns: GridColDef[] = [
    {
      field: "ecommerce",
      headerName: "Sàn thương mại",
    },
    {
      field: "price",
      headerName: "Giá",
    },
    {
      field: "stock",
      headerName: "Tổng số lượng",
    },
    {
      field: "action",
      headerName: "Chọn số lượng",
      renderCell: ({ row }) => (
        <FormControl>
          <Stack
            spacing={1}
            direction="row"
            sx={{ justifyContent: "center", alignItems: "center" }}
          >
            <TextField
              label="Chọn số lượng"
              variant="standard"
              size="small"
              onChange={(event) =>
                handleChangeChoiceQuantity(
                  row.id_external_product,
                  row.id_ecommerce,
                  event.target.value,
                )
              }
            />
            <IconButton
              onClick={() =>
                handleAddChosen(row.id_external_product, row.id_ecommerce)
              }
            >
              <Add />
            </IconButton>
          </Stack>
        </FormControl>
      ),
    },
  ];

  return (
    <Grid container spacing={2}>
      <Grid item xs={12} lg={4}>
        <ProductInfo
          name="Product Name"
          option={option_titles}
          chosenOption={chosenOption}
          handleChangeOption={handleChangeOption}
        />
      </Grid>
      <Grid item xs={12} lg={8}>
        <Paper sx={{ p: 2 }}>
          <DataGrid
            rows={choiceData}
            columns={choiceColumns}
            getRowId={({ id_external_product, id_ecommerce }) =>
              `${id_external_product}-${id_ecommerce}`
            }
          ></DataGrid>
        </Paper>
      </Grid>
      <Grid item xs={12}>
        <Paper sx={{ p: 2 }}>
          <DataGrid
            rows={chosenData}
            columns={chosenColumns}
            getRowId={({ id_external_product, id_ecommerce }) =>
              `${id_external_product}-${id_ecommerce}`
            }
          ></DataGrid>
        </Paper>
      </Grid>
    </Grid>
  );
};

export default ProductDetail;
