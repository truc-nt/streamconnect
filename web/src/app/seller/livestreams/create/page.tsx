"use client";
import React, { useState } from "react";
import { Button, message, Steps, Flex } from "antd";
import ChosenLivestreamVariant from "./component/ChosenLivestreamVariant";
import LivestreamInformation from "./component/LivestreamInformation";
import { useAppSelector, useAppDispatch } from "@/store/store";

const steps = [
  {
    title: "Thêm sản phẩm",
    component: <ChosenLivestreamVariant />,
  },
  {
    title: "Điền thông tin",
    component: <LivestreamInformation />,
  },
];

const Page = () => {
  const { currentStep } = useAppSelector(
    (state) => state.livestreamCreateReducer,
  );

  const items = steps.map((item) => ({ key: item.title, title: item.title }));

  return (
    <Flex vertical gap="large">
      <Steps current={currentStep} items={items} labelPlacement="vertical" />
      {steps[currentStep].component}
    </Flex>
  );
};

export default Page;
