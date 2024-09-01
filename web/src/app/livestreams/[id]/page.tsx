"use client"
import {sdkToken} from "@/api/axios";
import {MeetingProvider} from "@videosdk.live/react-sdk";
import LiveStreamPageContainer from "@/app/livestreams/[id]/IlsContainer";

const Page = ({ params }: { params: { id: number } }) => {
  const meetingId = "v00o-wrk7-p0dr";
  return (
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
        <LiveStreamPageContainer params={params} />
      </MeetingProvider>
  )
};

export default Page;
