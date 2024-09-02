"use client";
import React from "react";
import { Card as AntdCard, Typography, Flex, Space } from "antd";
import Image from "next/image";
import { theme } from "antd";
import { ILivestreamProduct } from "@/api/livestream";

interface ICard extends ILivestreamProduct {
  onClick: () => void;
}

const Card = ({ name, min_price, max_price, image_url, onClick }: ICard) => {
  const { token } = theme.useToken();

  return (
    <AntdCard
      onClick={onClick}
      hoverable
      className="w-full h-[160px] p-5 flex items-center"
      cover={
        <div style={{ height: "120px", width: "120px" }} className="relative">
          <Image src={image_url} alt={name} layout="fill" objectFit="contain" />
        </div>
      }
      styles={{
        body: {
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
          height: "100%",
          padding: "0 0 0 20px",
          width: "100%",
          maxWidth: "100%",
        },
      }}
    >
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
          {min_price !== max_price ? `${min_price} - ${max_price}` : min_price}
        </Typography.Text>
      </div>
    </AntdCard>
  );
};

export default Card;
