"use client";

import { useMeeting, useParticipant } from "@videosdk.live/react-sdk";
import { List, Avatar } from "antd";

const ParticipantPanel = () => {
  const { participants } = useMeeting();

  return (
    <div className="flex-1 overflow-y-scroll p-2">
      <List
        dataSource={Array.from(participants.keys())}
        renderItem={(participantId, index) => {
          const { micOn, webcamOn, displayName, isLocal, mode } =
            useParticipant(participantId);
          return (
            <List.Item>
              <List.Item.Meta
                style={{ alignItems: "center" }}
                avatar={<Avatar>{displayName?.charAt(0).toUpperCase()}</Avatar>}
                title={displayName}
              />
            </List.Item>
          );
        }}
      />
    </div>
  );
};

export default ParticipantPanel;
