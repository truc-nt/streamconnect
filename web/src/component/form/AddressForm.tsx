"use client";
import {
  Form,
  Input,
  InputNumber,
  Button,
  Select,
  Space,
  Flex,
  DatePicker,
  Row,
  Col,
  Switch,
  FormProps,
} from "antd";

const AddressForm = (props: FormProps) => {
  return (
    <Form layout="vertical" name="voucher" {...props}>
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            label="Họ và tên"
            name="name"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            label="Địa chỉ"
            name="address"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            label="Thành phố"
            name="city"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            label="Số điện thoại"
            name="phone"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            label="Địa chỉ mặc định"
            name="is_default"
            rules={[{ required: true }]}
            required={false}
            initialValue={false}
          >
            <Switch />
          </Form.Item>
        </Col>
      </Row>
    </Form>
  );
};

export default AddressForm;
