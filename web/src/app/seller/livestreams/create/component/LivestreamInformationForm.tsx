"use client";
import { useState } from "react";

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
import dayjs from "dayjs";

const LivestreamInformationForm = (props: FormProps) => {
  const [isStartTimeRequired, setIsStartTimeRequired] = useState(false);
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
        <Radio.Group
          onChange={(e) => setIsStartTimeRequired(e.target.value === 2)}
        >
          <Space direction="vertical">
            <Radio value={1}>Bắt đầu ngay bây giờ</Radio>
            <Flex align="center">
              <Radio value={2}>Lên lịch</Radio>
              <Form.Item
                name="startTimeValue"
                className="m-0"
                required={isStartTimeRequired}
                rules={[{ required: isStartTimeRequired }]}
              >
                <DatePicker
                  showTime
                  placeholder="Chọn thời gian bắt đầu"
                  format="YYYY-MM-DD HH:mm:ss"
                  disabledDate={(current) =>
                    current && current < dayjs().startOf("day")
                  }
                  disabledTime={(current) => ({
                    disabledHours: () => [],
                    disabledMinutes: (selectedHour) => {
                      if (selectedHour === dayjs().hour()) {
                        return Array.from(
                          { length: dayjs().minute() },
                          (_, i) => i,
                        );
                      }
                      return [];
                    },
                    disabledSeconds: () => [],
                  })}
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
