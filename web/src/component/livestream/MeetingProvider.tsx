"use client";
import dynamic from "next/dynamic";
import React, { memo } from "react";

const VideoSdkMeetingProvider = dynamic(
  () => import("@videosdk.live/react-sdk").then((mod) => mod.MeetingProvider),
  { ssr: false },
);

interface IViewerProviderProps {
  meetingId: string;
  mode: "VIEWER" | "CONFERENCE";
  name?: string;
  children?: React.ReactNode;
}

const MeetingProvider = ({
  meetingId,
  mode,
  name = "TestUser",
  children,
}: IViewerProviderProps) => {
  return (
    <VideoSdkMeetingProvider
      config={{
        meetingId: meetingId,
        micEnabled: mode === "VIEWER" ? false : true,
        webcamEnabled: mode === "VIEWER" ? false : true,
        name: name,
        mode: mode,
        multiStream: false,
        debugMode: true,
      }}
      token={process.env.NEXT_PUBLIC_VIDEOSDK_TOKEN ?? ""}
      reinitialiseMeetingOnConfigChange={true}
      joinWithoutUserInteraction={true}
    >
      {children}
    </VideoSdkMeetingProvider>
  );
};

const MemoizedMeetingProvider = memo(
  MeetingProvider,
  (prevProps, nextProps) => {
    return (
      prevProps.meetingId === nextProps.meetingId &&
      prevProps.name === nextProps.name &&
      prevProps.mode === nextProps.mode
    );
  },
);

export default MemoizedMeetingProvider;
