"use client";
import { Form, Input, Button, Row, Col, FormProps } from "antd";
import { useEffect } from "react";
import { IBaseProduct } from "@/model/product";
import { updateProduct } from "@/api/product";
import useLoading from "@/hook/loading";

const ProductForm = (product: IBaseProduct) => {
  const [form] = Form.useForm();
  const executeUpdateProduct = useLoading(
    updateProduct,
    "Cập nhật thông tin sản phẩm thành công",
    "Cập nhật thông tin sản phẩm thất bại",
  );

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();

      await executeUpdateProduct(product.id_product, { ...values });
    } catch (e) {}
  };

  useEffect(() => {
    if (product) {
      form.setFieldsValue({
        name: product.name,
        description: product.description,
      });
    }
  }, [product]);

  return (
    <Form form={form} layout="vertical">
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            name="name"
            label="Tên"
            required={false}
            className="w-full"
          >
            <Input placeholder="Tên sản phẩm" className="w-full" />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            name="description"
            label="Mô tả"
            required={false}
            className="w-full"
          >
            <Input.TextArea placeholder="Mô tả sản phẩm" className="w-full" />
          </Form.Item>
        </Col>
        <Form.Item className="w-full">
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
      </Row>
    </Form>
  );
};

export default ProductForm;
