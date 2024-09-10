"use client";
import React, { useState } from "react";
import { Button, message, Steps, Flex, Card } from "antd";
import LivestreamInformation from "./component/LivestreamInformation";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";

const Page = () => {
  return (
    <Flex vertical gap="large">
      <Card title="Thông tin">
        <LivestreamInformation />
      </Card>
      <Card title="Chọn sản phẩm">
        <ChosenLivestreamVariant shopId={1} />
      </Card>
    </Flex>
  );
};

export default Page;
