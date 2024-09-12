import React, { useState } from "react";
import {
  InputNumber,
  Table,
  Space,
  Card,
  TableProps,
  Avatar,
  Flex,
  Typography,
  Divider,
} from "antd";
import Tag from "@/component/core/Tag";
import Image from "next/image";

import PriceInfo from "@/component/info/PriceInfo";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";
import { IBaseOrderItem } from "@/model/order";
import { IBuyOrdersGetRequest } from "@/api/order";

interface IOrderTableProps extends IBuyOrdersGetRequest {}
const OrderTable = ({ shop_name, external_orders }: IOrderTableProps) => {
  const orderItems = external_orders.flatMap(
    ({ id_ecommerce, order_items }) => {
      return order_items.map((order_item) => ({ ...order_item, id_ecommerce }));
    },
  );

  const subTotal = orderItems.reduce(
    (total, { paid_price, quantity }) => total + paid_price * quantity,
    0,
  );

  const internalDiscount = external_orders.reduce(
    (total, { internal_discount }) => total + internal_discount,
    0,
  );

  const externalDiscount = 0;
  const shippingFee = 0;

  const columns: TableProps<IBaseOrderItem>["columns"] = [
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
      dataIndex: "id_ecommerce",
      key: "id_ecommerce",
      render: (_, { id_ecommerce }) => (
        <Avatar
          src={ECOMMERCE_LOGOS[id_ecommerce]}
          alt="Shopify Logo"
          size={40}
        />
      ),
    },
    {
      dataIndex: "option",
      key: "option",
      render: (_, { option }) => (
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
      render: (_, { paid_price, unit_price }) => (
        <span>
          {paid_price !== unit_price && (
            <span className="line-through">{unit_price}</span>
          )}
          {paid_price}
        </span>
      ),
    },
    {
      dataIndex: "quantity",
      key: "quantity",
    },
    {
      render: (_, { quantity, paid_price }) => (
        <span>{paid_price * quantity}</span>
      ),
    },
  ];

  return (
    <div className="relative">
      <Table
        columns={columns}
        dataSource={orderItems}
        rowKey={(row) => row.id_order_item}
        pagination={false}
      />
      <Card
        bordered={false}
        style={{
          position: "absolute",
          top: "100%", // Adjust based on your layout
          left: "0",
          width: "100%",
          marginTop: "-20px", // No space between card and table
          boxShadow: "none",
          display: "flex",
          justifyContent: "flex-end",
        }}
      >
        <div className="w-[400px]">
          <PriceInfo
            {...{ subTotal, internalDiscount, externalDiscount, shippingFee }}
          />
        </div>
      </Card>
    </div>
  );
};

export default OrderTable;
