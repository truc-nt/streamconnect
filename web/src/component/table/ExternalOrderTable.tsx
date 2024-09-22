"use client";
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
import { IOrderGetRequest } from "@/api/order";

interface IExternalOrderTableProps extends IOrderGetRequest {}
const ExternalOrderTable = ({
  id_external_order,
  id_ecommerce,
  external_order_id_mapping,
  shipping_fee,
  internal_discount,
  external_discount,
  order_items,
}: IOrderGetRequest) => {
  const subTotal = order_items.reduce(
    (total, { paid_price, quantity }) => total + paid_price * quantity,
    0,
  );
  const columns: TableProps<IBaseOrderItem>["columns"] = [
    {
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
      render: () => (
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
        dataSource={order_items}
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
            subTotal={subTotal}
            shippingFee={shipping_fee}
            internalDiscount={internal_discount}
            externalDiscount={external_discount}
          />
        </div>
      </Card>
    </div>
  );
};

export default ExternalOrderTable;
