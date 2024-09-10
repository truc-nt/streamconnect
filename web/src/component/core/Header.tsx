"use client";

import {
  Button,
  Input,
  Select,
  Space,
  Flex,
  Avatar,
  Modal,
  Form,
  Dropdown,
} from "antd";
import { UserOutlined, BellFilled, ShoppingFilled } from "@ant-design/icons";
import { useRouter } from "next/navigation";
import Link from "next/link";

import { Layout, MenuProps } from "antd";
import { useEffect, useState } from "react";
import LoginModal from "@/component/auth/LoginModal";
import RegisterModal from "@/component/auth/RegisterModal";

const Header = () => {
  const [isSignInModalVisible, setIsSignInModalVisible] = useState(false);
  const [isSignUpModalVisible, setIsSignUpModalVisible] = useState(false);
  const [isAuthorized, setIsAuthorized] = useState(false);

  const items: MenuProps["items"] = [
    {
      key: "1",
      label: <Link href="/user">Hồ sơ</Link>,
    },
    {
      key: "2",
      label: <Link href="/seller">Cửa hàng</Link>,
    },
    {
      key: "3",
      label: (
        <Button
          onClick={() => {
            localStorage.removeItem("token");
            setIsAuthorized(false);
          }}
        >
          Đăng xuất
        </Button>
      ),
    },
  ];

  const onClickSignIn = () => {
    setIsSignInModalVisible(true);
  };
  const onClickSignUp = () => {
    setIsSignUpModalVisible(true);
  };
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!!token) setIsAuthorized(true);
  }, [isSignInModalVisible, isSignUpModalVisible, isAuthorized]);
  return (
    <Layout.Header className="bg-white">
      <Flex className="justify-between items-center" gap="large">
        <Input.Search enterButton />
        {!isAuthorized ? (
          <Space>
            <Button onClick={onClickSignUp}>Đăng ký</Button>
            <RegisterModal
              openModal={isSignUpModalVisible}
              setOpenModal={setIsSignUpModalVisible}
            />

            <Button onClick={onClickSignIn} type="primary">
              Đăng nhập
            </Button>
            <LoginModal
              openModal={isSignInModalVisible}
              setOpenModal={setIsSignInModalVisible}
            />
          </Space>
        ) : (
          <Space>
            <Button type="text" shape="circle" icon={<BellFilled />} />
            <Link href="/cart">
              <Button type="text" shape="circle" icon={<ShoppingFilled />} />
            </Link>
            <Dropdown menu={{ items }} placement="bottomLeft">
              <Avatar size="default" icon={<UserOutlined />} />
            </Dropdown>
          </Space>
        )}
      </Flex>
    </Layout.Header>
  );
};

export default Header;
