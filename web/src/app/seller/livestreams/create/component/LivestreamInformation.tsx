import { Form, Input, Button, Checkbox, Flex, DatePicker } from "antd";
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import { useAppSelector, useAppDispatch } from "@/store/store";
import {
  setPrevStep,
  setLivestreamInformation,
  reset,
} from "@/store/livestream_create";
import { ILivestreamExternalVariant, createLivestream } from "@/api/livestream";
import useLoading from "@/hook/loading";

const LivestreamInformation = () => {
  const dispatch = useAppDispatch();
  const { title, description, startTime, livestreamExternalVariants } =
    useAppSelector((state: any) => state.livestreamCreateReducer);

  const handleCreateLivestream = useLoading(
    createLivestream,
    "Tạo livestream thành công",
    "Tạo livestream thất bại",
  );

  const handleSubmit = async (values: any) => {
    dispatch(
      setLivestreamInformation({
        title: values.title,
        description: values.description,
        startTime: values.startTime?.format("YYYY-MM-DD HH:mm:ss") ?? "",
      }),
    );

    const livestreamProducts: {
      id_product: number;
      priority: number;
      livestream_variants: {
        id_variant: number;
        livestream_external_variants: {
          id_external_variant: number;
          quantity: number;
        }[];
      }[];
    }[] = [];
    for (const _livestreamExternalVariant of livestreamExternalVariants) {
      const { productId, variantId, externalVariants } =
        _livestreamExternalVariant;

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
      title: title,
      description: description,
      start_time: startTime,
      livestream_products: livestreamProducts,
    };

    try {
      await handleCreateLivestream(1, livestreamCreateData);
      dispatch(reset());
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <Form
      name="login"
      layout="vertical"
      initialValues={{ remember: true }}
      onFinish={handleSubmit}
    >
      <Form.Item
        label="Tiêu đề"
        name="title"
        rules={[{ required: true, message: "Hãy điền tiêu đề cho livestream" }]}
      >
        <Input.TextArea rows={2} placeholder="Tiêu đề cho livestream" />
      </Form.Item>
      <Form.Item label="Mô tả" name="description">
        <Input.TextArea rows={4} placeholder="Mô tả cho livestream" />
      </Form.Item>
      <Form.Item label="Thời gian bắt đầu" name="startTime">
        <DatePicker showTime placeholder="Chọn thời gian bắt đầu" />
      </Form.Item>
      <Flex justify="end" gap="middle">
        <Form.Item>
          <Button onClick={() => dispatch(setPrevStep())}>Quay lại</Button>
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Hoàn thành
          </Button>
        </Form.Item>
      </Flex>
    </Form>
  );
};

export default LivestreamInformation;
