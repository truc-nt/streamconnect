"use client";
import { useState } from "react";
import { theme, List } from "antd";
import { Flex, Card } from "antd";

import AddressModal from "@/component/modal/AddressModal";
import AddressInfo from "@/component/info/AddressInfo";
import { useGetAddresses } from "@/hook/user";

const Page = () => {
  const [openAddModal, setOpenAddModal] = useState(false);
  const { token } = theme.useToken();
  const { data: addresses, mutate: mutateAddresses } = useGetAddresses();
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
        <List
          dataSource={addresses}
          renderItem={(item, index) => (
            <List.Item style={{ border: 0 }}>
              <Card className="w-full">
                <AddressInfo {...item} />
              </Card>
            </List.Item>
          )}
        />
      </Flex>
      <AddressModal
        open={openAddModal}
        onCancel={() => setOpenAddModal(false)}
        successfullySubmitPostAction={() => mutateAddresses()}
      />
    </>
  );
};

export default Page;
