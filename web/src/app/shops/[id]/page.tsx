"use client";
import { useParams } from "next/navigation";
import { Typography, Flex, Tabs } from "antd";

import ShopInfo from "@/component/info/ShopInfo";
import LivestreamCalendar from "../../../component/livestream/LivestreamCalendar";

import { useGetShop } from "@/hook/shop";
const Page = ({ params }: { params: { id: string } }) => {
  const { data: shop } = useGetShop(parseInt(params.id));
  const { id: shopId } = useParams();
  return (
    <Flex vertical gap="middle">
      <ShopInfo {...shop!} />
      <Tabs
        className="w-full"
        defaultActiveKey="1"
        items={[
          {
            label: "Giới thiệu",
            key: "1",
            children: (
              <Typography.Paragraph>{shop?.description}</Typography.Paragraph>
            ),
          },
          {
            label: "Lịch phát",
            key: "2",
            children: <LivestreamCalendar shopId={Number(shopId)} />,
          },
        ]}
      />
    </Flex>
  );
};

export default Page;
