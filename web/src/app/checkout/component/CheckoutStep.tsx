import AddressCard from "./AddressCard";
import { Row, Col, Button, Flex, Card, Radio, Typography } from "antd";
import PriceSummary from "./PriceSummary";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setPrevStep } from "@/store/checkout";
import useLoading from "@/hook/loading";
import { createOrder } from "@/api/order";
import { useRouter } from "next/navigation";

const CheckoutStep = () => {
  const { prices, cartItemIds } = useAppSelector(
    (state) => state.checkoutReducer,
  );
  const dispatch = useAppDispatch();
  const router = useRouter();
  const handleCreateOrder = useLoading(
    createOrder,
    "Tạo đơn hàng thành công",
    "Tạo đơn hàng thất bại",
  );

  const handleCheckout = async () => {
    try {
      await handleCreateOrder({
        id_user: 1,
        id_cart_items: cartItemIds,
        address: "Hà Nội",
      });
      router.push("/cart");
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Row gutter={[24, 24]}>
      <Col span={16}>
        {prices.map((price) => (
          <PriceSummary
            key={price.ecommerceId}
            subTotal={price.subTotal}
            shippingFee={price.shippingFee}
            discountTotal={price.discountTotal}
            ecommerceId={price.ecommerceId}
          />
        ))}
      </Col>
      <Col span={8}>
        <Flex vertical gap="large">
          <AddressCard />
          <Card
            styles={{
              body: {
                display: "flex",
                flexDirection: "column",
                height: "100%",
                gap: "0.5rem",
              },
            }}
          >
            <Card.Meta title="Phương thức thanh toán" />
            <Radio.Group name="radiogroup" defaultValue={1}>
              <Radio value={1} className="w-full">
                <Flex justify="space-between">
                  <Typography.Text>Thanh toán khi nhận hàng</Typography.Text>
                </Flex>
              </Radio>
            </Radio.Group>
          </Card>
        </Flex>
      </Col>
      <Flex gap="small" className="w-full" justify="end">
        <Button type="default" onClick={() => dispatch(setPrevStep())}>
          Quay lại
        </Button>
        <Button type="primary" onClick={handleCheckout}>
          Thanh toán
        </Button>
      </Flex>
    </Row>
  );
};

export default CheckoutStep;
