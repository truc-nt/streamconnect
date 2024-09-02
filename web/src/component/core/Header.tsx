"use client";

import {
  Button,
  Input,
  Select,
  Space,
  Flex,
  Avatar,
  Dropdown,
  Divider,
} from "antd";
import type { MenuProps } from "antd";
import { UserOutlined, BellFilled, ShoppingFilled } from "@ant-design/icons";
import { useRouter } from "next/navigation";
import Link from "next/link";

import { Layout } from "antd";
const Header = () => {
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
      label: "Đăng xuất",
    },
  ];

  return (
    <Layout.Header className="bg-white">
      <Flex className="justify-between items-center" gap="large">
        <Input.Search enterButton />
        <Space>
          <Button>Đăng ký</Button>
          <Button type="primary">Đăng nhập</Button>
        </Space>
        <Space>
          <Button type="text" shape="circle" icon={<BellFilled />} />
          <Link href="/cart">
            <Button type="text" shape="circle" icon={<ShoppingFilled />} />
          </Link>
          <Dropdown menu={{ items }} placement="bottomLeft">
            <Avatar size="default" icon={<UserOutlined />} />
          </Dropdown>
        </Space>
      </Flex>
    </Layout.Header>
  );
};

export default Header;
