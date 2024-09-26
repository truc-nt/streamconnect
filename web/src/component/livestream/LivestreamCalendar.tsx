"use client";

import { useParams } from "next/navigation";
import { useState } from "react";

import {
  Card,
  List,
  Calendar,
  Tag,
  Flex,
  Space,
  Button,
  Modal,
  Checkbox,
} from "antd";
import { CheckboxChangeEvent } from "antd/lib/checkbox";
import {
  HeartOutlined,
  PushpinOutlined,
  PlusCircleOutlined,
} from "@ant-design/icons";
import type { CalendarProps } from "antd";
import type { Dayjs } from "dayjs";
import dayjs from "dayjs";
import { Constants } from "@videosdk.live/react-sdk";

import { useGetAllLivestreams } from "@/hook/livestream";
import { useGetLivestreamProducts } from "@/hook/livestream";
import { notifyLivestreamProductFollowers } from "@/api/notification";
import useLoading from "@/hook/loading";
import {
  updateLivestreamProductPriority,
  IPinLivestreamProduct,
} from "@/api/livestream_product";
import { addLivestreamProduct, ILivestreamProduct } from "@/api/livestream";
import { registerLivestreamProductFollower } from "@/api/livestream_product";
import { IChosenLivestreamVariant } from "@/app/seller/livestreams/create/component/LivestreamCreate";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";
import LivestreamProductList from "@/component/list/LivestreamProductList";

const LivestreamCalendar = ({
  shopId,
  mode,
}: {
  shopId: number;
  mode: string;
}) => {
  const { data: livestreams } = useGetAllLivestreams(shopId);
  const [selectedLivestreamId, setSelectedLivestreamId] = useState<
    number | null
  >(null);

  const getListData = (value: Dayjs) => {
    const dateString = value.format("YYYY-MM-DD");
    return livestreams
      ?.filter(
        (livestream) =>
          dayjs(livestream.start_time).format("YYYY-MM-DD") === dateString,
      )
      .map((livestream) => ({
        title: livestream.title,
        livestreamId: livestream.id_livestream,
      }));
  };

  const dateCellRender = (value: Dayjs) => {
    const listData = getListData(value);
    return (
      <div>
        {listData?.map((livestream, index) => (
          <Tag
            key={index}
            color="blue"
            onClick={() => setSelectedLivestreamId(livestream.livestreamId)}
          >
            {livestream.title}
          </Tag>
        ))}
      </div>
    );
  };

  const cellRender: CalendarProps<Dayjs>["cellRender"] = (current, info) => {
    if (info.type === "date") return dateCellRender(current);
    return info.originNode;
  };
  return (
    <Flex gap="small">
      <Calendar cellRender={cellRender} />
      {selectedLivestreamId && (
        <Card
          className="h-full w-full flex flex-col"
          title="Các sản phẩm trong phiên livestream"
          styles={{
            body: {
              flex: "1 1 0%",
              display: "flex",
              flexDirection: "column",
              gap: "0.25rem",
              padding: "1rem",
              justifyContent: "center",
            },
          }}
        >
          <div className="flex flex-col gap-1 p-2">
            <LivestreamProductList
              shopId={shopId}
              livestreamId={selectedLivestreamId}
              mode={mode}
            />
          </div>
        </Card>
      )}
    </Flex>
  );
};

export default LivestreamCalendar;
