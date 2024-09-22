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
  FormProps,
} from "antd";
import { formatPrice } from "@/util/format";

const VoucherForm = (props: FormProps) => {
  return (
    <Form layout="vertical" name="voucher" {...props}>
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            label="Mã code"
            name="code"
            rules={[{ required: true }]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Giảm"
            name="discount"
            rules={[{ required: true }]}
            required={false}
            initialValue={0}
          >
            <InputNumber className="w-full" />
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Kiểu"
            name="type"
            rules={[{ required: true }]}
            required={false}
          >
            <Select>
              <Select.Option value="percentage">Phần trăm</Select.Option>
              <Select.Option value="fixed">Cố định</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={8}>
          <Form.Item
            label="Giảm tối đa"
            name="max_discount"
            rules={[{ required: true }]}
            required={false}
            initialValue={0}
          >
            <InputNumber className="w-full" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Phạm vi"
            name="target"
            rules={[{ required: true }]}
            required={false}
          >
            <Select>
              <Select.Option value="item">Sản phẩm</Select.Option>
              <Select.Option value="shipping">Vận chuyển</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Phương thức"
            name="method"
            rules={[{ required: true }]}
            required={false}
          >
            <Select>
              <Select.Option value="across">Toàn bộ</Select.Option>
              <Select.Option value="each">Riêng lẻ</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Tiền hàng tối thiểu"
            name="min_purchase"
            required={false}
            initialValue={0}
          >
            <InputNumber className="w-full" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            label="Số lượng"
            name="quantity"
            rules={[{ required: true }]}
            required={false}
            initialValue={0}
          >
            <InputNumber className="w-full" />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item label="Thời gian bắt đầu" name="start_time">
            <DatePicker
              showTime
              placeholder="Chọn thời gian bắt đầu"
              format="YYYY-MM-DD HH:mm:ss"
              className="w-full"
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item label="Thời gian kết thúc" name="end_time">
            <DatePicker
              showTime
              placeholder="Chọn thời gian kết thúc"
              format="YYYY-MM-DD HH:mm:ss"
              className="w-full"
            />
          </Form.Item>
        </Col>
      </Row>
    </Form>
  );
};

export default VoucherForm;
