"use client";
import React from "react";
import {
  Checkbox,
  Card as AntdCard,
  Typography,
  Flex,
  Space,
  Button,
} from "antd";
import Image from "next/image";
import { theme } from "antd";
import { CheckboxChangeEvent } from "antd/es/checkbox";
import { TruckOutlined, TagOutlined, GiftOutlined } from "@ant-design/icons";
import { addUserVoucher, IShopVoucherGetResponse } from "@/api/voucher";
import useLoading from "@/hook/loading";
import {
  Constants,
  useMeeting,
  useParticipant,
} from "@videosdk.live/react-sdk";

interface IVoucherItemProps extends IShopVoucherGetResponse {
  button: React.ReactNode;
}

const VoucherItem = ({
  id_voucher,
  code,
  discount,
  max_discount,
  method,
  type,
  min_purchase,
  start_time,
  end_time,
  is_saved,
  button,
}: IVoucherItemProps) => {
  return (
    <Flex gap="small" align="center" justify="center">
      <AntdCard
        hoverable
        className="w-full h-[120px] p-2 flex items-center"
        cover={
          <div
            style={{
              width: "60px",
              height: "100px",
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              borderRadius: "5px",
            }}
            className="bg-slate-300"
          >
            {type === "item" ? (
              <TagOutlined style={{ fontSize: "30px" }} />
            ) : (
              <GiftOutlined style={{ fontSize: "30px" }} />
            )}
          </div>
        }
        styles={{
          body: {
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
            height: "100%",
            padding: "0 0 0 5px",
            width: "100%",
            maxWidth: "100%",
          },
        }}
      >
        <Flex vertical gap="small">
          <Typography.Text style={{ fontWeight: "bold" }}>
            {`Giảm ${discount}${type === "percentage" ? "%" : "đ"}. ${max_discount ? `Tối đa ${max_discount}đ` : ""}`}
          </Typography.Text>
          <Typography.Text style={{ fontSize: "10px" }}>
            {`Đơn tối thiểu ${min_purchase}đ`}
          </Typography.Text>
          <Typography.Text style={{ fontSize: "0.5rem" }}>
            Thời gian: {new Date(start_time).toLocaleString()} -{" "}
            {new Date(end_time).toLocaleString()}
          </Typography.Text>
        </Flex>
        {button}
      </AntdCard>
    </Flex>
  );
};

export default VoucherItem;
