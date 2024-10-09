"use client";
import React, { useRef, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { Constants, useMeeting } from "@videosdk.live/react-sdk";
import dynamic from "next/dynamic";
import { Flex, Row, Col, Typography, Button, Tabs } from "antd";

import ShopInfo from "@/component/info/ShopInfo";
import LivestreamStatistic from "@/component/statistic/LivestreamStatistic";
import BottomBar from "./BottomBar";
import SideBar from "./SideBar";
import { useMeetingAppContext } from "./MeetingProvider";

const ConferenceViewGrid = dynamic(
  () => import("@/component/livestream/ConferenceViewGrid"),
  {
    ssr: false,
  },
);

const PresenterView = dynamic(
  () => import("@/component/livestream/PresenterView"),
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
  const [activeTab, setActiveTab] = useState("1");
  const [activePanel, setActivePanel] = useState("");
  const isPresenting = mMeeting.presenterId ? true : false;
  console.log("isPresenting", isPresenting);

  const mMeetingRef = useRef<any>(null);
  const { shopName, shopId } = useMeetingAppContext();
  useEffect(() => {
    mMeetingRef.current = mMeeting;
  }, [mMeeting]);

  if (!mMeeting || !mMeeting.localParticipant) {
    return <p>Loading...</p>;
  }

  return (
    <Flex vertical className="h-full w-full" gap="middle">
      {mMeeting.localParticipant.mode == Constants.modes.CONFERENCE && (
        <Tabs
          centered
          onChange={(key) => setActiveTab(key)}
          items={[
            {
              label: "Livestream",
              key: "1",
            },
            {
              label: "Thống kê",
              key: "2",
            },
          ]}
        />
      )}
      {activeTab === "1" ? (
        <>
          <Row gutter={[8, 0]} className="w-full h-full">
            <Col span={activePanel === "" ? 24 : 17}>
              {mMeeting.localParticipant.mode == Constants.modes.CONFERENCE ? (
                isPresenting ? (
                  <PresenterView />
                ) : (
                  <ConferenceViewGrid />
                )
              ) : (
                <Flex vertical gap="middle" className="h-full">
                  <div className="flex-1">
                    <ViewerView />
                  </div>
                  <ShopInfo
                    id_shop={shopId}
                    name={shopName}
                    is_following={false}
                    description=""
                  />
                </Flex>
              )}
            </Col>
            <Col span={activePanel === "" ? 0 : 7}>
              <SideBar activePanel={activePanel} />
            </Col>
          </Row>
          <BottomBar
            activePanel={activePanel}
            setActivePanel={setActivePanel}
          />
        </>
      ) : (
        <LivestreamStatistic />
      )}
    </Flex>
  );
};

export default MeetingContainer;
