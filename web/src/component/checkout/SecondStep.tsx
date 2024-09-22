import AddressCard from "@/app/checkout/component/AddressCard";
import { Row, Col, Button, Flex, Card, Radio, Typography } from "antd";
import PriceSummaryCard from "./PriceSummaryCard";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setPrevStep } from "@/store/checkout";
import useLoading from "@/hook/loading";
import { createOrder } from "@/api/order";
import { useRouter } from "next/navigation";
import { reset } from "@/store/checkout";

const CheckoutStep = () => {
  const { groupByShop, addressId } = useAppSelector(
    (state) => state.checkoutReducer,
  );
  const { userId } = useAppSelector((state) => state.authReducer);
  const dispatch = useAppDispatch();
  const router = useRouter();
  const handleCreateOrder = useLoading(
    createOrder,
    "Tạo đơn hàng thành công",
    "Tạo đơn hàng thất bại",
  );

  const handleCheckout = async () => {
    try {
      const orderPromises = groupByShop.flatMap((shop) =>
        shop.groupByEcommerce.map(async (externalOrder) => {
          const createOrderBody = {
            cart_item_ids: externalOrder.cartItemIds,
            shipping_fee: 0,
            shipping_fee_discount: 0,
            external_discount: 0,
            voucher_ids: externalOrder.voucherIds,
          };

          return handleCreateOrder({
            id_user_address: addressId,
            id_shop: shop.shopId,
            external_orders: [createOrderBody],
          });
        }),
      );

      await Promise.all(orderPromises);

      router.push("/cart");
      dispatch(reset());
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Row gutter={[24, 24]}>
      <Col span={16}>
        {groupByShop.map((shop) =>
          shop.groupByEcommerce.map((ecommerce) => (
            <PriceSummaryCard
              key={ecommerce.ecommerceId}
              subTotal={ecommerce.subTotal}
              shippingFee={0}
              internalDiscountTotal={ecommerce.internalDiscountTotal}
              externalDiscountTotal={0}
              ecommerceId={ecommerce.ecommerceId}
            />
          )),
        )}
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
