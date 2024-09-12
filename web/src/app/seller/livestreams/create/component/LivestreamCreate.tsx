"use client";
import React, { useState } from "react";
import { Button, message, Steps, Flex, Card, Form } from "antd";
import LivestreamInformationForm from "./LivestreamInformationForm";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";
import {
  ILivestreamExternalVariant,
  ILivestreamProduct,
  createLivestream,
} from "@/api/livestream";
import useLoading from "@/hook/loading";
import { useRouter } from "next/navigation";

export interface IChosenLivestreamVariant {
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

const LivestreamCreate = ({ shopId }: { shopId: number }) => {
  const router = useRouter();
  const [livestreamInformationForm] = Form.useForm();
  const [chosenLivestreamVariants, setChosenLivestreamVariants] = useState<
    IChosenLivestreamVariant[]
  >([]);

  const handleChangeChosenLivestreamVariants = (
    chosenLivestreamVariants: IChosenLivestreamVariant[],
  ) => {
    setChosenLivestreamVariants(chosenLivestreamVariants);
  };

  const handleCreateLivestream = useLoading(
    createLivestream,
    "Tạo livestream thành công",
    "Tạo livestream thất bại",
  );

  const handleSubmit = async () => {
    try {
      const values = await livestreamInformationForm.validateFields();

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
            priority: livestreamProducts.length,
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

      const livestreamCreateData = {
        title: values.title,
        description: values.description,
        start_time: values.startTimeValue
          ? values.startTimeValue.format("YYYY-MM-DDTHH:mm:ssZ")
          : null,
        livestream_products: livestreamProducts,
      };

      const res = await handleCreateLivestream(shopId, livestreamCreateData);
      if (values.startTimeOption === 1) router.push(`/livestreams/${res.data}`);
    } catch (error) {}
  };

  return (
    <Flex vertical gap="large">
      <Card title="Thông tin">
        <LivestreamInformationForm form={livestreamInformationForm} />
      </Card>
      <Card title="Chọn sản phẩm">
        <ChosenLivestreamVariant
          shopId={1}
          onChange={handleChangeChosenLivestreamVariants}
        />
      </Card>
      <Button type="primary" onClick={handleSubmit}>
        Hoàn thành
      </Button>
    </Flex>
  );
};

export default LivestreamCreate;
