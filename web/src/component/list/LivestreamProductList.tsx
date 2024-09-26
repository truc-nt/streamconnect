"use client";
import { useState } from "react";

import { Checkbox, List, Modal, Button, Space, Flex, Typography } from "antd";
import { CheckboxChangeEvent } from "antd/lib/checkbox";
import { Constants } from "@videosdk.live/react-sdk";
import {
  HeartOutlined,
  PlusCircleOutlined,
  PushpinOutlined,
  ArrowRightOutlined,
} from "@ant-design/icons";

import LivestreamProductSellingModal from "@/app/livestreams/[id]/component/LivestreamProductSellingModal";
import EditLivestreamProductModal from "@/app/livestreams/[id]/component/EditLivestreamProductModal";
import ProductItem from "@/component/list_item/ProductItem";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";
import useLoading from "@/hook/loading";
import { useGetLivestreamProducts } from "@/hook/livestream";
import {
  updateLivestreamProductPriority,
  IPinLivestreamProduct,
} from "@/api/livestream_product";
import { notifyLivestreamProductFollowers } from "@/api/notification";
import { addLivestreamProduct, ILivestreamProduct } from "@/api/livestream";
import {
  registerLivestreamProductFollower,
  updateLivestreamProduct,
} from "@/api/livestream_product";
import { IChosenLivestreamVariant } from "@/app/seller/livestreams/create/component/LivestreamCreate";
import { LivestreamStatus } from "@/constant/livestream";

