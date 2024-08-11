import DataGrid from "@/components/core/DataGrid";
import { useState, useEffect } from "react";
import {
  Stack,
  Chip,
  Button,
  IconButton,
  TextField,
  Grid,
} from "@mui/material";
import { GridColDef, GridEventListener, useGridApiRef } from "@mui/x-data-grid";
import { ArrowBack } from "@mui/icons-material";
import { useGetProducts } from "@/hook/shop";
import { getVariants, IVariant } from "@/api/product";
import { IProduct } from "@/api/shop";
import ProductInfo from "@/components/product/ProductInfo";
import { useAppDispatch } from "@/store/store";
import { setChosenLivestreamVariants } from "@/store/livestream_create";
import { setOpen } from "@/store/alert";

interface IProductItem extends IProduct {
  isSelected?: boolean;
}

interface ILivestreamVariant {
  idProduct: number;
  idVariant: number;
  name: string;
  option: Record<string, string>;
  externalVariants: IExternalVariant[];
}

interface IExternalVariant {
  idVariant: number;
  idExternalVariant: number;
  idEcommerce: number;
  price: number;
  stock: number;
  quantity: number;
}

const LivestreamVariants = () => {
  const [livestreamVariants, setLivestreamVariants] = useState<
    ILivestreamVariant[]
  >([]);
  const [openChangePage, setChangePage] = useState(false);
  const distpatch = useAppDispatch();

  const handleSubmit = () => {
    distpatch(setChosenLivestreamVariants(livestreamVariants));
    console.log(livestreamVariants);
  };

  return (
    <Stack gap={2}>
      {!openChangePage ? (
        <>
          <Button
            variant="contained"
            sx={{ marginRight: "auto" }}
            onClick={() => setChangePage(true)}
          >
            Thêm sản phẩm
          </Button>
          <LivestreamVariantDataGrid rows={livestreamVariants} />
          <Stack direction="row" gap={1} sx={{ marginLeft: "auto" }}>
            <Button variant="contained" onClick={handleSubmit}>
              Tiếp theo
            </Button>
          </Stack>
        </>
      ) : (
        <>
          <IconButton
            onClick={() => setChangePage(false)}
            sx={{ marginRight: "auto" }}
          >
            <ArrowBack />
          </IconButton>
          <LivestreamVariantsChange
            livestreamVariants={livestreamVariants}
            setLivestreamVariants={setLivestreamVariants}
          />
        </>
      )}
    </Stack>
  );
};

const LivestreamVariantDataGrid = ({
  rows,
}: {
  rows: ILivestreamVariant[];
}) => {
  const apiRef = useGridApiRef();
  useEffect(() => {
    /*apiRef.current.autosizeColumns({
      columns: columns.map((column: any) => column.field),
      expand: true,
    });*/
  }, [rows]);
  const columns: GridColDef<ILivestreamVariant>[] = [
    {
      field: "name",
      headerName: "Tên",
    },
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
      field: "price",
      headerName: "Giá",
      renderCell: ({ row }) => (
        <Stack direction="row" spacing={1}>
          {row.externalVariants.map((externalProduct, index) => (
            <Chip
              key={index}
              label={`Shopify: ${externalProduct.price}`}
              size="small"
            />
          ))}
        </Stack>
      ),
    },
    {
      field: "stock",
      headerName: "Số lượng đã chọn",
      renderCell: ({ row }) => (
        <Stack direction="row" spacing={1}>
          {row.externalVariants.map((externalProduct, index) => (
            <Chip
              key={index}
              label={`Shopify: ${externalProduct.quantity}`}
              size="small"
            />
          ))}
        </Stack>
      ),
    },
  ];
  return (
    <DataGrid
      apiRef={apiRef}
      rows={rows}
      columns={columns}
      getRowId={({ idVariant }) => `${idVariant}`}
      disableRowSelectionOnClick
    />
  );
};

