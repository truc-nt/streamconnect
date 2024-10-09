"use client";
import { Form, Input, Button, Select, Row, Col, DatePicker } from "antd";
import { useEffect } from "react";
import { updateShop } from "@/api/shop";
import { useGetShop } from "@/hook/shop";
import { useAppSelector } from "@/store/store";
import useLoading from "@/hook/loading";

const ShopForm = () => {
  const [form] = Form.useForm();
  const { userId } = useAppSelector((state) => state.authReducer);
  const { data: shop } = useGetShop(userId!);
  const executeUpdateShop = useLoading(
    updateShop,
    "Cập nhật thông tin cửa hàng thành công",
    "Cập nhật thông tin cửa hàng thất bại",
  );

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();

      const request = {
        ...values,
      };

      await executeUpdateShop(userId!, request);
    } catch (e) {}
  };

  useEffect(() => {
    if (shop) {
      form.setFieldsValue({
        name: shop.name,
        description: shop.description,
      });
    }
  }, [shop, form]);

  return (
    <Form form={form} layout="vertical">
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            name="name"
            label="Tên cửa hàng"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            name="description"
            label="Mô tả"
            required={false}
            className="w-full"
          >
            <Input.TextArea placeholder="Mô tả cửa hàng" className="w-full" />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              block
              className="w-full"
              onClick={handleSubmit}
            >
              Cập nhật
            </Button>
          </Form.Item>
        </Col>
      </Row>
    </Form>
  );
};

export default ShopForm;
