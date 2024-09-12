"use client";
import {
  Space,
  Flex,
  Table,
  Button,
  InputNumber,
  Input,
  Card,
  Row,
  Col,
} from "antd";
import { useGetCart } from "@/hook/cart";
import CartItemList from "./component/CartItemList";
import PriceSummaryCard from "@/component/checkout/PriceSummaryCard";
import AddressCard from "./component/AddressCard";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { useRouter } from "next/navigation";

const Page = () => {
  const dispatch = useAppDispatch();
  const router = useRouter();

  const { groupByShop } = useAppSelector((state) => state.checkoutReducer);

  const subTotal = groupByShop.reduce((total, shop) => {
    const shopSubTotal = shop.groupByEcommerce.reduce((subtotal, ecommerce) => {
      return subtotal + ecommerce.subTotal;
    }, 0);
    return total + shopSubTotal;
  }, 0);

  const { data: cart, error } = useGetCart();
  const handleCheckout = () => {
    router.push("/checkout");
  };

  return (
    <Row gutter={[24, 24]}>
      <Col span={16}>
        <CartItemList cart={cart || []} />
      </Col>
      <Col span={8}>
        <Flex vertical gap="large">
          <PriceSummaryCard
            subTotal={subTotal}
            internalDiscountTotal={0}
            externalDiscountTotal={0}
            shippingFee={0}
          />
          <Button
            type="primary"
            size="large"
            disabled={!groupByShop.length}
            onClick={handleCheckout}
          >
            Đặt hàng ({groupByShop.length})
          </Button>
        </Flex>
      </Col>
    </Row>
  );
};

export default Page;
