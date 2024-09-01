import {login, signUp, SignUpRequest} from "@/api/auth";
import {Button, Form, Input, message, Modal} from "antd";
import {AuthModalProps} from "@/component/auth/LoginModal";

export default function SignUpModal({openModal, setOpenModal}: AuthModalProps) {
    const onSubmit = async (values: any) => {
        try {
            const request: SignUpRequest = {
                username: values.username,
                password: values.password,
                email: values.email,
                fullName: values.fullName
            }
            await signUp(request);
            message.success("Đăng kí thành công!");
            setOpenModal(false)
        } catch (err) {
            console.log(err);
        }

    }
    const onCancel = () => {
        setOpenModal(false);
    }
    return (
        <Modal open={openModal} onCancel={onCancel} title={<div style={{ textAlign: "center", marginBottom: "10px" }}>Đăng kí</div>}
                footer={null} width={400} destroyOnClose={true}
        >
            <Form
                name="register"
                autoComplete="off"
                onFinish={onSubmit}
            >
                <Form.Item<any>
                    style={{marginBottom: 15}}
                    name="fullName"
                    rules={[{ required: true, message: 'Please input your full name!' }]}
                >
                    <Input placeholder={"Full Name"} />
                </Form.Item>

                <Form.Item<any>
                    style={{marginBottom: 15}}
                    name="email"
                    rules={[{ required: true, message: 'Please input your email!' }, { type: 'email', message: 'Please enter a valid email!' }]}
                >
                    <Input placeholder={"Email"} />
                </Form.Item>

                <Form.Item<any>
                    style={{marginBottom: 15}}
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input placeholder={"Username"} />
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