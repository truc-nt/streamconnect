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
import type { CalendarProps } from "antd";
import type { Dayjs } from "dayjs";
import dayjs from "dayjs";
import { Constants } from "@videosdk.live/react-sdk";

import { useGetAllLivestreams } from "@/hook/livestream";
import LivestreamProductList from "@/component/list/LivestreamProductList";

const LivestreamCalendar = ({
  shopId,
  mode,
  selectedLivestreamId,
  setSelectedLivestreamId,
}: {
  shopId: number;
  mode: string;
  selectedLivestreamId?: number | null;
  setSelectedLivestreamId?: (id: number | null) => void;
}) => {
  const { data: livestreams } = useGetAllLivestreams(shopId);
  /*const [selectedLivestreamId, setSelectedLivestreamId] = useState<
    number | null
  >(null);*/

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
            onClick={() => setSelectedLivestreamId?.(livestream.livestreamId)}
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
              status={
                livestreams?.find(
                  (livestream) =>
                    livestream.id_livestream === selectedLivestreamId,
                )?.status!
              }
            />
          </div>
        </Card>
      )}
    </Flex>
  );
};

export default LivestreamCalendar;
