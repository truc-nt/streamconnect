"use client";
import { useAppSelector } from "@/store/store";
import { useGetCartItemsByIds } from "@/hook/cart";
import {
  Tabs,
  Flex,
  Typography,
  Input,
  Card,
  Radio,
  Row,
  Col,
  Button,
  Form,
} from "antd";

import VoucherSection from "./VoucherSection";
import AddressCard from "./AddressCard";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";
import CartItemsGroupByShop from "../../app/checkout/component/CartItemsGroupByShop";
import PriceSummaryCard from "./PriceSummaryCard";
import { setNextStep } from "@/store/checkout";
import { useAppDispatch } from "@/store/store";

const FirstStep = () => {
  const { groupByShop } = useAppSelector((state) => state.checkoutReducer);
  const { data } = useGetCartItemsByIds(
    groupByShop.flatMap((shop) =>
      shop.groupByEcommerce.flatMap((ecommerce) => ecommerce.cartItemIds),
    ),
  );
  const dispatch = useAppDispatch();
  const subTotal =
    data?.reduce(
      (total, cartItemsByEcommerce) =>
        total +
        cartItemsByEcommerce.cart_items_group_by_shop.reduce(
          (total, cartItems) =>
            total +
            cartItems.cart_items.reduce(
              (total, cartItem) => total + cartItem.price * cartItem.quantity,
              0,
            ),
          0,
        ),
      0,
    ) ?? 0;

  const internalDiscountTotal = groupByShop.reduce((total, shop) => {
    return (
      total +
      shop.groupByEcommerce.reduce((ecommerceTotal, ecommerce) => {
        return ecommerceTotal + ecommerce.internalDiscountTotal;
      }, 0)
    );
  }, 0);

  const handleNextStep = () => {
    /*dispatch(
      setPrices([
        {
          subTotal,
          shippingFee: 0,
          discountTotal: 0,
          ecommerceId: 1,
        },
      ]),
    );*/
    dispatch(setNextStep());
  };

  return (
    <Flex vertical gap="large">
      <Tabs
        className="w-full"
        defaultActiveKey="1"
        items={data?.map((cartItemsByEcommerce) => {
          const platformLabel =
            ECOMMERCE_PLATFORMS[cartItemsByEcommerce.id_ecommerce] || "Unknown";
          return {
            label: platformLabel,
            key: cartItemsByEcommerce.id_ecommerce.toString(),
            children: cartItemsByEcommerce.cart_items_group_by_shop.map(
              (cartItems) => (
                <Row gutter={[24, 24]}>
                  <Col span={14}>
                    <Flex vertical gap="large">
                      <CartItemsGroupByShop
                        key={cartItems.id_shop}
                        {...cartItems}
                      />
                      <VoucherSection
                        shopId={groupByShop?.[0]?.shopId!}
                        ecommerceId={cartItemsByEcommerce.id_ecommerce}
                      />
                    </Flex>
                  </Col>
                  <Col span={10}>
                    <Flex vertical gap="large">
                      <AddressCard />
                      <Flex vertical>
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
                          <Card.Meta title="Hình thức giao hàng" />
                          <Radio.Group name="radiogroup" defaultValue={1}>
                            <Radio value={1} className="w-full">
                              <Flex justify="space-between">
                                <Typography.Text>Standard</Typography.Text>
                              </Flex>
                            </Radio>
                          </Radio.Group>
                        </Card>
                      </Flex>
                      <PriceSummaryCard
                        subTotal={subTotal}
                        shippingFee={0}
                        internalDiscountTotal={internalDiscountTotal}
                        externalDiscountTotal={0}
                      />
                    </Flex>
                  </Col>
                </Row>
              ),
            ),
          };
        })}
      />
      <Flex gap="small" className="w-full" justify="end">
        <Button type="primary" onClick={handleNextStep}>
          Tiếp theo
        </Button>
      </Flex>
    </Flex>
  );
};

export default FirstStep;
