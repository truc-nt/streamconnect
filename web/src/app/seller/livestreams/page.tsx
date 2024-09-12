"use client";
import { Flex, Button } from "antd";
import Link from "next/link";

import LivestreamCalendar from "@/component/livestream/LivestreamCalendar";
import { useAppSelector } from "@/store/store";

const Page = () => {
  const { userId } = useAppSelector((state) => state.authReducer);
  return (
    <Flex vertical gap="middle">
      <Link href="/seller/livestreams/create">
        <Button type="primary">Tạo livestream mới</Button>
      </Link>
      <LivestreamCalendar shopId={userId!} editable={true} />
    </Flex>
  );
};

export default Page;
