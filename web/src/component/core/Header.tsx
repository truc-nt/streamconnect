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
import {disconnectSocket} from "@/api/socket";
import {useWebSocket} from "@/hook/socket";
import NotificationDropdown from "@/component/core/NotificationMenu";
import {batchUpdateNotificationStatus, getNotifications, Notification} from "@/api/notification";

const Header = () => {

  const [isSignInModalVisible, setIsSignInModalVisible] = useState(false);
  const [isSignUpModalVisible, setIsSignUpModalVisible] = useState(false);
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [newNotificationCount, setNewNotificationCount] = useState(0);

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
            disconnectSocket();
            localStorage.removeItem("token");
            setIsAuthorized(false);
          }}
        >
          Đăng xuất
        </Button>
      ),
    },
  ];

  //handle function
  const onClickSignIn = () => {
    setIsSignInModalVisible(true);
  };
  const onClickSignUp = () => {
    setIsSignUpModalVisible(true);
  };
  const onClickOpenNotificationMenu = () => {
    const ids = notifications.map((notification) => notification.id);
    setNewNotificationCount(0);
    batchUpdateNotificationStatus(ids, "SEND");
  }
  const onNotificationReceive = (notification: Notification) => {
    setNotifications((prevNotifications) => [notification, ...prevNotifications]);
  }

  //use effect section
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!!token) setIsAuthorized(true);
  }, [isSignInModalVisible, isSignUpModalVisible, isAuthorized]);
  useEffect(() => {
    const fetchNotifications = async () => {
      const notifications = await getNotifications();
      setNotifications(notifications);
      // handle the fetched notifications
    };
    fetchNotifications();
  }, []);
  useEffect(() => {
    setNewNotificationCount(notifications.filter((notification) => notification.status === "NEW").length);
  }, [notifications]);

  //use hook section
  useWebSocket(onNotificationReceive);

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
            <NotificationDropdown items={notifications} newItemCount={newNotificationCount}>
              <Button onClick={onClickOpenNotificationMenu} type="text" shape="circle" icon={<BellFilled />} />
            </NotificationDropdown>
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
