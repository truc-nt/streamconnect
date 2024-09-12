"use client";
import React from "react";
import { Checkbox, Card as AntdCard, Typography, Flex, Space } from "antd";
import Image from "next/image";
import { theme } from "antd";
import { ILivestreamProductInformation } from "@/api/livestream";
import { CheckboxChangeEvent } from "antd/es/checkbox";

interface IProductItem extends ILivestreamProductInformation {
  checked?: boolean;
  onClick: () => void;
  onClickCheckbox?: (e: CheckboxChangeEvent) => void;
  button?: React.ReactNode;
  className?: string;
}

const ProductItem = ({
  name,
  min_price,
  max_price,
  image_url,
  checked,
  onClick,
  onClickCheckbox,
  button,
  className,
}: IProductItem) => {
  const { token } = theme.useToken();

  return (
    <Flex gap="small" align="center" justify="center">
      {onClickCheckbox && (
        <Checkbox checked={checked} onChange={onClickCheckbox} />
      )}
      <AntdCard
        onClick={onClick}
        hoverable
        className={`w-full h-[120px] p-2 flex items-center ${className}`}
        cover={
          <div style={{ height: "100px", width: "100px" }} className="relative">
            <Image
              src={image_url}
              alt={name}
              layout="fill"
              objectFit="contain"
            />
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
        {button}
      </AntdCard>
    </Flex>
  );
};

export default ProductItem;
