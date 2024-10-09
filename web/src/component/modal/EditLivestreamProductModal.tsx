import { useGetLivestreamProduct } from "@/hook/livestream_product";
import {
  ILivestreamExternalVariant,
  updateLivestreamExternalVariantQuantity,
  IUpdateLivestreamExternalVariantQuantity,
} from "@/api/livestream_product";
import { useState, useEffect } from "react";
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
  Space,
} from "antd";
import useLoading from "@/hook/loading";

interface IEditLivestreamProductModalProps extends ModalProps {
  livestreamProductId: number;
}

const EditLivestreamProductModal = ({
  livestreamProductId,
  ...props
}: IEditLivestreamProductModalProps) => {
  const { data: livestreamProduct } =
    useGetLivestreamProduct(livestreamProductId);

  const [chosenOption, setChosenOption] = useState<{ [key: string]: string }>(
    {},
  );
  const chosenVariant = livestreamProduct?.livestream_variants?.find(
    (variant) =>
      Object.entries(variant.option).every(
        ([key, value]) => chosenOption[key] === value,
      ),
  );

  const handleUpdateLivestreamExternalVariantQuantity = useLoading(
    updateLivestreamExternalVariantQuantity,
    "Cập nhật số lượng thành công",
    "Cập nhật số lượng thất bại",
  );

  const [
    updatedLivestreamExternalVariantQuantity,
    setUpdatedLivestreamExternalVariantQuantity,
  ] = useState<IUpdateLivestreamExternalVariantQuantity[]>([]);

  const columns: TableColumnType<ILivestreamExternalVariant>[] = [
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
      title: "Số lượng đã chọn",
      dataIndex: "quantity",
      key: "quantity",
    },
    {
      title: "Số lượng chọn",
      dataIndex: "quantity",
      key: "quantity",
      render: (_, { quantity, stock, id_livestream_external_variant }) => {
        return (
          <InputNumber
            min={0}
            max={stock}
            defaultValue={quantity ?? 0}
            onChange={(value) => {
              if (!value) {
                return;
              }

              setUpdatedLivestreamExternalVariantQuantity((prev) => {
                const index = prev.findIndex(
                  (variant) =>
                    variant.id_livestream_external_variant ===
                    id_livestream_external_variant,
                );
                if (index !== -1) {
                  prev[index].quantity = value;
                } else {
                  prev.push({
                    id_livestream_external_variant,
                    quantity: value,
                  });
                }
                return [...prev];
              });
            }}
          />
        );
      },
    },
  ];

  return (
    <Modal
      centered
      title="Thông tin sản phẩm"
      open={true}
      footer={null}
      width="60%"
      {...props}
    >
      <Flex vertical gap="middle">
        <ProductInfo
          product={{
            productId: livestreamProduct?.id_product!,
            name: livestreamProduct?.name!,
            image_url: livestreamProduct?.image_url!,
            variants:
              livestreamProduct?.livestream_variants?.map((variant) => ({
                id_variant: variant.id_variant,
                option: variant.option,
              })) ?? [],
          }}
          handleChangeOption={(option) => setChosenOption(option)}
        />

        {chosenVariant?.livestream_external_variants !== undefined && (
          <Table
            columns={columns}
            dataSource={chosenVariant?.livestream_external_variants}
            pagination={false}
            rowKey={(record) => record.id_livestream_external_variant}
          />
        )}
        <Flex gap="small" className="w-full">
          <Button
            type="primary"
            className="flex-1"
            disabled={updatedLivestreamExternalVariantQuantity.length === 0}
            onClick={async () => {
              try {
                await handleUpdateLivestreamExternalVariantQuantity(
                  updatedLivestreamExternalVariantQuantity,
                );
              } catch (error) {}
            }}
          >
            Lưu
          </Button>
        </Flex>
      </Flex>
    </Modal>
  );
};

export default EditLivestreamProductModal;