const LivestreamVariantsChange = ({
  livestreamVariants,
  setLivestreamVariants,
}: {
  livestreamVariants: ILivestreamVariant[];
  setLivestreamVariants: React.Dispatch<
    React.SetStateAction<ILivestreamVariant[]>
  >;
}) => {
  const { data: products, error } = useGetProducts(
    process.env.NEXT_PUBLIC_SHOP_ID,
  );
  const [product, setProduct] = useState<IProduct>();
  const [variants, setVariants] = useState<IVariant[]>([]);
  const [chosenVariantOption, setChosenVariantOption] = useState<
    Record<string, string>
  >({});
  const [externalProducts, setExternalProducts] = useState<IExternalVariant[]>(
    [],
  );

  useEffect(() => {
    if (
      !product ||
      Object.keys(chosenVariantOption).length <
        Object.keys(product?.option).length
    ) {
      return;
    }

    const filteredVariants = variants.find((variant) => {
      return Object.entries(chosenVariantOption).every(
        ([optionKey, optionValue]) => {
          return variant.option[optionKey] === optionValue;
        },
      );
    });

    if (!filteredVariants) {
      return;
    }

    const fileteredLivestreamVariants = livestreamVariants.find(
      (item) => item.idProduct === product.id_product,
    );

    setExternalProducts(
      filteredVariants.external_products.map((externalProduct) => ({
        idVariant: filteredVariants.id_variant,
        idExternalVariant: externalProduct.id_external_product,
        idEcommerce: externalProduct.id_ecommerce,
        ecommerce: externalProduct.ecommerce,
        option: filteredVariants.option,
        stock: externalProduct.stock,
        price: externalProduct.price,
        quantity:
          fileteredLivestreamVariants?.externalVariants.find(
            (item) =>
              item.idExternalVariant === externalProduct.id_external_product,
          )?.quantity || 0,
      })),
    );
  }, [chosenVariantOption]);

  const handleChangeOption = (key: string, value: string) => {
    setChosenVariantOption((prev) => ({ ...prev, [key]: value }));
  };

  const handleRowClick: GridEventListener<"rowClick"> = async (params) => {
    try {
      const product = params.row;
      const variants = await getVariants(product.id_product);
      setProduct(product);
      setVariants(variants.data);
      setChosenVariantOption({});
      setExternalProducts([]);
    } catch (error) {
      console.log(error);
    }
  };

  const handleChangeLivestreamVariantQuantity = (
    externalProduct: IExternalVariant,
    quantity: string,
  ) => {
    const parsedQuantity = parseInt(quantity, 10);

    if (isNaN(parsedQuantity) || parsedQuantity <= 0 || !product) {
      return;
    }

    const existingLivestreamVariant = livestreamVariants.findIndex(
      (item) => item.idVariant === externalProduct.idVariant,
    );

    if (existingLivestreamVariant > -1) {
      const existingExternalProduct = livestreamVariants[
        existingLivestreamVariant
      ].externalVariants.findIndex(
        (item) =>
          item.idExternalVariant === externalProduct.idExternalVariant &&
          item.idEcommerce === externalProduct.idEcommerce,
      );

      if (existingExternalProduct > -1) {
        livestreamVariants[existingLivestreamVariant].externalVariants[
          existingExternalProduct
        ].quantity = parsedQuantity;
        setLivestreamVariants([...livestreamVariants]);
        return;
      }

      livestreamVariants[existingLivestreamVariant].externalVariants = [
        ...livestreamVariants[existingLivestreamVariant].externalVariants,
        {
          ...externalProduct,
          quantity: parsedQuantity,
        },
      ];
      setLivestreamVariants([...livestreamVariants]);
    } else {
      setLivestreamVariants([
        ...livestreamVariants,
        {
          idProduct: product.id_product,
          idVariant: externalProduct.idVariant,
          name: product.name,
          option: chosenVariantOption,
          externalVariants: [
            {
              ...externalProduct,
              quantity: parsedQuantity,
            },
          ],
        },
      ]);
    }
  };

  return (
    <>
      {product && variants && (
        <Grid container spacing={2}>
          <Grid item xs={12} sm={6}>
            <ProductInfo
              name={product.name}
              option={product.option}
              chosenOption={chosenVariantOption}
              handleChangeOption={handleChangeOption}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <ExternalProductDataGrid
              rows={externalProducts}
              setRows={setExternalProducts}
              handleChangeLivestreamVariantQuantity={
                handleChangeLivestreamVariantQuantity
              }
            />
          </Grid>
        </Grid>
      )}
      {products && (
        <ProductDataGrid
          rows={products.map((product) => ({
            ...product,
            isSelected: livestreamVariants.some(
              (livestreamVariant) =>
                livestreamVariant.idProduct === product.id_product,
            ),
          }))}
          handleRowClick={handleRowClick}
        />
      )}
    </>
  );
};

const ExternalProductDataGrid = ({
  rows,
  setRows,
  handleChangeLivestreamVariantQuantity,
}: {
  rows: IExternalVariant[];
  setRows: React.Dispatch<React.SetStateAction<IExternalVariant[]>>;
  handleChangeLivestreamVariantQuantity: (
    externalProduct: IExternalVariant,
    quantity: string,
  ) => void;
}) => {
  const externalProductChoiceColumns: GridColDef<IExternalVariant>[] = [
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
      headerName: "Số lượng",
    },
    {
      field: "quantity",
      headerName: "Số lượng",
      renderCell: ({ row }) => (
        <TextField
          type="number"
          value={row.quantity}
          onChange={(e) => {
            handleChangeLivestreamVariantQuantity(row, e.target.value);
            setRows((prev) =>
              prev.map((item) =>
                item.idExternalVariant === row.idExternalVariant &&
                item.idEcommerce === row.idEcommerce
                  ? { ...item, quantity: parseInt(e.target.value, 10) }
                  : item,
              ),
            );
          }}
        />
      ),
    },
  ];

  return (
    <DataGrid
      rows={rows}
      columns={externalProductChoiceColumns}
      getRowId={(row) => row.idExternalVariant}
      disableRowSelectionOnClick
    ></DataGrid>
  );
};

const ProductDataGrid = ({
  rows,
  handleRowClick,
}: {
  rows: IProductItem[];
  handleRowClick: GridEventListener<"rowClick">;
}) => {
  const apiRef = useGridApiRef();

  useEffect(() => {
    rows?.forEach((row) => {
      if (row.isSelected) {
        apiRef.current.selectRow(row.id_product, true);
      }
    });
    /*apiRef.current.autosizeColumns({
      columns: columns.map((column: any) => column.field),
      expand: true,
    });*/
  }, [rows]);

  const columns: GridColDef<IProductItem>[] = [
    {
      field: "name",
      headerName: "Tên",
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
      field: "price",
      headerName: "Giá",
      renderCell: ({ row }) => (
        <span>
          {row.min_price} - {row.max_price}
        </span>
      ),
    },
    {
      field: "total_stock",
      headerName: "Số lượng",
    },
  ];
  return (
    <DataGrid
      checkboxSelection
      apiRef={apiRef}
      rows={rows}
      columns={columns}
      getRowId={(row) => row.id_product}
      onRowClick={handleRowClick}
      disableRowSelectionOnClick
    />
  );
};

export default LivestreamVariants;
