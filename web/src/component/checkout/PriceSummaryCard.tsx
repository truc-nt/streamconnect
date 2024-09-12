import { Avatar, Card, Divider, Flex, Typography, Button } from "antd";

import PriceInfo from "@/component/info/PriceInfo";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";

const PriceSummaryCard = ({
  ecommerceId,
  subTotal,
  internalDiscountTotal,
  externalDiscountTotal,
  shippingFee,
}: {
  ecommerceId?: number;
  subTotal: number;
  internalDiscountTotal: number;
  externalDiscountTotal: number;
  shippingFee: number;
}) => (
  <Card
    bordered={false}
    styles={{
      body: {
        display: "flex",
        flexDirection: "column",
        height: "100%",
        gap: "0.5rem",
      },
    }}
  >
    {!ecommerceId && <Card.Meta title="Tổng tiền" />}
    {ecommerceId && (
      <Flex align="center">
        <Avatar
          src={ECOMMERCE_LOGOS[ecommerceId]}
          alt="Shopify Logo"
          size={40}
        />
        <Typography.Title level={5} style={{ margin: 0 }}>
          Tổng tiền
        </Typography.Title>
      </Flex>
    )}
    <PriceInfo
      subTotal={subTotal}
      internalDiscount={internalDiscountTotal}
      externalDiscount={externalDiscountTotal}
      shippingFee={shippingFee}
    />
  </Card>
);

export default PriceSummaryCard;
