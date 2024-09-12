import { Card } from "antd";
import { useState } from "react";
import ParticipantPanel from "./ParticipantPanel";
import ChatPanel from "./ChatPanel";
import ProductPanel from "./ProductPanel";
import VoucherPanel from "./VoucherPanel";

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
    {
      value: "voucher",
      title: "Voucher",
      component: <VoucherPanel />,
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
              //height: "calc(100vh - 250px)",
              maxWidth: "100%",
              maxHeight: "calc(100% - 56px)",
              gap: "0.25rem",
              padding: "1rem",
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
