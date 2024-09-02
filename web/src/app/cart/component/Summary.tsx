import { Card, Divider, Flex, Typography, Button } from "antd";

const Summary = ({
  subTotal,
  discountTotal,
}: {
  subTotal: number;
  discountTotal: number;
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
    <Card.Meta title="Tổng tiền" />
    <Flex justify="space-between">
      <Typography.Text>Tạm tính</Typography.Text>
      <Typography.Text>{subTotal} đ</Typography.Text>
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

export default Summary;
