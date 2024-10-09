"use client";
import { useState } from "react";
import { Flex, Button } from "antd";
import Link from "next/link";
import { useRouter, redirect } from "next/navigation";
import { Constants } from "@videosdk.live/react-sdk";

import LivestreamCalendar from "@/component/livestream/LivestreamCalendar";
import { useAppSelector } from "@/store/store";

const Page = () => {
  const router = useRouter();
  const { userId } = useAppSelector((state) => state.authReducer);
  const [livestreamId, setLivestreamId] = useState<number | null>(null);
  if (!userId) return null;
  return (
    <Flex vertical gap="middle">
      <Flex gap="small">
        <Link href="/seller/livestreams/create">
          <Button type="primary">Tạo livestream mới</Button>
        </Link>
        {livestreamId !== null && (
          <Button
            type="default"
            onClick={() => {
              window.location.href = `/livestreams/${livestreamId}`;
            }}
          >
            Bắt đầu livestream
          </Button>
        )}
      </Flex>
      <LivestreamCalendar
        shopId={userId!}
        mode={Constants.modes.CONFERENCE}
        selectedLivestreamId={livestreamId}
        setSelectedLivestreamId={setLivestreamId}
      />
    </Flex>
  );
};

export default Page;
