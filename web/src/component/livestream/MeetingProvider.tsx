"use client";
import dynamic from "next/dynamic";
import React, {
  memo,
  useState,
  useEffect,
  createContext,
  useContext,
} from "react";
import { useGetLivstreamInfo } from "@/hook/livestream";
import { useAppDispatch, useAppSelector } from "@/store/store";

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
});

export const useMeetingAppContext = () => useContext(MeetingAppContext);

const MeetingProvider = ({ livestreamId }: { livestreamId: number }) => {
  const { data: getLivestreamInfoResponse } = useGetLivstreamInfo(
    Number(livestreamId),
  );
  const { username } = useAppSelector((state) => state.authReducer);

  return getLivestreamInfoResponse ? (
    <VideoSdkMeetingProvider
      config={{
        meetingId: getLivestreamInfoResponse.meeting_id,
        micEnabled: getLivestreamInfoResponse?.is_host ? true : false,
        webcamEnabled: getLivestreamInfoResponse?.is_host ? true : false,
        name: username ?? "",
        mode: getLivestreamInfoResponse?.is_host ? "CONFERENCE" : "VIEWER",
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
          shopId: getLivestreamInfoResponse?.id_shop,
          shopName: getLivestreamInfoResponse?.shop_name,
        }}
      >
        <MeetingContainer />
      </MeetingAppContext.Provider>
    </VideoSdkMeetingProvider>
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
