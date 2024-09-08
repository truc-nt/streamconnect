import { Card } from "antd";
import { useState } from "react";
import ParticipantPanel from "./ParticipantPanel";
import ChatPanel from "./ChatPanel";
import ProductPanel from "./ProductPanel";
import { useMeeting } from "@videosdk.live/react-sdk";

const SideBar = ({ activePanel }: { activePanel: string }) => {
  const panels = [
    {
      value: "participant",
      title: "Nguời xem livestream",
      component: <ParticipantPanel />,
    },
    {
      value: "chat",
      title: "Bình luận",
      component: <ChatPanel />,
    },
    {
      value: "product",
      title: "Sản phẩm",
      component: <ProductPanel />,
    },
  ];
  return (
    <div className="relative h-full w-full">
      <div className="absolute left-0 top-0 bottom-0 right-0">
        <Card
          className="h-full w-full flex flex-col"
          title={panels.find((panel) => panel.value === activePanel)?.title}
          styles={{
            body: {
              flex: "1 1 0%",
              display: "flex",
              flexDirection: "column",
              height: "calc(100vh - 56px)",
              maxHeight: "calc(100vh - 64px)",
              gap: "1rem",
            },
          }}
        >
          {panels.find((panel) => panel.value === activePanel)?.component}
        </Card>
      </div>
    </div>
  );
};

export default SideBar;
