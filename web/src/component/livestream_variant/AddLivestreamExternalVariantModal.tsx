import { useGetProductsByShopId } from "@/hook/product";
import { getVariantsByProductId } from "@/api/product";
import { useState, useEffect } from "react";
import {
  IProductInformation,
  IExternalVariant,
} from "@/app/seller/livestreams/create/component/ProductInformation";
import ProductInfo from "@/component/info/ProductInfo";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";
import {
  TableColumnType,
  Modal,
  InputNumber,
  Table,
  Flex,
  Select,
  Image,
  Button,
  ModalProps,
} from "antd";

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

interface IAddLivestreamExternalVariantProps extends ModalProps {
  shopId: number;
  chosenLivestreamVariants: IChosenLivestreamVariant[];
  setChosenLivestreamVariants: React.Dispatch<
    React.SetStateAction<IChosenLivestreamVariant[]>
  >;
}

const AddLivestreamVariantModal = ({
  shopId,
  chosenLivestreamVariants,
  setChosenLivestreamVariants,
  ...props
}: IAddLivestreamExternalVariantProps) => {
  const { data: products } = useGetProductsByShopId(shopId);

  const [selectedProduct, setSelectedProduct] =
    useState<IProductInformation | null>();

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
      const product = products?.find((product) => product.id_product === value);
      if (!product) return;
      setSelectedProduct({
        productId: product.id_product,
        name: product.name,
        image_url: product.image_url,
        variants: res,
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
    setChosenLivestreamVariants((prev) => {
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
        const quantity = chosenLivestreamVariants
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
      centered
      title="Danh sách sản phẩm"
      open={true}
      footer={null}
      width="60%"
      {...props}
    >
      <Flex vertical gap="middle">
        <Select
          showSearch
          placeholder="Chọn sản phẩm"
          optionFilterProp="label"
          onChange={handleOnChange}
          //onSearch={onSearch}
          options={
            products?.map((product) => ({
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
        {!selectedProduct && (
          <div className="text-center mt-4 border-2 border-dashed border-gray-300 bg-gray-200 rounded-lg h-[300px]" />
        )}
        {Object.keys(selectedProduct ?? {}).length > 0 && (
          <ProductInfo
            product={{
              productId: selectedProduct?.productId!,
              name: selectedProduct?.name!,
              image_url: selectedProduct?.image_url!,
              variants:
                selectedProduct?.variants?.map((variant) => ({
                  id_variant: variant.id_variant,
                  option: variant.option,
                })) ?? [],
            }}
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

export default AddLivestreamVariantModal;
