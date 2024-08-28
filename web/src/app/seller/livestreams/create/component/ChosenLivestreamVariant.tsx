import {
  Button,
  Flex,
  Table,
  Modal,
  Select,
  InputNumber,
  TableColumnType,
  ModalProps,
  Tag,
} from "antd";
import React, { useState } from "react";
import { useGetProductsByShopId } from "@/hook/product";
import { getVariantsByProductId } from "@/api/product";
import Image from "next/image";
import ProductInformation, {
  IExternalVariant,
  IProductInformation,
} from "@/app/seller/livestreams/create/component/ProductInformation";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";
import { setChosenLivestreamVariants } from "@/store/livestream_create";
import { useAppDispatch, useAppSelector } from "@/store/store";

interface IAddLivestreamVariantProps extends ModalProps {
  chosenLivestreamVariant: IChosenLivestreamVariant[];
  setChosenLivestreamVariant: React.Dispatch<
    React.SetStateAction<IChosenLivestreamVariant[]>
  >;
}
const AddLivestreamVariant = ({
  chosenLivestreamVariant,
  setChosenLivestreamVariant,
  ...props
}: IAddLivestreamVariantProps) => {
  const { data: products } = useGetProductsByShopId(1);
  const [selectedProduct, setSelectedProduct] =
    useState<IProductInformation | null>(null);
  const [chosenOption, setChosenOption] = useState<{ [key: string]: string }>(
    {},
  );
  const chosenVariant = selectedProduct?.variants?.find((variant) =>
    Object.entries(variant.option).every(
      ([key, value]) => chosenOption[key] === value,
    ),
  );

  const handleOnChange = async (value: number) => {
    try {
      const res = await getVariantsByProductId(value);
      const product = products?.data.find(
        (product) => product.id_product === value,
      );
      if (!product) return;
      setSelectedProduct({
        productId: product.id_product,
        name: product.name,
        image_url: product.image_url,
        variants: res.data,
      });
    } catch (error) {
      console.error(error);
    }
  };

  const handleChangeQuantity = (
    externalVariantId: number,
    ecommerceId: number,
    quantity: number,
    price: number,
  ) => {
    setChosenLivestreamVariant((prev) => {
      const index = prev.findIndex((variant) =>
        variant.externalVariants.some(
          (externalVariant) =>
            externalVariant.externalVariantId === externalVariantId,
        ),
      );
      if (index === -1) {
        return [
          ...prev,
          {
            productId: selectedProduct?.productId!,
            variantId: chosenVariant?.id_variant,
            name: selectedProduct?.name,
            imageUrl: selectedProduct?.image_url,
            option: chosenOption,
            externalVariants: [
              {
                externalVariantId,
                ecommerceId,
                quantity,
                price,
              },
            ],
          } as IChosenLivestreamVariant,
        ];
      }
      return prev.map((variant, i) => {
        if (i === index) {
          return {
            ...variant,
            externalVariants: variant.externalVariants.map(
              (externalVariant) => {
                if (externalVariant.externalVariantId === externalVariantId) {
                  return {
                    ...externalVariant,
                    quantity,
                  };
                }
                return externalVariant;
              },
            ),
          };
        }
        return variant;
      });
    });
  };

  const columns: TableColumnType<IExternalVariant>[] = [
    {
      title: "Sàn thương mại",
      dataIndex: "id_ecommerce",
      key: "id_ecommerce",
      render: (_, { id_ecommerce }) => {
        return <span>{ECOMMERCE_PLATFORMS[id_ecommerce]}</span>;
      },
    },
    {
      title: "Giá",
      dataIndex: "price",
      key: "price",
    },
    {
      title: "Số lượng",
      dataIndex: "stock",
      key: "stock",
    },
    {
      title: "Số lượng chọn",
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { id_external_variant, id_ecommerce, stock, price }) => {
        const quantity = chosenLivestreamVariant
          .find((variant) =>
            variant.externalVariants?.some(
              (externalVariant) =>
                externalVariant.externalVariantId === id_external_variant,
            ),
          )
          ?.externalVariants.find(
            (externalVariant) =>
              externalVariant.externalVariantId === id_external_variant,
          )?.quantity;

        return (
          <InputNumber
            min={0}
            max={stock}
            defaultValue={quantity ?? 0}
            onChange={(value) =>
              handleChangeQuantity(
                id_external_variant,
                id_ecommerce,
                value!,
                price,
              )
            }
          />
        );
      },
    },
  ];

  return (
    <Modal
      {...props}
      centered
      title="Danh sách sản phẩm"
      open={true}
      footer={null}
      width="60%"
    >
      <Flex vertical gap="middle">
        <Select
          showSearch
          placeholder="Chọn sản phẩm"
          optionFilterProp="label"
          onChange={handleOnChange}
          //onSearch={onSearch}
          options={
            products?.data.map((product) => ({
              label: product.name,
              value: product.id_product,
              image_url: product.image_url,
            })) ?? []
          }
          optionRender={(option) => (
            <Flex align="center" gap="small">
              <Image
                src={option.data.image_url}
                alt={option.data.label}
                width={50}
                height={50}
              />
              {option.data.label}
            </Flex>
          )}
        />
        {Object.keys(selectedProduct ?? {}).length > 0 && (
          <ProductInformation
            product={selectedProduct!}
            handleChangeOption={(option) => setChosenOption(option)}
          />
        )}
        {chosenVariant?.external_variants !== undefined && (
          <Table
            columns={columns}
            dataSource={chosenVariant?.external_variants}
            pagination={false}
            rowKey={(record) => record.id_external_variant}
          />
        )}
      </Flex>
    </Modal>
  );
};

