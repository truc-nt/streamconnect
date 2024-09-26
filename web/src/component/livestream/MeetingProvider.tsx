"use client";
import dynamic from "next/dynamic";
import React, {
  memo,
  useState,
  useEffect,
  createContext,
  useContext,
} from "react";
import { Flex, Typography } from "antd";

import { Constants } from "@videosdk.live/react-sdk";
import { useGetLivestream } from "@/hook/livestream";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { LivestreamStatus } from "@/constant/livestream";

const VideoSdkMeetingProvider = dynamic(
  () => import("@videosdk.live/react-sdk").then((mod) => mod.MeetingProvider),
  { ssr: false },
);

const MeetingContainer = dynamic(
  () => import("@/component/livestream/MeetingContainer"),
  {
    ssr: false,
  },
);

export const MeetingAppContext = createContext({
  livestreamId: 0,
  shopId: 0,
  shopName: "",
  livestreamStatus: "",
});

export const useMeetingAppContext = () => useContext(MeetingAppContext);

const MeetingProvider = ({ livestreamId }: { livestreamId: number }) => {
  const { data: livestream } = useGetLivestream(Number(livestreamId));
  const { username } = useAppSelector((state) => state.authReducer);
  console.log(livestream, livestream?.meeting_id);

  return livestream?.meeting_id &&
    livestream?.status !== LivestreamStatus.ENDED ? (
    <VideoSdkMeetingProvider
      config={{
        meetingId: livestream?.meeting_id,
        micEnabled: livestream?.is_host ? true : false,
        webcamEnabled: livestream?.is_host ? true : false,
        name: username ?? "",
        mode: livestream?.is_host ? "CONFERENCE" : "VIEWER",
        multiStream: false,
        debugMode: true,
      }}
      token={process.env.NEXT_PUBLIC_VIDEOSDK_TOKEN ?? ""}
      reinitialiseMeetingOnConfigChange={true}
      joinWithoutUserInteraction={true}
    >
      <MeetingAppContext.Provider
        value={{
          livestreamId,
          shopId: livestream?.id_shop,
          shopName: livestream?.shop_name,
          livestreamStatus: livestream?.status,
        }}
      >
        <MeetingContainer />
      </MeetingAppContext.Provider>
    </VideoSdkMeetingProvider>
  ) : livestream?.status === LivestreamStatus.ENDED ? (
    <Flex
      vertical
      justify="center"
      align="center"
      className="h-full bg-gray-800 rounded-lg"
    >
      <Typography.Title level={3} style={{ color: "white" }}>
        Người bán đã kết thúc buổi livestream
      </Typography.Title>
    </Flex>
  ) : (
    <h1>Loading...</h1>
  );
};

const MemoizedMeetingProvider = memo(
  MeetingProvider,
  (prevProps, nextProps) => {
    return prevProps.livestreamId === nextProps.livestreamId;
  },
);

export default MemoizedMeetingProvider;
