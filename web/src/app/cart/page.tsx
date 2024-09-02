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
import SummaryCard from "./component/Summary";
import AddressCard from "./component/AddressCard";
import { useAppSelector, useAppDispatch } from "@/store/store";
import useLoading from "@/hook/loading";
import { setCartItemIds } from "@/store/checkout";
import { useRouter } from "next/navigation";

const Page = () => {
  const { data: cart, error } = useGetCart(1);
  const dispatch = useAppDispatch();
  const router = useRouter();

  const { cartItems } = useAppSelector((state) => state.cartItemIdsSelection);
  const subTotal =
    cartItems.reduce(
      (total, cartItem) => total + cartItem.price * cartItem.quantity,
      0,
    ) || 0;

  const handleCheckout = () => {
    dispatch(setCartItemIds(cartItems.map((cartItem) => cartItem.cartItemId)));
    router.push("/checkout");
  };

  return (
    <Row gutter={[24, 24]}>
      <Col span={16}>
        <CartItemList cart={cart || []} />
      </Col>
      <Col span={8}>
        <Flex vertical gap="large">
          <AddressCard />
          <SummaryCard subTotal={subTotal} discountTotal={0} />
          <Button
            type="primary"
            size="large"
            disabled={!cartItems.length}
            onClick={handleCheckout}
          >
            Đặt hàng ({cartItems.length})
          </Button>
        </Flex>
      </Col>
    </Row>
  );
};

export default Page;
