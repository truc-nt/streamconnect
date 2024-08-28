"use client";
import { Space, Flex, Table, Button, InputNumber, Input } from "antd";
import { PlusOutlined, MinusOutlined } from "@ant-design/icons";
import { useGetCart } from "@/hook/cart";
import CartGroupByShop from "./component/CartGroupByShop";

const Page = () => {
  const data = [
    {
      key: "1",
      name: "John Brown",
      price: 32,
      quantity: 2,
      action: "Edit",
    },
  ];

  const { data: cart, error } = useGetCart(1);

  return (
    <Space.Compact direction="vertical" style={{ display: "flex" }}>
      <div></div>
      <Flex vertical>
        {cart?.data.map((cartGroupByShop, index) => (
          <CartGroupByShop key={index} {...cartGroupByShop} />
        ))}
      </Flex>
    </Space.Compact>
  );
};

export default Page;
