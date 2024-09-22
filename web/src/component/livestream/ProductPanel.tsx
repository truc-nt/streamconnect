import { useGetLivestreamProducts } from "@/hook/livestream";
import ProductItem from "@/component/list_item/ProductItem";
import { List } from "antd";
import LivestreamProductSellingModal from "@/app/livestreams/[id]/component/LivestreamProductSellingModal";
import EditLivestreamProductModal from "@/app/livestreams/[id]/component/EditLivestreamProductModal";
import { useState } from "react";
import { useParams } from "next/navigation";
import { Button, Space, Modal, Flex } from "antd";
import {
  PushpinOutlined,
  PlusCircleOutlined,
  ArrowRightOutlined,
  HeartOutlined,
} from "@ant-design/icons";
import { CheckboxChangeEvent } from "antd/lib/checkbox";
import {
  pinLivestreamProduct,
  IPinLivestreamProduct,
} from "@/api/livestream_product";
import useLoading from "@/hook/loading";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";
import { addLivestreamProduct, ILivestreamProduct } from "@/api/livestream";
import { IChosenLivestreamVariant } from "@/app/seller/livestreams/create/component/LivestreamCreate";
import { useAppSelector } from "@/store/store";
import {
  Constants,
  useMeeting,
  useParticipant,
} from "@videosdk.live/react-sdk";

const ProductPanel = () => {
  const { id: livestreamId } = useParams();

  const { data: livestreamProducts, mutate: getLivestreamProducts } =
    useGetLivestreamProducts(Number(livestreamId));

  const { localParticipant } = useMeeting();
  const { mode } = useParticipant(localParticipant?.id ?? "");

  const { userId } = useAppSelector((state) => state.authReducer);

  const [pinnedLivestreamProductIds, setPinnedLivestreamProductIds] = useState<
    number[]
  >([]);

  const [openAddModal, setOpenAddModal] = useState(false);

  const [livestreamProductId, setLivestreamProductId] = useState<number | null>(
    null,
  );

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
        Number(livestreamId),
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

  return (
    <>
      <div className="flex-1 overflow-y-scroll p-2">
        <List
          grid={{ gutter: [2, 2], column: 1 }}
          dataSource={livestreamProducts || []}
          renderItem={(item, index) => (
            <List.Item style={{ padding: 0 }}>
              <ProductItem
                {...item}
                checked={pinnedLivestreamProductIds.includes(
                  item.id_livestream_product,
                )}
                onClick={() => {
                  setLivestreamProductId(item.id_livestream_product);
                }}
                {...(mode === Constants.modes.CONFERENCE
                  ? {
                      onClickCheckbox: (e) =>
                        handleCheckboxClick(e, item.id_livestream_product),
                    }
                  : {})}
                button={mode === Constants.modes.VIEWER && <HeartOutlined />}
                className={index !== 0 ? "opacity-50" : ""}
              />
            </List.Item>
          )}
          //className="overflow-y-scroll overflow-x-visible p-1"
        />
      </div>
      {livestreamProductId !== null &&
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
      {mode === Constants.modes.CONFERENCE && (
        <Space.Compact className="m-auto">
          {true && (
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
          )}
          <Button
            icon={<PlusCircleOutlined />}
            onClick={() => setOpenAddModal(true)}
          />
          <Button
            icon={<ArrowRightOutlined />}
            onClick={async () => {
              try {
                const updatedLivestreamProducts = [...livestreamProducts];
                const firstElement = updatedLivestreamProducts?.shift();
                updatedLivestreamProducts?.push(firstElement);
                console.log(updatedLivestreamProducts);
                await handlePinLivestreamProduct(updatedLivestreamProducts.map((product, index) => ({
                  ...product,
                  priority: index,
                })));
                getLivestreamProducts();
              } catch (e) {}
            }}
          />
        </Space.Compact>
      )}
      {mode === Constants.modes.CONFERENCE && openAddModal && (
        <Modal
          open={true}
          footer={null}
          onCancel={() => setOpenAddModal(false)}
          centered
          width="80%"
        >
          <ChosenLivestreamVariant
            shopId={userId!}
            initialChosenLivestreamVariants={[]}
            onSubmit={handleSubmitAddLivestreamProduct}
          />
        </Modal>
      )}
    </>
  );
};

export default ProductPanel;
