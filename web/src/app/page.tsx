"use client"
import Image from "next/image";
import dynamic from "next/dynamic";
import CategorySlider from "@/component/livestream/CategorySlider";
import LivestreamPreview from "@/component/livestream/LivestreamPreview";
import { Typography } from "antd";
import {MeetingProvider} from "@videosdk.live/react-sdk";
import {sdkToken} from "@/api/axios";
import LivestreamsList from "@/component/livestream/LivestreamsList";

export default function Home() {
  const meetingId = "v00o-wrk7-p0dr";
  return (
      <>
        <CategorySlider />
        <MeetingProvider
            config={{
              meetingId: meetingId,
              micEnabled: true,
              webcamEnabled: true,
              name: "TestUser",
              mode: "VIEWER",
              multiStream: false,
              debugMode: true
            }}
            token={sdkToken}
            reinitialiseMeetingOnConfigChange={true}
            joinWithoutUserInteraction={true}
        >
          <LivestreamPreview/>
        </MeetingProvider>
        {/*<Typography.Title level={5} style={{ fontSize: "20px", fontWeight: "bold", color: "white" }}>*/}
        {/*  Livestreams Được Đề Xuất*/}
        {/*</Typography.Title>*/}
        {/*<LivestreamsList></LivestreamsList>*/}
      </>
  );
}
