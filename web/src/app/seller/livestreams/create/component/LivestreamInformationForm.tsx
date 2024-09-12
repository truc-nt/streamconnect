import {
  Form,
  Input,
  Button,
  Checkbox,
  Flex,
  DatePicker,
  FormProps,
  Radio,
  Space,
} from "antd";
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setLivestreamInformation, reset } from "@/store/livestream_create";

const LivestreamInformationForm = (props: FormProps) => {
  return (
    <Form
      name="login"
      layout="vertical"
      initialValues={{ remember: true }}
      {...props}
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
      <Form.Item
        label="Thời gian bắt đầu"
        name="startTimeOption"
        rules={[{ required: true, message: "Hãy chọn thời gian bắt đầu" }]}
      >
        <Radio.Group>
          <Space direction="vertical">
            <Radio value={1}>Bắt đầu ngay bây giờ</Radio>
            <Flex align="center">
              <Radio value={2}>Lên lịch</Radio>
              <Form.Item name="startTimeValue" className="m-0">
                <DatePicker
                  showTime
                  placeholder="Chọn thời gian bắt đầu"
                  format="YYYY-MM-DD HH:mm:ss"
                />
              </Form.Item>
            </Flex>
          </Space>
        </Radio.Group>
      </Form.Item>
    </Form>
  );
};

export default LivestreamInformationForm;
