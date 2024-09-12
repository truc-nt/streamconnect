"use client";

import { useParams } from "next/navigation";
import { useState } from "react";

import { Card, List, Calendar, Tag, Flex, Space, Button, Modal } from "antd";
import { CheckboxChangeEvent } from "antd/lib/checkbox";
import {
  HeartOutlined,
  PushpinOutlined,
  PlusCircleOutlined,
} from "@ant-design/icons";
import type { CalendarProps } from "antd";
import type { Dayjs } from "dayjs";
import dayjs from "dayjs";

import LivestreamInfoItem from "@/component/list_item/LivestreamInfoItem";
import LivestreamProductSellingModal from "@/app/livestreams/[id]/component/LivestreamProductSellingModal";
import EditLivestreamProductModal from "@/app/livestreams/[id]/component/EditLivestreamProductModal";
import ProductItem from "@/component/list_item/ProductItem";
import { useGetAllLivestreams } from "@/hook/livestream";
import { useGetLivestreamProducts } from "@/hook/livestream";
import useLoading from "@/hook/loading";
import {
  pinLivestreamProduct,
  IPinLivestreamProduct,
} from "@/api/livestream_product";
import { addLivestreamProduct, ILivestreamProduct } from "@/api/livestream";
import { IChosenLivestreamVariant } from "@/app/seller/livestreams/create/component/LivestreamCreate";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";

const LivestreamCalendar = ({
  shopId,
  editable,
}: {
  shopId: number;
  editable: boolean;
}) => {
  const { data: livestreams } = useGetAllLivestreams(shopId);
  const [selectedLivestreamId, setSelectedLivestreamId] = useState<
    number | null
  >(null);
  const [livestreamProductId, setLivestreamProductId] = useState<number | null>(
    null,
  );
  const { data: livestreamProducts, mutate: getLivestreamProducts } =
    useGetLivestreamProducts(Number(selectedLivestreamId)) ?? [];

  const [openAddModal, setOpenAddModal] = useState(false);
  const [pinnedLivestreamProductIds, setPinnedLivestreamProductIds] = useState<
    number[]
  >([]);
  const handlePinLivestreamProduct = useLoading(
    pinLivestreamProduct,
    "Đã ghim sản phẩm thành công",
    "Ghim sản phẩm thất bại",
  );
  const _handleSubmitAddLivestreamProduct = useLoading(
    addLivestreamProduct,
    "Thêm sản phẩm thành công",
    "Thêm sản phẩm thất bại",
  );

  const handleSubmitAddLivestreamProduct = async (
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

      await _handleSubmitAddLivestreamProduct(
        Number(selectedLivestreamId),
        livestreamProducts,
      );
      setOpenAddModal(false);
      getLivestreamProducts();
    } catch (e) {}
  };

  const handleCheckboxClick = (e: CheckboxChangeEvent, productId: number) => {
    if (e.target.checked) {
      setPinnedLivestreamProductIds((prev) => [...prev, productId]);
    } else {
      setPinnedLivestreamProductIds((prev) =>
        prev.filter((id) => id !== productId),
      );
    }
  };

  const getListData = (value: Dayjs) => {
    const dateString = value.format("YYYY-MM-DD");
    return livestreams
      ?.filter(
        (livestream) =>
          dayjs(livestream.start_time).format("YYYY-MM-DD") === dateString,
      )
      .map((livestream) => ({
        title: livestream.title,
        livestreamId: livestream.id_livestream,
      }));
  };

  const dateCellRender = (value: Dayjs) => {
    const listData = getListData(value);
    return (
      <div>
        {listData?.map((livestream, index) => (
          <Tag
            key={index}
            color="blue"
            onClick={() => setSelectedLivestreamId(livestream.livestreamId)}
          >
            {livestream.title}
          </Tag>
        ))}
      </div>
    );
  };

  const cellRender: CalendarProps<Dayjs>["cellRender"] = (current, info) => {
    if (info.type === "date") return dateCellRender(current);
    return info.originNode;
  };
  return (
    <Flex gap="small">
      <Calendar cellRender={cellRender} />
      {selectedLivestreamId && (
        <Card
          className="h-full w-full flex flex-col"
          title="Các sản phẩm trong phiên livestream"
          styles={{
            body: {
              flex: "1 1 0%",
              display: "flex",
              flexDirection: "column",
              gap: "0.25rem",
              padding: "1rem",
              justifyContent: "center",
            },
          }}
        >
          <div className="flex-1 overflow-y-scroll p-2">
            <List
              grid={{ gutter: [2, 2], column: 1 }}
              dataSource={livestreamProducts}
              renderItem={(item) => (
                <List.Item style={{ padding: 0 }}>
                  <ProductItem
                    {...item}
                    checked={pinnedLivestreamProductIds.includes(
                      item.id_livestream_product,
                    )}
                    onClick={() => {
                      setLivestreamProductId(item.id_livestream_product);
                      console.log(item.id_livestream_product);
                    }}
                    button={!editable && <HeartOutlined />}
                    {...(editable
                      ? {
                          onClickCheckbox: (e) =>
                            handleCheckboxClick(e, item.id_livestream_product),
                        }
                      : {})}
                  />
                </List.Item>
              )}
              //className="overflow-y-scroll overflow-x-visible p-1"
            />
            {editable && (
              <Space.Compact className="m-auto">
                <Button
                  icon={<PushpinOutlined />}
                  disabled={pinnedLivestreamProductIds.length === 0}
                  onClick={async () => {
                    try {
                      const pinLivestreamProduct: IPinLivestreamProduct[] =
                        pinnedLivestreamProductIds.map((id, index) => ({
                          id_livestream_product: id,
                          priority: index,
                        }));

                      const unpinLivestreamProduct: IPinLivestreamProduct[] = (
                        livestreamProducts || []
                      )
                        .filter(
                          (product) =>
                            !pinnedLivestreamProductIds.includes(
                              product.id_livestream_product,
                            ),
                        )
                        .map((product, index) => ({
                          id_livestream_product: product.id_livestream_product,
                          priority:
                            product.priority +
                            pinnedLivestreamProductIds.length +
                            index,
                        }));

                      await handlePinLivestreamProduct([
                        ...pinLivestreamProduct,
                        ...unpinLivestreamProduct,
                      ]);
                      getLivestreamProducts();
                      setPinnedLivestreamProductIds([]);
                    } catch (e) {}
                  }}
                />
                <Button
                  icon={<PlusCircleOutlined />}
                  onClick={() => setOpenAddModal(true)}
                />
              </Space.Compact>
            )}
          </div>
        </Card>
      )}
      {livestreamProductId !== null &&
        (editable ? (
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
      {editable && openAddModal && (
        <Modal
          open={true}
          footer={null}
          onCancel={() => setOpenAddModal(false)}
          centered
          width="80%"
        >
          <ChosenLivestreamVariant
            shopId={shopId}
            initialChosenLivestreamVariants={[]}
            onSubmit={handleSubmitAddLivestreamProduct}
          />
        </Modal>
      )}
    </Flex>
  );
};

export default LivestreamCalendar;
