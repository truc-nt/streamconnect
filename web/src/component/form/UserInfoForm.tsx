"use client";
import { Form, Input, Button, Select, Row, Col, DatePicker } from "antd";
import { useGetUser } from "@/hook/user";
import { useEffect } from "react";
import { updateUser } from "@/api/user";
import useLoading from "@/hook/loading";
import dayjs from "dayjs";

const UserInfoForm = () => {
  const [form] = Form.useForm();
  const { data: user } = useGetUser();
  const executeUpdateUser = useLoading(
    updateUser,
    "Cập nhật thông tin người dùng thành công",
    "Cập nhật thông tin người dùng thất bại",
  );

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();

      const request = {
        ...values,
      };

      console.log(request);

      await executeUpdateUser(request);
    } catch (e) {}
  };

  useEffect(() => {
    if (user) {
      form.setFieldsValue({
        username: user.username,
        email: user.email,
        gender: user.gender,
        birthdate: user.birthdate ? dayjs(user.birthdate) : null,
      });
    }
  }, [user, form]);

  return (
    <Form form={form} layout="vertical">
      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            name="username"
            label="Username"
            rules={[{ required: true, message: "Xin hãy điền username" }]}
            required={false}
          >
            <Input disabled />
          </Form.Item>
        </Col>
        <Col span={24}>
          <Form.Item
            name="email"
            label="Email"
            rules={[
              { required: true, message: "Xin hãy điền email" },
              { type: "email", message: "Xin hãy điền đúng tên email phù hợp" },
            ]}
            required={false}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            name="gender"
            label="Giới tính"
            required={false}
            initialValue={user?.email}
          >
            <Select
              placeholder="Chọn giới tính"
              optionFilterProp="children"
              value={"Nam"}
              options={[
                { label: "Nam", value: "male" },
                { label: "Nữ", value: "female" },
                { label: "Khác", value: "other" },
              ]}
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item name="birthdate" label="Sinh nhật" required={false}>
            <DatePicker className="w-full" />
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

export default UserInfoForm;
