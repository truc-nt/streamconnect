import {Button, Form, Input, message, Modal} from "antd";
import {Dispatch, SetStateAction} from "react";
import {login} from "@/api/auth";
export interface AuthModalProps {
  openModal: boolean;
  setOpenModal:  Dispatch<SetStateAction<boolean>>
}
export default function LoginModal({openModal, setOpenModal}: AuthModalProps) {
  const onSubmit = async (values: any) => {
    try {
      const response = await login(values.username, values.password);
      localStorage.setItem("token", response.data.token);
      message.success("Đăng nhập thành công");
      setOpenModal(false)
    } catch (err) {
      console.log(err);
    }

  }
  const onCancel = () => {
    setOpenModal(false);
  }
  return (
      <Modal  open={openModal} onCancel={onCancel} title={<div style={{ textAlign: "center", marginBottom: "10px" }}>Đăng nhập</div>}
             footer={null} width={400} destroyOnClose={true}
      >
        <Form
            name="basic"
            autoComplete="off"
            onFinish={onSubmit}
        >
          <Form.Item<any>
              style={{marginBottom: 15}}
              name="username"
              rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input placeholder={"Username / Email"} />
          </Form.Item>

          <Form.Item<any>
              style={{marginBottom: 15}}
              name="password"
              rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password placeholder={"Password"}/>
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 9, span: 16 }} style={{marginBottom: 0}}>
            <Button style={{margin: 0}} type="primary" htmlType="submit">
              Submit
            </Button>
          </Form.Item>
        </Form>
      </Modal>
  )
}