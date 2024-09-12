"use client";
import { Flex, Steps } from "antd";
import FirstStep from "@/component/checkout/FirstStep";
import { useAppSelector, useAppDispatch } from "@/store/store";
import SecondStep from "@/component/checkout/SecondStep";

const steps = [
  {
    title: "Kiểm tra sản phẩm đặt hàng",
    component: <FirstStep />,
  },
  {
    title: "Thanh toán",
    component: <SecondStep />,
  },
];

const Page = () => {
  const { currentStep } = useAppSelector((state) => state.checkoutReducer);
  return (
    <Flex gap="small" className="w-full">
      <Steps
        direction="vertical"
        progressDot
        current={currentStep}
        items={steps.map((item) => ({ key: item.title, title: item.title }))}
        className="w-[200px]"
      />
      <div className="flex-1">{steps[currentStep].component}</div>
    </Flex>
  );
};

export default Page;