const LivestreamProductList = ({
  shopId,
  livestreamId,
  mode,
  status,
}: {
  shopId: number;
  livestreamId: number;
  mode: "CONFERENCE" | "VIEWER";
  status: LivestreamStatus;
}) => {
  const { data: livestreamProducts, mutate: getLivestreamProducts } =
    useGetLivestreamProducts(Number(livestreamId)) ?? [];
  const livestreamProductList = [
    livestreamProducts?.filter(
      (livestreamProduct) => !livestreamProduct.is_livestreamed,
    ) ?? [],
    livestreamProducts?.filter(
      (livestreamProduct) => livestreamProduct.is_livestreamed,
    ) ?? [],
  ];

  const [pinnedLivestreamProductIds, setPinnedLivestreamProductIds] = useState<
    number[]
  >([]);
  const [livestreamProductId, setLivestreamProductId] = useState<number | null>(
    null,
  );
  const [openAddModal, setOpenAddModal] = useState(false);

  const executeUpdateLivestreamProductPriority = useLoading(
    updateLivestreamProductPriority,
    "Đã ghim sản phẩm thành công",
    "Ghim sản phẩm thất bại",
  );
  const executeAddLivestreamProduct = useLoading(
    addLivestreamProduct,
    "Thêm sản phẩm thành công",
    "Thêm sản phẩm thất bại",
  );
  const executeRegisterLivestreamProductFollower = useLoading(
    registerLivestreamProductFollower,
    "Đăng ký sản phẩm thành công",
    "Đăng ký sản phẩm thất bại",
  );

  const executeNotifyLivestreamProductFollowers = useLoading(
    notifyLivestreamProductFollowers,
  );

  const executeUpdateLivestreamProduct = useLoading(updateLivestreamProduct);

  const handleAddLivestreamProduct = async (
    chosenLivestreamVariants: IChosenLivestreamVariant[],
  ) => {
    try {
      const livestreamProducts: ILivestreamProduct[] = [];
      for (const chosenLivestreamVariant of chosenLivestreamVariants) {
        const { productId, variantId, externalVariants } =
          chosenLivestreamVariant;

        let livestreamProductIndex = livestreamProducts.findIndex(
          (product) => product.id_product === productId,
        );

        if (livestreamProductIndex === -1) {
          livestreamProducts.push({
            id_product: productId,
            priority:
              livestreamProducts.length + (livestreamProducts.length || 0),
            livestream_variants: [],
          });
          livestreamProductIndex = livestreamProducts.length - 1;
        }

        let livestreamVariantIndex = livestreamProducts[
          livestreamProductIndex
        ].livestream_variants.findIndex(
          (variant) => variant.id_variant === variantId,
        );
        if (livestreamVariantIndex === -1) {
          livestreamProducts[livestreamProductIndex].livestream_variants.push({
            id_variant: variantId,
            livestream_external_variants: [],
          });
          livestreamVariantIndex =
            livestreamProducts[livestreamProductIndex].livestream_variants
              .length - 1;
        }

        livestreamProducts[livestreamProductIndex].livestream_variants[
          livestreamVariantIndex
        ].livestream_external_variants.push(
          ...externalVariants.map((externalVariant) => ({
            id_external_variant: externalVariant.externalVariantId,
            quantity: externalVariant.quantity,
          })),
        );
      }

      await executeAddLivestreamProduct(
        Number(livestreamId),
        livestreamProducts,
      );
      setOpenAddModal(false);
      getLivestreamProducts();
    } catch (e) {}
  };
  const handleCheckboxClick = (e: CheckboxChangeEvent, productId: number) => {
    e.stopPropagation();
    if (e.target.checked) {
      setPinnedLivestreamProductIds((prev) => [...prev, productId]);
    } else {
      setPinnedLivestreamProductIds((prev) =>
        prev.filter((id) => id !== productId),
      );
    }
  };
  const handlePinLivestreamProduct = async () => {
    try {
      const pinLivestreamProduct: IPinLivestreamProduct[] =
        pinnedLivestreamProductIds.map((id, index) => ({
          id_livestream_product: id,
          priority: index,
          is_livestreamed: false,
        }));

      const unpinLivestreamProduct: IPinLivestreamProduct[] = (
        livestreamProducts || []
      )
        .filter(
          (product) =>
            !pinnedLivestreamProductIds.includes(product.id_livestream_product),
        )
        .map((product, index) => ({
          id_livestream_product: product.id_livestream_product,
          priority: pinnedLivestreamProductIds.length + index,
        }));

      await executeUpdateLivestreamProduct(livestreamId, [
        ...pinLivestreamProduct,
        ...unpinLivestreamProduct,
      ]);
      await getLivestreamProducts();
      await executeNotifyLivestreamProductFollowers(
        livestreamProducts?.[0]?.id_livestream_product,
      );
      setPinnedLivestreamProductIds([]);
    } catch (e) {}
  };
  const handleNextLivestreamProduct = async () => {
    const newLivestreamProducts = [
      ...livestreamProducts?.slice(1)!,
      livestreamProducts?.[0],
    ];
    await executeUpdateLivestreamProduct(
      livestreamId,
      newLivestreamProducts.map((livestreamProduct, index) => ({
        ...livestreamProduct,
        priority: index,
        ...(index === newLivestreamProducts.length - 1
          ? { is_livestreamed: true }
          : {}),
      })),
    );

    await getLivestreamProducts();

    await executeNotifyLivestreamProductFollowers(
      livestreamProducts?.[0]?.id_livestream_product,
    );
  };

  return (
    <>
      <div className="flex-1 overflow-y-scroll p-2">
        <List
          itemLayout="horizontal"
          dataSource={livestreamProductList}
          renderItem={(livestreamProducts, index) => (
            <List.Item style={{ width: "100%" }}>
              <List
                grid={{ gutter: [2, 2], column: 1 }}
                dataSource={livestreamProducts}
                renderItem={(item, index) => (
                  <List.Item style={{ padding: 0, width: "100%" }}>
                    <ProductItem
                      {...item}
                      onClick={() => {
                        setLivestreamProductId(item.id_livestream_product);
                      }}
                      {...(mode === Constants.modes.CONFERENCE
                        ? {
                            Checkbox: (
                              <Checkbox
                                checked={pinnedLivestreamProductIds.includes(
                                  item.id_livestream_product,
                                )}
                                onChange={(e) => {
                                  e.stopPropagation();
                                  handleCheckboxClick(
                                    e,
                                    item.id_livestream_product,
                                  );
                                }}
                              />
                            ),
                          }
                        : {})}
                      Button={
                        mode === Constants.modes.VIEWER && (
                          <HeartOutlined
                            onClick={async (e) => {
                              e.stopPropagation();
                              await executeRegisterLivestreamProductFollower(
                                livestreamId,
                                [item.id_livestream_product],
                              );
                            }}
                          />
                        )
                      }
                      className={item.is_livestreamed ? "opacity-50" : ""}
                    />
                  </List.Item>
                )}
              />
            </List.Item>
          )}
        />
      </div>
      {mode === Constants.modes.CONFERENCE &&
        (status == LivestreamStatus.CREATED ||
          status == LivestreamStatus.PLAYED) && (
          <Space.Compact className="m-auto">
            {true && (
              <Button
                icon={<PushpinOutlined />}
                disabled={pinnedLivestreamProductIds.length === 0}
                onClick={handlePinLivestreamProduct}
              />
            )}
            <Button
              icon={<PlusCircleOutlined />}
              onClick={() => setOpenAddModal(true)}
            />
            <Button
              icon={<ArrowRightOutlined />}
              onClick={async () =>
                //await executeNotifyLivestreamProductFollowers(1)
                handleNextLivestreamProduct()
              }
            />
          </Space.Compact>
        )}
      {livestreamProductId !== null &&
        status !== LivestreamStatus.ENDED &&
        (mode === Constants.modes.CONFERENCE ? (
          <EditLivestreamProductModal
            livestreamProductId={livestreamProductId}
            onCancel={() => setLivestreamProductId(null)}
          />
        ) : (
          <LivestreamProductSellingModal
            livestreamProductId={livestreamProductId}
            onCancel={() => setLivestreamProductId(null)}
          />
        ))}
      {mode === Constants.modes.CONFERENCE &&
        (status == LivestreamStatus.CREATED ||
          status == LivestreamStatus.PLAYED) && (
          <Modal
            open={openAddModal}
            footer={null}
            onCancel={() => setOpenAddModal(false)}
            centered
            width="80%"
          >
            <ChosenLivestreamVariant
              shopId={shopId}
              initialChosenLivestreamVariants={[]}
              onSubmit={handleAddLivestreamProduct}
            />
          </Modal>
        )}
    </>
  );
};

export default LivestreamProductList;
