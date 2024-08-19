"use client";

import { Button, Input, Select, Space, Flex, Avatar } from "antd";
import { UserOutlined, BellFilled, ShoppingFilled } from "@ant-design/icons";
import { useRouter } from "next/navigation";
import Link from "next/link";

import { Layout } from "antd";
const Header = () => {
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
          <Avatar size="default" icon={<UserOutlined />} />
        </Space>
      </Flex>
    </Layout.Header>
  );
};

export default Header;
