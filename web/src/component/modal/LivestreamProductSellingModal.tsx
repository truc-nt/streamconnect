import {
  Flex,
  Row,
  Col,
  Typography,
  Button,
  Avatar,
  Space,
  theme,
  InputNumber,
  Modal,
  ModalProps,
  Popconfirm,
} from "antd";
import Image from "next/image";
import { useState, useEffect } from "react";

import ProductInfo from "@/component/info/ProductInfo";
import { ECOMMERCE_PLATFORMS, ECOMMERCE_LOGOS } from "@/constant/ecommerce";
import { getProductOptionFromVariantOptions } from "@/util/product";
import { addToCart } from "@/api/cart";
import { createOrderWithLivestreamExternalVariantId } from "@/api/order";
import useLoading from "@/hook/loading";
import { useGetLivestreamProduct } from "@/hook/livestream_product";
import { useAppSelector } from "@/store/store";

interface IProductInformationProps extends ModalProps {
  livestreamProductId: number;
  handleChangeOption?: (option: { [key: string]: string | number }) => void;
}

const LivestreamProductSellingModal = ({
  livestreamProductId,
  handleChangeOption,
  ...props
}: IProductInformationProps) => {
  const { token } = theme.useToken();
  const { data: livestreamProduct } =
    useGetLivestreamProduct(livestreamProductId);
  const [chosenOption, setChosenOption] = useState<{
    [key: string]: string | number;
  }>({});
  const [quantity, setQuantity] = useState(1);
  const [openPopconfirm, setOpenPopconfirm] = useState(false);
  const handleAddToCart = useLoading(
    addToCart,
    "Thêm vào giỏ thành công",
    "Thêm vào giỏ thất bại",
  );

  const option = getProductOptionFromVariantOptions(
    livestreamProduct?.livestream_variants.map(
      (livestreamVariant) => livestreamVariant.option,
    ) ?? [],
  );

  const ecommerceIds: number[] = [];
  const seenIds = new Set<number>();

  const executeCreateOrderWithLivestreamExternalVariantId = useLoading(
    createOrderWithLivestreamExternalVariantId,
    "Đặt hàng thành công",
    "Đặt hàng thất bại",
  );

  livestreamProduct?.livestream_variants.forEach((livestream_variant) => {
    livestream_variant.livestream_external_variants.forEach(
      (livestream_external_variant) => {
        const id_ecommerce = livestream_external_variant.id_ecommerce;
        if (!seenIds.has(id_ecommerce)) {
          seenIds.add(id_ecommerce);
          ecommerceIds.push(id_ecommerce);
        }
      },
    );
  });

  const livestreamVariant = livestreamProduct?.livestream_variants.find(
    (variant) =>
      Object.entries(variant.option).every(
        ([key, value]) => chosenOption[key] === value,
      ),
  );
  const livestreamExternalVariant =
    livestreamVariant?.livestream_external_variants.find(
      (variant) => variant.id_ecommerce === chosenOption.ecommerceId,
    );

  useEffect(() => {
    setChosenOption({});
  }, [livestreamProduct?.name]);

  console.log(livestreamExternalVariant);

  return (
    <Modal
      {...props}
      centered
      title="Thông tin sản phẩm"
      open={true}
      footer={null}
      width="60%"
    >
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
      >
        <div>
          <Typography.Title level={5}>Sàn thương mại</Typography.Title>
          <Flex gap="small">
            {ecommerceIds.map((ecommerceId) => (
              <Button
                key={ecommerceId}
                size="small"
                type={
                  chosenOption.ecommerceId === ecommerceId
                    ? "primary"
                    : "default"
                }
                onClick={() =>
                  setChosenOption({ ...chosenOption, ecommerceId })
                }
              >
                <Flex align="center">
                  <Avatar
                    src={ECOMMERCE_LOGOS[ecommerceId]}
                    alt="Shopify Logo"
                    size={20}
                    className="me-1"
                  />
                  <span>{ECOMMERCE_PLATFORMS[ecommerceId]}</span>
                </Flex>
              </Button>
            ))}
          </Flex>
        </div>
        <Typography.Title
          level={2}
          style={{ color: token.colorPrimaryText, margin: 0 }}
        >
          {livestreamExternalVariant?.price}
        </Typography.Title>
        <Space size="small">
          <InputNumber
            min={1}
            max={livestreamExternalVariant?.quantity}
            defaultValue={quantity}
            disabled={!livestreamExternalVariant}
            onChange={(value) => setQuantity(value!)}
          />
          <Button
            type="primary"
            disabled={!livestreamExternalVariant}
            onClick={async () => {
              await handleAddToCart([
                {
                  id_livestream_external_variant:
                    livestreamExternalVariant?.id_livestream_external_variant ??
                    0,
                  quantity,
                },
              ]);
            }}
          >
            Thêm vào giỏ
          </Button>
          <Popconfirm
            title="Mua ngay"
            description="Mua ngay sản phẩm"
            open={openPopconfirm}
            onConfirm={async () => {
              await executeCreateOrderWithLivestreamExternalVariantId(
                livestreamExternalVariant?.id_livestream_external_variant!,
              );
            }}
            onCancel={() => setOpenPopconfirm(false)}
          >
            <Button
              type="default"
              disabled={!livestreamExternalVariant}
              onClick={() => setOpenPopconfirm(true)}
            >
              Mua ngay
            </Button>
          </Popconfirm>
        </Space>
      </ProductInfo>
    </Modal>
  );
};

export default LivestreamProductSellingModal;
