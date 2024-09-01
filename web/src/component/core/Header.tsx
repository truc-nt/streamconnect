"use client";

import {Button, Input, Select, Space, Flex, Avatar, Modal, Form, Dropdown} from "antd";
import { UserOutlined, BellFilled, ShoppingFilled } from "@ant-design/icons";
import Link from "next/link";
import styles from "./Header.module.css";

import { Layout } from "antd";
import {useEffect, useState} from "react";
import LoginModal from "@/component/auth/LoginModal";
import SignUpModal from "@/component/auth/SignUpModal";
import ActionMenu from "@/component/core/ActionMenu";
const Header = () => {
  const [isSignInModalVisible, setIsSignInModalVisible] = useState(false);
  const [isSignUpModalVisible, setIsSignUpModalVisible] = useState(false);
  const [isAuthorized, setIsAuthorized] = useState(false);
  const onClickSignIn = () => {
    setIsSignInModalVisible(true);
  }
  const onClickSignUp = () => {
    setIsSignUpModalVisible(true);
  }
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!!token) setIsAuthorized(true);
  });
  return (
    <Layout.Header className="bg-white">
      <Flex className="justify-between items-center" gap="large">
        <Input.Search enterButton />
        {!isAuthorized ? <Space>
          <Button onClick={onClickSignUp}>Đăng ký</Button>
          <SignUpModal openModal={isSignUpModalVisible} setOpenModal={setIsSignUpModalVisible} />

          <Button onClick={onClickSignIn} type="primary">Đăng nhập</Button>
          <LoginModal openModal={isSignInModalVisible} setOpenModal={setIsSignInModalVisible} />
        </Space> : <></>}
        <Space>
          <Button type="text" shape="circle" icon={<BellFilled />} />
          <Link href="/cart">
            <Button type="text" shape="circle" icon={<ShoppingFilled />} />
          </Link>
          <Dropdown dropdownRender={ActionMenu} trigger={['click']}>
            <Avatar size="default" icon={<UserOutlined />} className={styles.avatar} />
          </Dropdown>
        </Space>
      </Flex>
    </Layout.Header>
  );
};

export default Header;
