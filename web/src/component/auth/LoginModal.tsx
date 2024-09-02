import { Button, Form, Input, message, Modal, ModalProps } from "antd";
import { Dispatch, SetStateAction } from "react";
import { login } from "@/api/auth";
import useLoading from "@/hook/loading";
export interface AuthModalProps {
  openModal: boolean;
  setOpenModal: Dispatch<SetStateAction<boolean>>;
}
export default function LoginModal({
  openModal,
  setOpenModal,
}: AuthModalProps) {
  const handleLogin = useLoading(
    login,
    "Đăng nhập thành công",
    "Đăng nhập thất bại",
  );

  const onSubmit = async (values: any) => {
    try {
      const res = await handleLogin(values.username, values.password);
      localStorage.setItem("token", res.data.token);
      setOpenModal(false);
    } catch (err) {}
  };
  const onCancel = () => {
    setOpenModal(false);
  };
  return (
    <Modal
      open={openModal}
      onCancel={onCancel}
      title={<div className="text-center">Đăng nhập</div>}
      footer={null}
      width={400}
      centered
    >
      <Form name="login" onFinish={onSubmit} layout="vertical">
        <Form.Item
          name="username"
          label="Username"
          rules={[{ required: true, message: "Xin hãy điền username" }]}
          required={false}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="password"
          label="Mật khẩu"
          rules={[{ required: true, message: "Xin hãy điền mật khẩu" }]}
          required={false}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Đăng nhập
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  );
}
