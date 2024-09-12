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
import { useEffect, useState, useLayoutEffect } from "react";
import LoginModal from "@/component/auth/LoginModal";
import RegisterModal from "@/component/auth/RegisterModal";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { setLogin, setLogout } from "@/store/auth";

import { decodeJwt } from "@/util/auth";
import { toggleLoginModal } from "@/store/auth";

const Header = () => {
  const dispatch = useAppDispatch();
  const [isSignUpModalVisible, setIsSignUpModalVisible] = useState(false);
  const { userId } = useAppSelector((state) => state.authReducer);
  console.log(userId);

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
      key: "divider",
      type: "divider",
    },
    {
      key: "3",
      label: (
        <div
          onClick={() => {
            localStorage.removeItem("token");
            dispatch(setLogout());
          }}
        >
          Đăng xuất
        </div>
      ),
    },
  ];

  const onClickSignUp = () => {
    setIsSignUpModalVisible(true);
  };

  useLayoutEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      const userInfo = decodeJwt(token);
      dispatch(setLogin(userInfo));
    } else {
      dispatch(setLogout());
    }
  }, []);

  return (
    <Layout.Header className="bg-white">
      <Flex className="justify-between items-center" gap="large">
        <Input.Search enterButton />
        {userId === null ? (
          <Space>
            <Button onClick={onClickSignUp}>Đăng ký</Button>
            <RegisterModal
              openModal={isSignUpModalVisible}
              setOpenModal={setIsSignUpModalVisible}
            />

            <Button onClick={() => dispatch(toggleLoginModal())} type="primary">
              Đăng nhập
            </Button>
            <LoginModal />
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
