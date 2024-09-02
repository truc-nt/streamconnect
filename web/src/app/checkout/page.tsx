"use client";
import { Flex, Steps } from "antd";
import CartItemsReviewStep from "./component/CartItemsReviewStep";
import { useAppSelector, useAppDispatch } from "@/store/store";
import CheckoutStep from "./component/CheckoutStep";

const steps = [
  {
    title: "Kiểm tra sản phẩm đặt hàng",
    component: <CartItemsReviewStep />,
  },
  {
    title: "Thanh toán",
    component: <CheckoutStep />,
  },
];

const Page = () => {
  const { currentStep } = useAppSelector((state) => state.checkoutReducer);
  return (
    <Flex vertical gap="large" align="center" className="w-full">
      <Steps
        labelPlacement="vertical"
        current={currentStep}
        items={steps.map((item) => ({ key: item.title, title: item.title }))}
        className="w-[800px]"
      />
      {steps[currentStep].component}
    </Flex>
  );
};

export default Page;
