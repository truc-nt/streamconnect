"use client";
import React from "react";
import { Card as AntdCard, Typography, Flex, Space } from "antd";
import Image from "next/image";
import { theme } from "antd";
import { IBaseLivestreamProduct } from "@/model/livestream";
import { CheckboxChangeEvent } from "antd/es/checkbox";

interface IProductItem extends IBaseLivestreamProduct {
  onClick: () => void;
  className?: string;
  Button?: React.ReactNode;
  Checkbox?: React.ReactNode;
}

const ProductItem = ({
  name,
  min_price,
  max_price,
  image_url,
  onClick,
  className,
  Button,
  Checkbox,
}: IProductItem) => {
  const { token } = theme.useToken();

  const handleClick = (e: React.MouseEvent) => {
    // Check if the click event is from a checkbox
    if (e.target instanceof HTMLInputElement && e.target.type === "checkbox") {
      return; // Prevent triggering the card's onClick
    }
    onClick(); // Trigger the card's onClick
  };

  return (
    <Flex gap="small" align="center" justify="center">
      <AntdCard
        onClick={handleClick}
        hoverable
        className={`w-full h-[120px] p-2 flex items-center ${className}`}
        cover={
          <div
            style={{
              display: "flex",
              flexDirection: "row",
              justifyContent: "space-between",
              gap: "5px",
            }}
          >
            {Checkbox}
            <div
              style={{ height: "100px", width: "100px" }}
              className="relative"
            >
              <Image
                src={image_url}
                alt={name}
                layout="fill"
                objectFit="contain"
              />
            </div>
          </div>
        }
        styles={{
          body: {
            display: "flex",
            flexDirection: "row",
            justifyContent: "space-between",
            height: "100%",
            padding: "0 0 0 20px",
            width: "100%",
            maxWidth: "100%",
          },
        }}
      >
        <Flex gap="small" justify="space-around" vertical>
          <div className="h-[60%] overflow-hidden text-ellipsis">
            <Typography.Text>{name}</Typography.Text>
          </div>
          <div>
            <Typography.Text
              style={{
                color: token.colorPrimaryText,
                fontWeight: "bold",
                margin: 0,
              }}
            >
              {min_price !== max_price
                ? `${min_price} - ${max_price}`
                : min_price}
            </Typography.Text>
          </div>
        </Flex>
        {Button}
      </AntdCard>
    </Flex>
  );
};

export default ProductItem;
