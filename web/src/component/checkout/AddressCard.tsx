import { Card, Typography, Divider, Flex } from "antd";

import AddressInfo from "@/component/info/AddressInfo";
import { useGetDefaultAddress } from "@/hook/user";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setAddressId } from "@/store/checkout";
import { useEffect } from "react";

const AddressCard = () => {
  const dispatch = useAppDispatch();
  const { data } = useGetDefaultAddress();

  useEffect(() => {
    if (data) {
      console.log("hello", data);
      dispatch(setAddressId(data.id_user_address));
    }
  }, []);

  return (
    <Card
      bordered={false}
      styles={{
        body: {
          display: "flex",
          flexDirection: "column",
          height: "100%",
          gap: "0.5rem",
        },
      }}
    >
      <Card.Meta
        title={
          <Flex justify="space-between">
            <Typography.Text style={{ fontSize: "16px" }}>
              Địa chỉ
            </Typography.Text>
            <Typography.Text>Thay đổi</Typography.Text>
          </Flex>
        }
      />
      <AddressInfo {...data!} />
    </Card>
  );
};

export default AddressCard;