interface IChosenLivestreamVariant {
  productId: number;
  variantId: number;
  name: string;
  imageUrl: string;
  option: { [key: string]: string };
  externalVariants: {
    externalVariantId: number;
    ecommerceId: number;
    price: number;
    quantity: number;
  }[];
}

const ChosenLivestreamVariant = () => {
  const distpatch = useAppDispatch();
  const { livestreamExternalVariants } = useAppSelector(
    (state) => state.livestreamCreateReducer,
  );
  const [chosenLivestreamVariant, setChosenLivestreamVariant] = useState<
    IChosenLivestreamVariant[]
  >(livestreamExternalVariants);
  const [openAddModal, setOpenAddModal] = useState(false);

  const columns: TableColumnType<IChosenLivestreamVariant>[] = [
    {
      title: "Tên sản phẩm",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Option",
      dataIndex: "option",
      key: "option",
      render: (_, { option }) =>
        Object.entries(option).map(([key, value]) => (
          <Tag key={key}>
            {key}: {value}
          </Tag>
        )),
    },
    {
      title: "Giá",
      dataIndex: "price",
      key: "price",
      render: (_, { externalVariants }) =>
        Object.values(externalVariants).map((externalVariant) => (
          <Tag key={externalVariant.externalVariantId}>
            {ECOMMERCE_PLATFORMS[externalVariant.ecommerceId]}:{" "}
            {externalVariant.price}
          </Tag>
        )),
    },
    {
      title: "Số lượng",
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { externalVariants }) =>
        Object.values(externalVariants).map((externalVariant) => (
          <Tag key={externalVariant.externalVariantId}>
            {ECOMMERCE_PLATFORMS[externalVariant.ecommerceId]}:{" "}
            {externalVariant.quantity}
          </Tag>
        )),
    },
  ];

  const handleSubmit = () => {
    distpatch(setChosenLivestreamVariants(chosenLivestreamVariant));
  };
  return (
    <Flex vertical gap="large">
      <Flex>
        <Button type="primary" onClick={() => setOpenAddModal(true)}>
          Thêm sản phẩm
        </Button>
      </Flex>
      <Table
        columns={columns}
        dataSource={chosenLivestreamVariant}
        rowKey={(row) => row.variantId}
      />
      {openAddModal && (
        <AddLivestreamVariant
          chosenLivestreamVariant={chosenLivestreamVariant}
          setChosenLivestreamVariant={setChosenLivestreamVariant}
          onCancel={() => setOpenAddModal(false)}
        />
      )}
      <Flex justify="end">
        <Button
          type="primary"
          disabled={Object.keys(chosenLivestreamVariant).length === 0}
          onClick={handleSubmit}
        >
          Tiếp theo
        </Button>
      </Flex>
    </Flex>
  );
};

export default ChosenLivestreamVariant;
