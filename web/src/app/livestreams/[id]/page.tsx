import { Col, Row } from "antd";
import dynamic from "next/dynamic";

import MeetingProvider from "@/component/livestream/MeetingProvider";

const MeetingContainer = dynamic(
  () => import("@/component/livestream/MeetingContainer"),
  {
    ssr: false,
  },
);

const Page = ({ params }: { params: { id: number } }) => {
  const meetingId = "klu7-zzex-qlgu";
  return (
    <MeetingProvider meetingId={meetingId} mode="CONFERENCE">
      <MeetingContainer />
    </MeetingProvider>
  );
};

export default Page;
