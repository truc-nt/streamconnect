import { Card, Divider, Flex, Typography, Button, Avatar } from "antd";
import { ECOMMERCE_LOGOS } from "@/constant/ecommerce";

const PriceSummary = ({
  subTotal,
  shippingFee,
  discountTotal,
  ecommerceId,
}: {
  subTotal: number;
  shippingFee: number;
  discountTotal: number;
  ecommerceId?: number;
}) => {
  console.log(ecommerceId);
  return (
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
      <Flex justify="space-between">
        <Typography.Text>Tạm tính</Typography.Text>
        <Typography.Text>{subTotal} đ</Typography.Text>
      </Flex>
      <Flex justify="space-between">
        <Typography.Text>Tiền giao hàng</Typography.Text>
        <Typography.Text>{shippingFee} đ</Typography.Text>
      </Flex>
      <Flex justify="space-between">
        <Typography.Text>Tổng giảm giá</Typography.Text>
        {discountTotal > 0 ? `-${discountTotal} đ` : `${discountTotal} đ`}
      </Flex>
      <Divider className="my-2" />
      <Flex justify="space-between">
        <Typography.Text style={{ fontWeight: "bold" }}>Tổng</Typography.Text>
        <Typography.Text style={{ fontWeight: "bold" }}>
          {subTotal - discountTotal}
        </Typography.Text>
      </Flex>
    </Card>
  );
};

export default PriceSummary;
