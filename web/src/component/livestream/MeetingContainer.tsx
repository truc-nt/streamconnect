"use client";
import React, { useRef, useEffect, useState } from "react";
import { Constants, useMeeting } from "@videosdk.live/react-sdk";
import dynamic from "next/dynamic";
import { Flex, Row, Col } from "antd";
import BottomBar from "./BottomBar";
import SideBar from "./SideBar";

const ConferenceViewGrid = dynamic(
  () => import("@/component/livestream/ConferenceViewGrid"),
  {
    ssr: false,
  },
);

const ViewerView = dynamic(() => import("@/component/livestream/ViewerView"), {
  ssr: false,
});

const MeetingContainer = () => {
  const mMeeting = useMeeting({
    onMeetingJoined: () => {
      //Pin the local participant if he joins in CONFERENCE mode
      if (mMeetingRef.current.localParticipant.mode == "CONFERENCE") {
        mMeetingRef.current.localParticipant.pin();
      }
    },
  });
  const [activePanel, setActivePanel] = useState("");

  const mMeetingRef = useRef<any>(null);
  useEffect(() => {
    mMeetingRef.current = mMeeting;
  }, [mMeeting]);

  if (!mMeeting || !mMeeting.localParticipant) {
    return <p>Loading...</p>;
  }

  return (
    <Flex vertical className="h-full w-full" gap="middle">
      <Row gutter={[8, 0]} className="w-full h-full">
        <Col span={activePanel === "" ? 24 : 17}>
          {mMeeting.localParticipant.mode == Constants.modes.CONFERENCE ? (
            <ConferenceViewGrid />
          ) : (
            <ViewerView />
          )}
        </Col>
        <Col span={activePanel === "" ? 0 : 7}>
          <SideBar activePanel={activePanel} />
        </Col>
      </Row>
      <BottomBar activePanel={activePanel} setActivePanel={setActivePanel} />
    </Flex>
  );
};

export default MeetingContainer;
