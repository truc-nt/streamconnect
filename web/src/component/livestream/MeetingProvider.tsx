"use client";

import { MeetingProvider as VideoSdkMeetingProvider } from "@videosdk.live/react-sdk";
import { sdkToken } from "@/api/axios";
import ReactPlayer from "react-player";

interface IMeetingProviderProps {
  meetingId: string;
  name?: string;
  mode?: "CONFERENCE" | "VIEWER";
}

const MeetingProvider = ({
  meetingId,
  name = "TestUser",
  mode = "VIEWER",
}: IMeetingProviderProps) => {
  return (
    <VideoSdkMeetingProvider
      config={{
        meetingId: meetingId,
        micEnabled: true,
        webcamEnabled: true,
        name: name,
        mode: mode,
        multiStream: false,
        debugMode: true,
      }}
      token={sdkToken}
      reinitialiseMeetingOnConfigChange={true}
      joinWithoutUserInteraction={true}
    >
        <ReactPlayer
          //
          playsinline // extremely crucial prop
          pip={false}
          light={false}
          controls={false}
          muted={true}
          playing={true}
          //
          url={videoStream}
          //
          height={"200px"}
          width={"300px"}
          onError={(err) => {
            console.log(err, "participant video error");
          }}
        />
    </VideoSdkMeetingProvider>
  );
};

export default MeetingProvider;