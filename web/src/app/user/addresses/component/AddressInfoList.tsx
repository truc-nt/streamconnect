"use client";
import { useGetAddresses } from "@/hook/user";
import { List } from "antd";
import AddressInfo from "./AddressInfo";

const AddressInfoList = () => {
  const { data } = useGetAddresses();

  return (
    <List
      dataSource={data}
      renderItem={(item, index) => (
        <List.Item style={{ border: 0 }}>
          <AddressInfo {...item} />
        </List.Item>
      )}
    />
  );
};

export default AddressInfoList;
