"use client";
import { Segmented, Card, Space, Flex } from "antd";
import { SkinFilled, MessageFilled } from "@ant-design/icons";
import LivestreamProductSegmented from "./LivestreamProductSegmented";
import { useState } from "react";
import LivestreamProductInfo from "./LivestreamProductInfo";

const Sidebar = ({ livestreamId }: { livestreamId: number }) => {
  const segmentedOptions = [
    {
      label: "Sản phẩm",
      value: "product",
      icon: <SkinFilled />,
      component: <LivestreamProductSegmented livestreamId={livestreamId} />,
    },
    {
      label: "Bình luận",
      value: "comment",
      icon: <MessageFilled />,
      component: <div>Bình luận</div>,
    },
  ];

  const [activeSegmented, setActiveSegmented] = useState(
    segmentedOptions[0].value,
  );

  return (
    <Card
      className="absolute left-0 top-0 bottom-0 right-0"
      styles={{
        body: {
          display: "flex",
          flexDirection: "column",
          height: "100%",
          gap: "1rem",
        },
      }}
    >
      <Segmented
        options={segmentedOptions}
        value={activeSegmented}
        onChange={(value) => setActiveSegmented(value)}
        block
      />
      {
        segmentedOptions.find((option) => option.value === activeSegmented)
          ?.component
      }
    </Card>
  );
};

export default Sidebar;
