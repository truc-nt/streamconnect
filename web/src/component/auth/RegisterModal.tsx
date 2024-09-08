import { register } from "@/api/auth";
import { Button, Form, Input, message, Modal } from "antd";
import { Dispatch, SetStateAction } from "react";
import useLoading from "@/hook/loading";

export interface AuthModalProps {
  openModal: boolean;
  setOpenModal: Dispatch<SetStateAction<boolean>>;
}

export default function RegisterModal({
  openModal,
  setOpenModal,
}: AuthModalProps) {
  const handleRegister = useLoading(
    register,
    "Đăng kí thành công",
    "Đăng kí thất bại",
  );

  const onSubmit = async (values: any) => {
    try {
      const request = {
        username: values.username,
        password: values.password,
        email: values.email,
      };
      await handleRegister(request);
      setOpenModal(false);
    } catch (err) {
      console.log(err);
    }
  };
  const onCancel = () => {
    setOpenModal(false);
  };
  return (
    <Modal
      open={openModal}
      onCancel={onCancel}
      title={<div className="text-center">Đăng kí</div>}
      footer={null}
      width={400}
      destroyOnClose={true}
    >
      <Form name="register" onFinish={onSubmit} layout="vertical">
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
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="password"
          label="Password"
          rules={[{ required: true, message: "Xin hãy điền mật khẩu" }]}
          required={false}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Submit
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  );
}
