"use client";
import React, { useMemo, useState, useEffect } from "react";
import { Constants, useMeeting } from "@videosdk.live/react-sdk";
import ConferenceView from "./ConferenceView";
import { Flex, Row, Col } from "antd";
import BottomBar from "./BottomBar";
import SideBar from "./SideBar";

const ConferenceViewGrid = () => {
  const { participants, hlsState } = useMeeting();
  const [activePanel, setActivePanel] = useState("");

  const hosts = useMemo(() => {
    const speakerParticipants = Array.from(participants.values()).filter(
      (participant) => {
        return participant.mode === Constants.modes.CONFERENCE;
      },
    );
    return speakerParticipants;
  }, [participants]);

  const itemPerRow = hosts.length == 1 ? 1 : hosts.length <= 4 ? 2 : 3;
  return (
    <Row gutter={[4, 4]} className="w-full h-full">
      {hosts.map((participant) => (
        <Col key={participant.id} span={24 / itemPerRow}>
          <ConferenceView participantId={participant.id} />
        </Col>
      ))}
    </Row>
  );
};

export default ConferenceViewGrid;
