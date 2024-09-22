import { Button, Form, Input, message, Modal, ModalProps } from "antd";
import { Dispatch, SetStateAction } from "react";
import { login } from "@/api/auth";
import useLoading from "@/hook/loading";
import { useAppDispatch } from "@/store/store";
import { setLogin } from "@/store/auth";
import { decodeJwt } from "@/util/auth";
import { useAppSelector } from "@/store/store";
import { toggleLoginModal } from "@/store/auth";

export default function LoginModal() {
  const dispatch = useAppDispatch();
  const handleLogin = useLoading(
    login,
    "Đăng nhập thành công",
    "Đăng nhập thất bại",
  );
  const { isShowLoginModal } = useAppSelector((state) => state.authReducer);

  const onSubmit = async (values: any) => {
    try {
      const res = await handleLogin(values.username, values.password);
      localStorage.setItem("token", res.token);
      dispatch(toggleLoginModal());
      const userInfo = decodeJwt(res.token);
      dispatch(setLogin(userInfo));
    } catch (err) {}
  };

  return (
    <Modal
      open={isShowLoginModal}
      onCancel={() => dispatch(toggleLoginModal())}
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
