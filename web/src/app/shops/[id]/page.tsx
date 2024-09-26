"use client";
import { useParams } from "next/navigation";
import { Typography, Flex, Tabs } from "antd";
import { Constants } from "@videosdk.live/react-sdk";

import ShopInfo from "@/component/info/ShopInfo";
import LivestreamCalendar from "@/component/livestream/LivestreamCalendar";

import { useAppSelector } from "@/store/store";

import { useGetShop } from "@/hook/shop";
const Page = ({ params }: { params: { id: string } }) => {
  const { data: shop } = useGetShop(parseInt(params.id));
  const { id: shopId } = useParams();
  const { userId } = useAppSelector((state) => state.authReducer);

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
            children: (
              <LivestreamCalendar
                shopId={Number(shopId)}
                mode={
                  userId == shopId
                    ? Constants.modes.CONFERENCE
                    : Constants.modes.VIEWER
                }
              />
            ),
          },
        ]}
      />
    </Flex>
  );
};

export default Page;
