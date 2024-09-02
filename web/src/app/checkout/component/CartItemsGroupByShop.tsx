import React, { useState } from "react";
import {
  InputNumber,
  Table,
  Space,
  TableProps,
  Avatar,
  Flex,
  Input,
  Typography,
} from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ICart, ICartItem } from "@/api/cart";
import Tag from "@/component/core/Tag";
import Image from "next/image";
import type { Key } from "react";
import { setSelectedCartItems } from "@/store/cart_item_ids_selection";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { updateQuantity } from "@/api/cart";
import useLoading from "@/hook/loading";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";

const CartItemsGroupByShop = ({ shop_name, cart_items }: ICart) => {
  const columns: TableProps<ICartItem>["columns"] = [
    {
      title: () => <span>{shop_name}</span>,
      dataIndex: "name",
      key: "name",
      render: (_, { name, image_url }) => (
        <Space size="middle" align="center">
          <Image src={image_url} alt={name} width={50} height={50} />
          <span>{name}</span>
        </Space>
      ),
    },
    {
      dataIndex: "option",
      key: "option",
      render: (option) => (
        <Space.Compact block>
          {Object.entries(option).map(([key, value]) => (
            <Tag key={key} label={`${key}: ${value}`} />
          ))}
        </Space.Compact>
      ),
    },
    {
      dataIndex: "price",
      key: "price",
    },
    {
      dataIndex: "quantity",
      key: "quantity",
    },
    {
      key: "total_price",
      render: (_, { quantity, price }) => <span>{price * quantity}</span>,
    },
  ];

  return (
    <Table
      columns={columns}
      dataSource={cart_items}
      rowKey={(row) => row.id_cart_item}
      pagination={false}
    />
  );
};

export default CartItemsGroupByShop;
