"use client";
import { useState } from "react";
import { theme } from "antd";
import AddButton from "./component/AddButton";
import AddressInfoList from "./component/AddressInfoList";
import { Flex } from "antd";

import AddressModal from "@/component/modal/AddressModal";
const Page = () => {
  const [openAddModal, setOpenAddModal] = useState(false);
  const { token } = theme.useToken();
  return (
    <>
      <Flex gap="middle" vertical>
        <div
          style={{
            display: "flex",
            height: "80px",
            alignItems: "center",
            justifyContent: "center",
            color: token.colorTextTertiary,
            backgroundColor: token.colorFillAlter,
            borderRadius: token.borderRadiusLG,
            border: `1px dashed ${token.colorBorder}`,
            cursor: "pointer",
          }}
          onClick={() => setOpenAddModal(true)}
        >
          Thêm địa chỉ
        </div>
        <AddressInfoList />
      </Flex>
      <AddressModal
        open={openAddModal}
        onCancel={() => setOpenAddModal(false)}
      />
    </>
  );
};

export default Page;
