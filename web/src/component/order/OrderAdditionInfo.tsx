"use client";
import { Card, Flex, Typography } from "antd";

import { IBaseOrderAdditionInfo } from "@/model/order";
import { SHIPPING_METHODS, PAYMENT_METHODS } from "@/constant/order";

interface IOrderAdditionInfoProps extends IBaseOrderAdditionInfo {}

const OrderAdditionInfo = ({
  name,
  phone,
  address,
  city,
  id_shipping_method,
  id_payment_method,
}: IOrderAdditionInfoProps) => {
  return (
    <Flex gap="large">
      <Card
        title="Thông tin người nhận"
        className="flex-1"
        styles={{
          body: {
            display: "flex",
            flexDirection: "column",
            gap: "0.5rem",
          },
        }}
      >
        <Typography.Text style={{ fontWeight: "bold" }}>{name}</Typography.Text>
        <Typography.Text>Số điện thoại: {phone}</Typography.Text>
        <Typography.Text>
          Địa chỉ: {address} {city}
        </Typography.Text>
      </Card>
      <Card title="Phương thức vận chuyển" className="flex-1">
        <Typography.Text>
          Giao hàng {SHIPPING_METHODS[id_shipping_method]}
        </Typography.Text>
      </Card>
      <Card title="Phương thức thanh toán" className="flex-1">
        <Typography.Text>{PAYMENT_METHODS[id_payment_method]}</Typography.Text>
      </Card>
    </Flex>
  );
};

export default OrderAdditionInfo;
