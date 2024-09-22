import { Flex, Divider, Typography } from "antd";
interface IPriceInfo {
  subTotal: number;
  internalDiscount: number;
  externalDiscount: number;
  shippingFee: number;
}
const PriceInfo = ({
  subTotal,
  internalDiscount,
  externalDiscount,
  shippingFee,
}: IPriceInfo) => {
  return (
    <>
      <Flex justify="space-between">
        <Typography.Text>Tạm tính</Typography.Text>
        <Typography.Text>{subTotal} đ</Typography.Text>
      </Flex>
      <Flex justify="space-between">
        <Typography.Text>Tiền giao hàng</Typography.Text>
        <Typography.Text>{shippingFee} đ</Typography.Text>
      </Flex>
      <Flex justify="space-between">
        <Typography.Text>Tổng giảm giá sàn</Typography.Text>
        {internalDiscount > 0
          ? `-${internalDiscount} đ`
          : `${internalDiscount} đ`}
      </Flex>
      <Flex justify="space-between">
        <Typography.Text>Tổng giảm giá ngoài sàn</Typography.Text>
        {externalDiscount > 0
          ? `-${externalDiscount} đ`
          : `${externalDiscount} đ`}
      </Flex>
      <Divider className="my-2" />
      <Flex justify="space-between">
        <Typography.Text style={{ fontWeight: "bold" }}>Tổng</Typography.Text>
        <Typography.Text style={{ fontWeight: "bold" }}>
          {subTotal - internalDiscount - externalDiscount} đ
        </Typography.Text>
      </Flex>
    </>
  );
};

export default PriceInfo;
