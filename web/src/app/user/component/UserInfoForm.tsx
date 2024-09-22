"use client";
import { Form, Input, Button } from "antd";
import { useGetUserInfo } from "@/hook/user";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { useEffect } from "react";

const UserInfoForm = () => {
  const [form] = Form.useForm();
  const { data } = useGetUserInfo();

  useEffect(() => {
    if (data) {
      form.setFieldsValue({
        username: data.username,
        email: data.email,
      });
    }
  }, [data, form]);

  return (
    <Form form={form} layout="vertical">
      <Form.Item
        name="username"
        label="Username"
        rules={[{ required: true, message: "Xin hãy điền username" }]}
        required={false}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="email"
        label="Email"
        rules={[
          { required: true, message: "Xin hãy điền email" },
          { type: "email", message: "Xin hãy điền đúng tên email phù hợp" },
        ]}
        required={false}
        initialValue={data?.email}
      >
        <Input />
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit" block>
          Cập nhật
        </Button>
      </Form.Item>
    </Form>
  );
};

export default UserInfoForm;
